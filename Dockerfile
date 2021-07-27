# Start from the latest alpine based golang base image
FROM golang:alpine as builder

# Build stamp argument
ARG buildstamp

# Git Commit Id
ARG gitCommitId

# Git Primary Branch
ARG gitPrimaryBranch

# Git Repository
ARG gitRepository

# Git username
ARG gitUsername

# Hostname
ARG hostname

# App Version
ARG appVersion

# Add maintainer info
LABEL maintainer="Shubham Sinha <sinhashubham95@gmail.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build Moxy
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/sinhashubham95/go-actuator/core.BuildStamp=$buildstamp -X github.com/sinhashubham95/go-actuator/core.GitCommitID=$gitCommitId -X github.com/sinhashubham95/go-actuator/core.GitPrimaryBranch=$gitPrimaryBranch -X github.com/sinhashubham95/go-actuator/core.GitURL=https://github.com/$gitRepository -X github.com/sinhashubham95/go-actuator/core.Username=$gitUsername -X github.com/sinhashubham95/go-actuator/core.HostName=$hostname  -X github.com/sinhashubham95/go-actuator/core.GitCommitTime=$buildstamp -X github.com/sinhashubham95/go-actuator/core.GitCommitAuthor=$gitUsername"

# Start again from scratch
FROM scratch

# Copy the binary
COPY --from=builder /app/moxy /moxy

# Set the working directory to data so all created files
# can be mapped to physical files on disk
WORKDIR /data

# Run binary
ENTRYPOINT ["../moxy", "--env=dev", "--name=moxy", "--version=$appVersion", "--port=$port", "--persistence-path=persistence.db"]
