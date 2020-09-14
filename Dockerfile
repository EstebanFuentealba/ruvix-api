# # Base container for run service
FROM alpine:latest

# Define service name
ARG SVC=uluru-api
ENV SVC=${SVC}

# Go to workdir
WORKDIR /src/${SVC}

# Copy binaries and json seed files
COPY bin/${SVC} /usr/bin/${SVC}
COPY misc/seeds/goals.json goals.json
COPY misc/seeds/institutions.json institutions.json
COPY misc/seeds/subscriptions.json subscriptions.json

# Run service
ENTRYPOINT ["uluru-api"]