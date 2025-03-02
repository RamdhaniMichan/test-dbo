# 🚀 Go Gin Boilerplate

This boilerplate is a template for building a **REST API** using **Go Gin**, featuring **JWT Authentication** and **RBAC (Role-Based Access Control)**. It includes **Docker** and **PostgreSQL** for easy deployment.

## 🛠️ Tech Stack

- **Go** (Gin Framework)
- **PostgreSQL** (Database)
- **GORM** (ORM for database operations)
- **JWT** (JSON Web Token for authentication)
- **Docker** (Containerization)

## 📌 Key Features

- ✅ **JWT Authentication** (Login & Register)
- ✅ **RBAC (Role-Based Access Control)**
- ✅ **CRUD for Users & Orders**
- ✅ **Middleware for authorization**
- ✅ **Data Preloading with GORM**
- ✅ **Dockerized Setup (App)**

## 📂 Project Structure

```
/go-gin-boilerplate
│── config/
│   ├── config.go  # Application configuration
│── controllers/
│   ├── auth_controller.go  # Register & Login
│   ├── order_controller.go  # Order API
│   ├── product_controller.go  # Product API
│   ├── user_controller.go  # User API
│── database/
│   ├── database.go  # PostgreSQL database configuration
│── middleware/
│   ├── jwt.go  # JWT Middleware
│   ├── logger.go  # Logger Middleware
│   ├── rbac.go  # Role-Based Access Control Middleware
│── models/
│   ├── order_item.go  # Order Item Model
│   ├── order.go  # Order Model
│   ├── product.go  # Product Model
│   ├── user.go  # User Model
│── repositories/
│   ├── auth_repository.go  # Auth Repository
│   ├── order_repository.go  # Order Repository
│   ├── product_repository.go  # Product Repository
│   ├── user_repository.go  # User Repository
│── routes/
│   ├── auth_route.go  # Auth Routes
│   ├── order_route.go  # Order Routes
│   ├── product_route.go  # Product Routes
│   ├── user_route.go  # User Routes
│── services/
│   ├── auth_service.go  # Auth Service
│   ├── order_service.go  # Order Service
│   ├── product_service.go  # Product Service
│   ├── user_service.go  # User Service
│── utils/
│   ├── jwt.go  # JWT Utility
│   ├── pagination.go  # Pagination Utility
│── main.go  # Application entry point
│── Dockerfile  # Container configuration
│── docker-compose.yml  # Service orchestration with database
```

## 🚀 How to Run

1. **Run with Docker**
```sh
docker build -t test-dbo .
docker run --env-file .env -p 8080:8080 test-dbo
```

2. **Run manually** (without Docker)
```sh
go mod tidy  # Install dependencies
go run main.go  # Start the application
```

## 📢 API Endpoints

### Authentication

| Method | Endpoint     | Description |
|--------|-------------|-------------|
| POST   | `/api/login`  | User login  |
| POST   | `/api/register` | User registration |

---

### Users

| Method | Endpoint      | Description |
|--------|--------------|-------------|
| GET    | `/api/users`  | Get all users (Admin only) |
| POST   | `/api/users`  | Create a new user (Admin only) |
| GET    | `/api/user/:id` | Get user by ID (Admin only) |
| DELETE | `/api/user/:id` | Delete user (Admin only) |
| PUT    | `/api/user/:id` | Update user (Admin only) |

---

### Products

| Method | Endpoint       | Description |
|--------|---------------|-------------|
| POST   | `/api/product`  | Create a new product (Admin only) |
| GET    | `/api/products` | Get all products |
| GET    | `/api/product/:id` | Get product by ID |
| PUT    | `/api/product/:id` | Update product (Admin only) |
| DELETE | `/api/product/:id` | Delete product (Admin only) |

---

### Orders

| Method | Endpoint      | Description |
|--------|--------------|-------------|
| POST   | `/api/order`  | Create a new order (Authenticated users) |
| GET    | `/api/orders` | Get all orders (Authenticated users) |
| GET    | `/api/order/:id` | Get order by ID (Authenticated users) |

## 🛠️ cURL Examples

### **🔐 Register**
```sh
curl -X POST http://localhost:8080/api/register \
     -H "Content-Type: application/json" \
     -d '{"name": "John Doe", "email": "johndoe@example.com", "password": "password123"}'
```

### **🔑 Login**
```sh
curl -X POST http://localhost:8080/api/login \
     -H "Content-Type: application/json" \
     -d '{"email": "johndoe@example.com", "password": "password123"}'
```

### **📦 Fetch Orders (JWT Required)**
```sh
curl -X GET http://localhost:8080/api/order \
     -H "Authorization: Bearer <TOKEN_JWT>"
```

## 🔑 Environment Variables (Optional)
If running without Docker, create a `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=mydatabase
JWT_SECRET=mysecretkey
```

## 📊 Entity Relationship Diagram (ERD)
The following is the ERD of the database used in this project:

![alt text](https://res.cloudinary.com/dwckpepep/image/upload/v1740932595/ERD_test_dbo_xl6oko.png)

To see more ERDs, visit the following link: [dbdiagram.io](https://dbdiagram.io/d/ERD-test-dbo-67c482dc263d6cf9a0f3563d)

## 🎯 TODO (Future Enhancements)
- 🔄 **Refresh Token Implementation**
- 📝 **Logging & Error Handling Improvements**
- 📊 **Monitoring with Prometheus**

---
Happy coding! 🚀

