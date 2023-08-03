# Use the official PostgreSQL image from Docker Hub
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydb

# Copy your SQL script into the container
COPY schema.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port (optional but useful for external connections)
EXPOSE 5432