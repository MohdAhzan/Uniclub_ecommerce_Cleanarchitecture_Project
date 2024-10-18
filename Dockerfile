FROM golang:1.23.2 AS builder
WORKDIR /uniclub_app
COPY . .
RUN go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/uniclub_project ./cmd
RUN go build -o ./cmd/uniclub_project ./cmd

FROM debian:bookworm-slim
COPY --from=builder /uniclub_app/cmd/uniclub_project  /uniclub_app/cmd/
COPY --from=builder  /uniclub_app/.env /uniclub_app/
COPY --from=builder  /uniclub_app/template/ /uniclub_app/template
WORKDIR /uniclub_app
CMD ["./cmd/uniclub_project"]

