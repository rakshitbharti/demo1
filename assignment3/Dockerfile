FROM golang:1.17-apline3.15

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .


RUN  go mod download
ENV PORT 8111
EXPOSE 8111

RUN  go build

CMD ["./assignment3"]
