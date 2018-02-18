FROM gobuilder
WORKDIR /go/src/app
COPY . .
RUN dep ensure && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app . && echo "Building done !"

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app .
CMD ["./app"]  

