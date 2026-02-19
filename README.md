# Arbolito API

Arbolito is a Go-based API designed to aggregate and provide the average "Dolar Blue" rates in Argentina from multiple sources. It serves as a unified interface to access financial data regarding the unofficial dollar exchange rate.

## Purpose

The main purpose of this API is to:
- Fetch multiple dollar rates (Blue, Oficial, MEP, CCL, Tarjeta) from various external APIs (DolarAPI, Bluelytics, Criptoya).
- Calculate an average rate from these sources.
- Provide a reliable and cached response to minimize latency and external dependency usage.

## Tech Stack

- **Language**: [Go](https://go.dev/) (Golang)
- **Database**: [MongoDB](https://www.mongodb.com/) (for caching and storage)
- **Documentation**: [Swagger](https://swagger.io/) (via `swaggo/http-swagger`)
- **External Libraries**:
  - `github.com/joho/godotenv`: For loading environment variables.
  - `github.com/stretchr/testify`: For testing.
  - `go.mongodb.org/mongo-driver`: MongoDB driver for Go.

## Endpoints

### Health Check
- **URL**: `/health`
- **Method**: `GET`
- **Description**: Returns a simple "OK" to indicate the service is running.

### Get Dollar Rates
- **URLs**: `/dolar-blue`, `/dolar-oficial`, `/dolar-mep`, `/dolar-ccl`, `/dolar-tarjeta`
- **Method**: `GET`
- **Description**: Returns the calculated average dollar rate (buy and sell) for the specified quote type, aggregating configured sources that provide it.

### Swagger Documentation
- **URL**: `/swagger/index.html`
- **Method**: `GET`
- **Description**: Interactive API documentation.

## Configuration

The application uses an `.env` file for configuration. The following environment variables are supported:

| Variable | Description | Default Value |
|----------|-------------|---------------|
| `SERVER_PORT` | Port where the API will listen | `8080` |
| `MONGO_URI` | MongoDB connection URI | `mongodb://localhost:27017` |
| `MONGO_DB_NAME` | Name of the MongoDB database | `arbolito` |
| `DOLAR_API_URL` | URL for DolarAPI source | `https://dolarapi.com/v1/dolares` |
| `BLUELYTICS_API_URL` | URL for Bluelytics source | `https://api.bluelytics.com.ar/v2/latest` |
| `CRIPTOYA_API_URL` | URL for Criptoya source | `https://criptoya.com/api/dolar` |
| `MONGO_USER` | MongoDB username (optional) | `admin` |
| `MONGO_PASSWORD` | MongoDB password (optional) | `password` |

## Running the Project

### Prerequisites
- Go 1.25 or higher
- MongoDB instance running

### Steps

1. **Clone the repository**:
   ```sh
   git clone <repository-url>
   cd arbolito
   ```

2. **Set up environment variables**:
   Create a `.env` file in the root directory or use the defaults.

3. **Run the application**:
   ```sh
   go run main.go
   ```

4. **Access the API**:
   Open your browser or an API client and navigate to `http://localhost:8080/swagger/index.html` to explore the endpoints.
