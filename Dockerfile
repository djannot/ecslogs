FROM google/golang
WORKDIR /go/src
RUN git clone https://github.com/djannot/ecslogs.git
WORKDIR /go/src/ecslogs
RUN go get "golang.org/x/crypto/ssh"
RUN go build .
