# PostBox

PostBox is an open-source, developer-friendly email-sending service that simplifies the process of sending batch emails using Go, Redis, PostgreSQL, and MongoDB. With integration to AWS SES, PostBox ensures reliable email delivery for various applications, making it a lightweight alternative to major players.

## Features

- Easy-to-use API for sending emails and batch processing.
- Integration with AWS SES for email delivery.
- Supports Redis for efficient queuing.
- Uses PostgreSQL and MongoDB for relational and JSON data storage.
- Containerized with Docker for smooth deployment.
- Designed for scalability with CI/CD integration (GitLab CI).

## Technology Stack

- **Backend**: Go
- **Queue**: Redis
- **Databases**: PostgreSQL (relational data), MongoDB (JSON records)
- **Email Provider**: AWS SES
- **Deployment**: Docker, GitLab CI/CD (Kubernetes and Jenkins support coming soon)

## Installation

### Prerequisites

- Go (1.18 or higher)
- Docker and Docker Compose
- Redis
- PostgreSQL
- MongoDB
- AWS (account)

### Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/PostBox.git
   cd PostBox
   ```

2. **Set up environment variables**  
   Create a `.env` file based on `.env.example` and fill in the necessary configuration details:
