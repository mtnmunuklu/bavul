# Use the official Golang image as the base image
FROM golang

# Define an environment variable for the app directory
ENV APP=/app

# Set up Services
# Create a directory for services
WORKDIR ${APP}

# Copy all services to the app directory
COPY authsvc authsvc
COPY vulnsvc vulnsvc
COPY apisvc apisvc

# Expose ports for the services
# These ports should match the ports that your services are listening on
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002

# Define the command to run the services
# Replace the command with the correct paths to your service binaries
CMD ["/app/authsvc", "/app/vulnsvc", "/app/apisvc"]