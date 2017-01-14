
SRC := $(shell find . -name "plug*.go")
SHARED_LID := $(SRC:.go=.so)


all: $(SHARED_LID)

%.so : %.go
	go build -buildmode=plugin -o $@ $<

clean:
	rm -fr $(SHARED_LID)
