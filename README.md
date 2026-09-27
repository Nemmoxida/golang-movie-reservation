# Golang Movie Reservation API

A Go-based movie reservation backend built with Gin and PostgreSQL. The application supports user authentication, movie management, seat reservation flows, and Midtrans payment initiation.

## Overview

This project exposes a REST API for a movie ticket booking system with two main user roles:

- Admin: manages movies and schedules
- Client: signs up, logs in, reserves seats, and cancels reservations

The server runs on port 3001.

## Tech Stack

- Go
- Gin Gonic
- PostgreSQL
- pgx
- JWT (Golang JWT)
- bcrypt
- Midtrans Snap
- dotenv

## Project Structure

```text
.
├── cmd/
│   └── golang-movie-reservation/
│       └── main.go
├── database/
│   └── database.go
├── internals/
│   ├── router/
│   │   └── router.go
│   └── services/
│       ├── admin.go
│       ├── cllient.go
│       ├── login.go
│       ├── movie.go
│       ├── payment.go
│       └── signup.go
├── middleware/
│   └── auth.go
├── pkg/
│   ├── extrackClaim.go
│   └── verifyToken.go
├── .env
├── database.sql
├── go.mod
├── go.sum
├── req.txt
└── README.md
```

## Prerequisites

- Go 1.20+
- PostgreSQL database
- A Midtrans account for payment integration
- Environment variables configured in a `.env` file

## Environment Variables

Create a `.env` file in the project root with values similar to:

```env
JWT_SECRET=your_jwt_secret
DB_LINK=postgres://username:password@localhost:5432/movie_reservation
MIDTRANS_CLIENT=your_midtrans_client_key
MIDTRANS_SERVER=your_midtrans_server_key
```

The project uses `godotenv` to load these values when the app starts.

## Database Setup

1. Create a PostgreSQL database.
2. Import the schema from `database.sql`.
3. Ensure the `DB_LINK` value matches your local Postgres credentials and database name.

The SQL file includes:

- user roles enum
- movie table
- showtimes table
- seats table
- basic constraints and relationships

## Running the API

Install dependencies:

```bash
go mod download
```

Start the server:

```bash
go run ./cmd/golang-movie-reservation
```

The application will listen on:

```text
http://localhost:3001
```

## API Endpoints

### Public Endpoints

| Method | Route       | Description                       |
| ------ | ----------- | --------------------------------- |
| POST   | `/signup`   | Register a new client account     |
| POST   | `/login`    | Login and receive a JWT token     |
| GET    | `/getmovie` | Fetch all movies                  |
| POST   | `/payment`  | Create a Midtrans payment session |

### Admin Endpoints

All admin routes require an authorization token and admin role.

| Method | Route             | Description       |
| ------ | ----------------- | ----------------- |
| POST   | `/addmovie`       | Add a movie       |
| PUT    | `/editmovie`      | Edit a movie      |
| DELETE | `/deletemovie`    | Delete a movie    |
| POST   | `/addschedule`    | Add a showtime    |
| PUT    | `/editschedule`   | Edit a showtime   |
| DELETE | `/deleteschedule` | Delete a showtime |

### Client Endpoints

All client routes require an authorization token.

| Method | Route             | Description               |
| ------ | ----------------- | ------------------------- |
| POST   | `/reserve`        | Reserve one or more seats |
| POST   | `/reserve/cancel` | Cancel a reservation      |

## Authentication

The app uses JWT authentication.

Send the token in the request header:

```http
Authorization: Bearer <token>
```

The middleware checks the token and validates the user role before allowing access to protected routes.

## Example Requests

### Signup

```http
POST /signup
Content-Type: application/json

{
  "username": "alice",
  "password": "secret123"
}
```

### Login

```http
POST /login
Content-Type: application/json

{
  "username": "alice",
  "password": "secret123"
}
```

Response example:

```json
{
  "status": "success",
  "token": "<jwt_token>"
}
```

### Add Movie (Admin)

```http
POST /addmovie
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "name": "Inception",
  "creator": "Christopher Nolan",
  "genre": ["Sci-Fi", "Thriller"],
  "release_date": "2026-09-27T00:00:00Z",
  "desc_movie": "A mind-bending thriller.",
  "poster_image_url": "https://example.com/inception.jpg",
  "ticket_price": 45000
}
```

### Reserve Seats (Client)

```http
POST /reserve
Authorization: Bearer <client_token>
Content-Type: application/json

{
  "showtime_id": "8f193a18-7d41-4e46-a61f-7f9b8317d0a7",
  "seats": ["A1", "A2"],
  "user_id": "0a3d1af1-7d4f-4ed8-9d3d-7df5d7d58dd1"
}
```

## Notes and Caveats

- The current project is a backend-only API there is yet a frontend included.
- Authentication and authorization are implemented in middleware and JWT claims.
- Payment creation uses Midtrans Snap and returns a redirect token for frontend payment flow.

## Future Improvement Ideas

- Add validation and sanitization for all request payloads
- Add explicit admin/client user creation flows
- Add endpoints to retrieve schedules and seat availability
- Improve error handling consistency across routes
- Add unit/integration tests for auth and booking logic
- Align route definitions with actual delete semantics and frontend requirements
