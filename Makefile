# NOTE: For a list of targets, run `make help`. <skr 2022-09-10>

.PHONY: all help test

SHELL = /usr/bin/env bash

ifneq ($(TRACE),yes)
Q = @
endif

help: BD := $(shell tput bold)
help: BL := $(shell tput setaf 4)
help: CY := $(shell tput setaf 6)
help: RS := $(shell tput sgr0)

help: HELP = '\n\
$(BD)weird targets:$(RS)\n\
  $(CY)all$(RS)   - $(BL)Test all solutions$(RS)\n\
  $(CY)N$(RS)     - $(BL)Test solution N$(RS)\n\
\n'

help:
	$(Q)printf $(HELP)
SUBDIRS := $(wildcard */.)

all: SUBDIRS := $(wildcard src/*/.)  # e.g. "foo/. bar/."
all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@ all

test:
	https://stackoverflow.com/a/11206700

