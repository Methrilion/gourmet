FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-storage-reader /app
EXPOSE 9093
CMD ["./svc-storage-reader"]