include ../Makefile.defs


TARGET := vellun-daemon

.PHONY: all $(TARGET)

all: $(TARGET)

$(TARGET):
	$(GO_BUILD) -o $@

clean:
	@rm -f $(TARGET)
