FROM golang:1.16-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build
RUN CGO_ENABLED=0 GOARCH=amd6

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./banking"]