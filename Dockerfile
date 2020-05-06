# Base container for run service
FROM scratch

# Define service name
ARG SVC=uluru-api
ENV SVC=${SVC}

# Go to workdir
WORKDIR /src/${SVC}

# Copy binaries
COPY bin/${SVC} /usr/bin/${SVC}

# Expose service port
EXPOSE 5000

# Run service
ENTRYPOINT ["uluru-api"]