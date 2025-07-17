# 🧑‍💼 Job Portal API

This is a RESTful API for a Job Portal system built with **Golang**, **Gin Framework**, **GORM**, and **PostgreSQL**. It supports user registration, authentication, job operations, and role-based access control (admin-only actions).

---

## 🚀 Features

- User Registration & Login with JWT Authentication
- Password hashing with bcrypt
- Admin-only job creation, update, and deletion
- Public job listing and detail fetching
- Middleware for authentication and authorization
- PostgreSQL as the database with GORM ORM

---

## 🧱 Tech Stack

- **Language**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Environment Config**: dotenv

---
## 🔐 Create .env File
PORT=8000.  
DATABASE_URL=postgres://user:password@localhost:5432/job_portal_db.  
JWT_SECRET=your_jwt_secret_key.  

---

## Install Dependencies

go mod tidy

---
## Run the Server
go run ./cmd/main.go
