# Captcha Service

[![Go](https://img.shields.io/badge/Go-1.20%2B-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Overview

The Captcha Service is a Go-based microservice designed to generate and verify CAPTCHAs (Completely Automated Public Turing test to tell Computers and Humans Apart). It utilizes the Steambap algorithm for captcha generation and provides a simple API for integration into web applications.

## Features

*   **Captcha Generation:** Generates image-based CAPTCHAs using the Steambap algorithm.
*   **Captcha Verification:** Verifies user-submitted CAPTCHA codes.
*   **Customizable Captcha:** Allows customization of captcha attributes like width, height, length, and noise level.
*   **Structured Logging:** Employs `zap` for structured logging, making it easy to analyze and debug.
*   **Redis Integration:** Uses Redis for storing and managing CAPTCHA data.
*   **Middleware Logging:** Provides middleware for logging HTTP requests and responses.
*   **Configuration Management:** Supports configuration via environment variables or a configuration file.

## Technologies Used

*   **Go:** Programming language.
*   **zap:** Structured logging library.
*   **go-redis/redis:** Redis client for Go.
*   **Echo:** High performance, minimalist Go web framework.
* **Testify:** Test suite for Go.

## Getting Started

### Prerequisites

*   **Go:** Version 1.20 or higher.
*   **Redis:** A running Redis server.

### Installation

1.  **Clone the repository:** 
2.  **Install dependencies:** `go mod tidy`
3.  **Configure the application:**

    *   Create a `.env` file in the root directory (or use environment variables directly).
    *   Set the following environment variables:
        *   `UNDER_MAINTENANCE`: set `true` or `false`, `true` for under maintenance.
        *   `HTTP_SERVER_TIMEOUT`: set time out server (e.g., `10`).
        *   `HTTP_PORT`: port access service (e.g., `8080`).
        *   `JWT_KEY`: Jwt key to create encryption.
        *   `REDIS_ADDRESS`: Address of the Redis server (e.g., `localhost:6379`).
        *   `REDIS_PASSWORD`: Password for Redis (if any).
        *   `REDIS_DB`: Redis database number.
        *   `REDIS_MAX_IDLE`: Redis maximum open connection.
        *   `REDIS_MAX_RETRIES`: Redis maximum retries connection.
    * Example `.env` file:

4.  **Run the application:**
go run main.go

## Usage

### API Endpoints

*   **`POST /api/v1/captcha/generate`**
    *   **Description:** Generates a new CAPTCHA.
    *   **Request Body:**
        json { "width": 150, "height": 80, "length": 4, "noise": 1.0 }
    *   `width`: Width of the CAPTCHA image (optional, default: 0).
    *   `height`: Height of the CAPTCHA image (optional, default: 0).
    *   `length`: Length of the CAPTCHA code (optional, default: 0).
    *   `noise`: Noise level (optional, default: 0).

*   **Response Body:**
    json { "captcha_id": "unique-captcha-id", "captcha_image": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD...", "captcha_expired_time": "10m0s" }
*   `captcha_id`: Unique identifier for the CAPTCHA.
    *   `captcha_image`: Base64 encoded CAPTCHA image.
    * `captcha_expired_time`: Captcha expired time.

## Code Structure

*   **`app/`:** Application-specific code.
    *   **`config/`:** Configuration management.
    *   **`db/`:** Database-related code (Redis).
    *   **`logger/`:** Logging setup.
    *   **`middleware/`:** HTTP middleware.
    *   **`models/`:** Data models.
    *   **`server/`:** Server-related code.
    *   **`utils/`:** Utility functions.
*   **`pkg/`:** Reusable packages.
    *   **`api/`:** API-related code.
*   **`main.go`:** Main application entry point.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

This project is licensed under the MIT License.

## Future Improvements

*   **Authorization** Adding authorization to access endpoint.
*   **Connection RPC** Adding connection RPC.
*   **More Captcha Algorithms:** Add support for other captcha algorithms.
*   **Rate Limiting:** Implement rate limiting to prevent abuse.
*   **API Documentation:** Generate comprehensive API documentation (e.g., using Swagger).
*   **More Robust Error Handling:** Add more specific error types.
* **Dependency Injection:** Implement dependency injection.
* **More Test:** Add more test case.

## Contact

If you have any questions or suggestions, please contact miniwormtail@gmail.com.