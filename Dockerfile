FROM golang
ADD . /go/src/agro
RUN go get -d -v agro
RUN go install agro
EXPOSE 8082
