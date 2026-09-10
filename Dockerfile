FROM golang:1.26 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest

LABEL maintainer="asami, ahmsalem" \
      description="Ascii-Art-Web - converts text to Ascii art banners" \
      version="1.0"

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/formats ./formats
COPY --from=builder /app/style ./style  

EXPOSE 8080

CMD ["./server"]