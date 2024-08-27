FROM golang:1.23 as builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .

RUN build/build.sh



FROM alpine

ENV DB_DSN=""
ENV REDIS_DSN=""

RUN apk add tzdata
WORKDIR /
COPY --from=builder /app/build/output .
EXPOSE 443
EXPOSE 80
CMD ["/run.sh"]
