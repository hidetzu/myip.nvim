TARGET:=myip

build:
	go build -o $(TARGET)

install: build
	mv $(TARGET) $(GOPATH)/bin

clean:
	rm -rf $(GOPATH)/bin/$(TARGET)

test:
	nvim --cmd "set rtp+=./"

