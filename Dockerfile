# Set the base image to use for subsequent instructions
FROM golang:1.22.2 AS build-stage

# Set the Current Working Directory inside the container
WORKDIR /app
# Now, the actual directory inside the container is /app.

# Copy go mod and sum files into /app 
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-example-executable ./cmd/my_app

# Multi-stage build

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

# Copy the executable from the build-stage to the distroless image
COPY --from=build-stage /golang-example-executable /golang-example-executable
COPY --from=build-stage /app/templates /templates

# Expose port 8080 to the outside world
EXPOSE 8080

 
USER nonroot:nonroot

ENTRYPOINT ["/golang-example-executable"]


# Build the Docker image using the following command
# docker build -t go-docker .

# Run the Docker container using the following command
# docker run -p 8080:8080 go-docker

# Open your web browser and navigate to http://localhost:8080 to see the output.

# To stop the container, press Ctrl + C in the terminal where the container is running.

# To remove the container, run the following command
# docker rm -f <container_id>

# To remove the image, run the following command
# docker rmi -f go-docker

# To list all the containers, run the following command
# docker ps -a

# To list all the images, run the following command
# docker images

# To list all the volumes, run the following command
# docker volume ls

# To list all the networks, run the following command
# docker network ls

# To list all the Docker logs, run the following command
# docker logs <container_id>
