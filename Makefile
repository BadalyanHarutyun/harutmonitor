.PHONY: install

BINARY_NAME=harutmonitor
BINARY_URL=https://raw.githubusercontent.com/BadalyanHarutyun/harutmonitor/main/binaries/linux/$(BINARY_NAME)

install:
	wget -O $(BINARY_NAME) $(BINARY_URL)
	./install.sh
