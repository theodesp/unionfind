# Go parameters
BUILDPATH=$(CURDIR)
GOCMD=go
GODOC=godoc
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOBENCH=$(GOCMD) test -check.b
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
GOGET=$(GOCMD) get
GOTESTFLAGS=-check.vv
GOCOVERFLAGS=-covermode=count -coverprofile=count.out
GOFLAGS=-race

# Package lists
TOPLEVEL_PKG := github.com/theodesp/unionfind

get:
	@echo "getting all dependencies"
	$(GOGET)

build:
	$(GOINSTALL) $(GOFLAGS) $(TOPLEVEL_PKG)
	@echo "build sucessfully"

clean:
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

test:
	@echo "Running Tests..."
	$(GOTEST) $(GOTESTFLAGS) $(TOPLEVEL_PKG)


bench:
	@echo "Running Benchmarks..."
	$(GOBENCH) $(TOPLEVEL_PKG)

cover:
	@echo "Running Coverage Report..."
	$(GOTEST) $(GOCOVERFLAGS) $(TOPLEVEL_PKG)


format:
	$(GOFMT) $(TOPLEVEL_PKG)

all: get build