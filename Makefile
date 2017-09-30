EXAMPLES := $(shell find examples/ -maxdepth 1 -type d -exec sh -c 'echo $(basename {})' \;)
EXLIST := $(subst examples/,,$(EXAMPLES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean lint test coverage $(EXLIST)

lint:
	@script/lint

test:
	@script/test

coverage:
	@script/coverage

$(EXLIST):
	@echo $@
	@go test -v ./examples/$@
	@gocov test ./examples/$@ | gocov report

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean lint test coverage $(EXLIST)
