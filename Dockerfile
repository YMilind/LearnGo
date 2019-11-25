FROM golang
COPY ./web ./
WORKDIR /go/src 
CMD ./wiki
#ENTRYPOINT ./wiki

