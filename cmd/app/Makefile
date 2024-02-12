GO=go 
DESTDIR=../../../
DESTNAME=app.o
TARGET=$(DESTDIR)$(DESTNAME)
LDFLAGS = -s -w -extldflags -static
GOFLAGS = -trimpath -tags release -buildmode exe

build:
	$(GO) env -w CGO_ENABLED=1 
	# -ldflags "$(LDFLAGS)"
	$(GO) build $(GOFLAGS) -o $(TARGET) 

run: build
	echo "cd $(DESTDIR); ./$(DESTNAME)" | bash