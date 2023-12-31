
	FROM golang:1.20-alpine

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

EXPOSE 8080

# create executable
RUN go build -o fastprint

# menjalankan file executablenya
CMD ["./fastprint"]
	