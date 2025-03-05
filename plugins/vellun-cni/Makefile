include ../../Makefile.defs

TARGET := vellun-cni

.PHONY: all $(TARGET)

all: $(TARGET)

$(TARGET):
	$(GO_BUILD) -o $@

clean:
	@rm -f $(TARGET)
