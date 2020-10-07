.DEFAULT_GOAL = test
.PHONY: FORCE

export GOPROXY = https://goproxy.cn,direct

# Build

build: go-build
.PHONY: build

clean: go-clean
.PHONY: clean

lint: go-lint
.PHONY: lint

test: go-test
.PHONY: test

# Non-PHONY targets (real files)

go-build: FORCE
	./script/build.sh

go-clean: FORCE
	./script/clean.sh

go-lint: FORCE
	./script/lint.sh

go-test: FORCE
	./script/test.sh
