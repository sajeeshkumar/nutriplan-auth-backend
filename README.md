# Golang Web API

This is a simple Golang web API that demonstrates how to build and deploy a Go application. It listens on a configurable port and responds with a basic message. This application is intended to be deployed on platforms like Render, but it can be used locally for development and testing.

## Table of Contents
- [Installation](#installation)
- [Running Locally](#running-locally)
- [Deploying to Render](#deploying-to-render)
- [Environment Variables](#environment-variables)
- [Docker Setup (Optional)](#docker-setup-optional)
- [License](#license)

## Installation

To get started with the Golang Web API, you'll need to have **Go 1.18+** installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/golang-web-api.git
   cd golang-web-api

2. Initialize Go modules
    ```bash
    go mod tidy

## Running Locally

1. Set environment variables
    ```makefile
    USERS=<username:password comma separated for different ones>

2. Build and run the application
    ```bash
    go run main.go