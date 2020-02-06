# Setup name variables for the package/tool
NAME := checkcapabilities
PKG := github.com/sanddorn/$(NAME)

CGO_ENABLED := 0

# Set any default go build tags.
BUILDTAGS :=

include basic.mk

.PHONY: prebuild
prebuild:
