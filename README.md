# btc-email

This is a sample Gin application that demonstrates various functionalities, including retrieving exchange rate, subscribing to emails, and sending emails. The application will be run inside a Docker container.

## Getting Started

To get started with the application in Docker, follow the instructions below.

### Prerequisites

- Docker

### Installation

1. Clone the repository:

```shell
git clone <repository-url>
```

2. Change into the project directory:

```shell
cd btc-email
```

### Configuration

Before running the application in a Docker container, make sure to set the following environment variables:

- `EMAIL_SENDER`: The email address from which you want to send the emails.
- `SMTP_HOST`: The SMTP server host.
- `SMTP_PORT`: The SMTP server port.
- `SMTP_USER`: The SMTP server username.
- `SMTP_PASSWORD`: The SMTP server password.

### Build and Run with Docker

1. Build the Docker image:

```shell
docker build -t btc-email .
```

2. Run the Docker container:

```shell
docker run -p 8080:8080 --env-file <path-to-env-file> btc-email
```

Make sure to replace `<path-to-env-file>` with the path to a file containing the environment variable configurations.

The application will start inside the Docker container and listen on `http://localhost:8080`.

#### API Endpoints

- **GET /rate**: Retrieves the current exchange rate of Bitcoin to UAH.
- **POST /subscribe**: Subscribes an email address to receive notifications.
- **POST /sendEmails**: Sends emails to all subscribed email addresses with the latest exchange rate.

## License

This project is licensed under the [MIT License](LICENSE).