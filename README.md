# Go Microservice & Api Gateway Sample
- Microservices architecture implemented in Go.
- Utilizes gRPC for communication between services.
- Consists of 3 separate Microservices and 1 API Gateway.
- API Gateway manages incoming HTTP requests.
- gRPC used to route and forward requests to Microservices.
- Incorporates JWT authentication for enhanced security.


## Microservices with Docker Compose: Getting Started Guide

TStep-by-step instructions to set up and run a microservices architecture using Docker Compose. The architecture consists of three microservices—`auth-svc`, `order-svc`, and `product-svc`—communicating through gRPC, all managed by an API Gateway.

## Prerequisites

Ensure you have the following software installed on your system:

- Docker
- Docker Compose

## Setup

1. **Clone the Repository:**

   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```

2. **Build images and start service**
  
   ```bash
   docker-compose up -d
   ```

3. **Access the service**
  
   ## Base Url: http://localhost:3001 


    ## Auth
    ### 1. POST /auth/register
    - Description: Register a new user.
    - Handler function: `Register`

    ### 2. POST /auth/login
    - Description: Authenticate and log in a user.
    - Handler function: `Login`


    ## Product
    ### 1. POST /product/
    - Description: Create a new product.
    - Handler function: `CreateProduct`

    ### 2. GET /product/
    - Description: Retrieve a list of all products.
    - Handler function: `FindAll`

    ### 3. GET /product/:id
    - Description: Retrieve details of a specific product by its ID.
    - Handler function: `FindOne`


    ## Order
    ### 1. POST /order/
    - Description: Create a new order.
    - Handler function: `CreateOrder`





