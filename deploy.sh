#!/usr/bin/env bash

make vendor
make tests
make lint
git add . && git commit -am 'tracing-0.0.1' && git push
git tag v0.0.1 && git push --tags
