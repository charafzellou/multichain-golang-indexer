# Use the official Golang image as the base image
FROM golang:1.19.4-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the local source files to the container
COPY . .

# Expose the app's port to the host
EXPOSE 3200

# Set the command to run the application when the container starts
CMD ["go", "run", "."]