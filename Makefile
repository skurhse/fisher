# NOTE: For a list of targets, run `make help`. <skr 2022-09-10>

SHELL = /usr/bin/env bash

.PHONY: help test

help: BD := $(shell tput bold)
help: BL := $(shell tput setaf 4)
help: CY := $(shell tput setaf 6)
help: RS := $(shell tput sgr0)

help: HELP = '\n\
$(BD)weird targets:$(RS)\n\
  $(CY)test$(RS)   - $(BL)Test all appoaches$(RS)\n\
\n'

help:
	$(Q)printf $(HELP)

test:
	go test -v ./...
