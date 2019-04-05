FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-storage-writer /app
EXPOSE 9091
CMD ["./svc-storage-writer"]