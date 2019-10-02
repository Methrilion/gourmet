FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-analytics /app
EXPOSE 9095
CMD ["./svc-analytics"]