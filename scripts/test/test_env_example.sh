export TEST_PHOENIX_PORT=3002
export TEST_PHOENIX_PROTOCOL=http
export TEST_PHOENIX_JWT_KEY=123456789
export TEST_PHOENIX_PATH=$GOPATH/github.com/jtorz/phoenix-backend
export TEST_PHOENIX_LOGGING_LEVEL=debug
export TEST_PHOENIX_DB_MAIN_CONNECTION="host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
export TEST_PHOENIX_REDIS_ADDRESS=127.0.0.1:6379
export TEST_PHOENIX_REDIS_PASSWORD=