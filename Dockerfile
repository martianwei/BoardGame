# Use an official Go runtime as the base image
FROM golang:1.20.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Install make
RUN apk add --no-cache make

# Copy the source code from the host to the container
COPY . .

# Set the command to run when the container starts
CMD ["make","run-boardgame"]