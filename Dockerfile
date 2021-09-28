# start golang from bse image
FROM golang:alpine As builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-chache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# validate dependencies and install it if needed
RUN go mod tidy

# 
RUN go build -o binary

# Expose port 8080 to the outside world
EXPOSE 8080