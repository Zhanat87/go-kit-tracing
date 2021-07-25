#!/usr/bin/env bash

make vendor
make tests
make lint
git add . && git commit -am 'tracing-0.0.2' && git push
git tag v0.0.2 && git push --tags
