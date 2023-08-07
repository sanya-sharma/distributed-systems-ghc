# Use the official PostgreSQL image from Docker Hub
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=db

# Copy your SQL script into the container
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port (optional but useful for external connections)
EXPOSE 5432