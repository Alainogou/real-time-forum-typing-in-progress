#Use base image with golang latest version
FROM golang:1.21.1


LABEL AUTHORS="@alo - @alogou"

#Define the working directory
WORKDIR /REALTIMEFORUM

#Copy the go.mod and go.sum files into the container
COPY go.mod .
COPY go.sum .

#Download go dependencies
RUN go mod download

#Copy source code to container
COPY . .

#Compile the application
RUN go build -o REALTIMEFORUM

#Expose application listening port
EXPOSE 8081

#Application start command
CMD ["./REALTIMEFORUM"]