FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-statistics /app
EXPOSE 9094
CMD ["./svc-statistics"]