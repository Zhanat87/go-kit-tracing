package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zhanat87/common-libs/httphandlers"
	"github.com/Zhanat87/common-libs/loggers"
	"github.com/Zhanat87/common-libs/tracers"
	"github.com/Zhanat87/go-kit-tracing/factory"
	"github.com/Zhanat87/go-kit-tracing/middleware"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	apphttp "github.com/Zhanat87/go-kit-tracing/transport/http"
	"github.com/go-kit/kit/log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	httpAddr := os.Getenv("HTTP_ADDR")
	logger := new(loggers.GoKitLoggerFactory).CreateLogger()
	httpLogger := log.With(logger, "component", "http")
	serviceName := os.Getenv("SERVICE_NAME")
	err := tracers.InitZipkinTracerAndZipkinHTTPReporter(serviceName, ":0")
	if err != nil {
		panic(err)
	}
	defer tracers.ZipkinReporter.Close()
	mux := http.NewServeMux()
	mux.Handle(pong.BaseURL, apphttp.MakePongHandler(
		middleware.MakePongEndpoints(new(factory.PongServiceFactory).CreateHTTPService(pong.PackageName, httpLogger, tracers.ZipkinTracer)),
		httpLogger, pong.BaseURL))
	httphandlers.InitDefaultHandlers(mux)
	errs := make(chan error, 2)
	go func() {
		_ = logger.Log("transport", "http", "address", httpAddr, "msg", "listening "+serviceName+" api")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	_ = logger.Log("terminated", <-errs)
}
