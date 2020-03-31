# Base container for run service
FROM alpine

# Define service name
ARG SVC=uluru-api

# Go to workdir
WORKDIR /src/uluru-api

# Copy binaries
COPY bin/uluru-api /usr/bin/uluru-api
COPY bin/goose /usr/bin/goose
COPY bin/wait-db /usr/bin/wait-db

# Copy all database migrations
COPY database/migrations/* /src/uluru-api/migrations/

# Expose service port
EXPOSE 5000

# Run service
CMD ["/bin/sh", "-l", "-c", "wait-db && cd /src/$SVC/migrations/ && goose postgres ${DATABASE_URL} up && uluru-api"]