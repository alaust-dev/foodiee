# Foodiee API

## Local Development

### Utilizing Docker
Experience seamless local deployment with Docker by simply utilizing the provided docker-compose file.
```bash
docker-compose up --build
```

### Dockerless Option
Prefer a faster setup without Docker? Follow these steps:

Setup prerequisites:
```bash
export DB_FILE_PATH="foodiee.db"
sqlite3 $DB_FILE_PATH < schema.sql
```

Launch the application:
```bash
go run cmd/foodiee_backend/main.go
```

## Swagger-UI
Explore the comprehensive API documentation conveniently via http://localhost:8080/swagger-ui
