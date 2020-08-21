# Setup for ssh onto github
FROM xbuilder as builder
RUN echo 'export PATH=$PATH:/root/go/bin' >> /root/.bashrc
RUN echo 'source  /root/.bashrc'
ENV PATH="/root/go/bin:${PATH}"

WORKDIR /gwd/
RUN mkdir -p /repo
COPY . .
ENV ANNOTATIONS="third-party/googleapis"
RUN gnostic --grpc-out=. bookstore.yaml
RUN protoc --proto_path=. --proto_path=${ANNOTATIONS} --go_out=plugins=grpc:bookstore bookstore.proto
RUN protoc --proto_path=. --proto_path=${ANNOTATIONS} --grpc-gateway_out=bookstore bookstore.proto
RUN protoc --proto_path=. --proto_path=${ANNOTATIONS} --swagger_out=bookstore bookstore.proto
ENV GO111MODULE=on
RUN go mod init gwd
RUN go mod tidy
#RUN go build -a -o twillo .
#ENTRYPOINT ["./twillo"]

FROM  twillodotio/depmanger as builder2
RUN mkdir -p /repo
COPY --from=builder /gwd /go/src/github.com/twillo/gwd
WORKDIR /go/src/github.com/twillo/gwd/
ENV GO111MODULE=on
RUN GOOS=linux cgo_enabled=0 go build -a -o twillo .
ENTRYPOINT ["./twillo"]


FROM alpine:edge

RUN addgroup -S 1024
RUN adduser root 1024

# RUN apk add --update nodejs npm
# RUN apk add --update npm
# RUN npm install --global yo

# RUN npm install -g yo generator-xi5s
# RUN sed -i -e '/rootCheck/d' '/usr/lib/node_modules/yo/lib/cli.js'

RUN apk add git
RUN apk add --no-cache \
  openssh-client

COPY --from=builder2 /go/src/github.com/twillo/gwd/twillo /root/twillo
COPY --from=builder2 /go/src/github.com/twillo/gwd/start.sh /root/start.sh


CMD ["sh", "start.sh"]
