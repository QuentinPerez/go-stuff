FROM library/golang:1.8

RUN apt-get update && \
    apt-get upgrade -y && \
	apt-get install -y build-essential

# go1.8rc1: plugin doesn't work, hence we download the sources from tip
RUN cd /tmp && \
    git clone https://go.googlesource.com/go && \
	cd go/src && \
	GOROOT_BOOTSTRAP=/usr/local/go/ ./make.bash && \
	cd .. && \
    cp bin/go /usr/local/go/bin/go


