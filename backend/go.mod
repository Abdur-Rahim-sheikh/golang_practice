module ecommerce

go 1.26.5

require github.com/joho/godotenv v1.5.1 // direct

require github.com/golang-jwt/jwt/v5 v5.3.1 // direct

require (
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.12.3
	github.com/rubenv/sql-migrate v1.8.1
)

require github.com/go-gorp/gorp/v3 v3.1.0 // indirect
