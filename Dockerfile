
FROM postgres:15-alpine

# Set environment variables for Postgres
ENV POSTGRES_PASSWORD=pismo4321
ENV POSTGRES_DB=pismo
ENV POSTGRES_USER=pismo


WORKDIR /app

# Copy the Go binary from builder
COPY transaction-service .
COPY app.yaml .

# Copy and prepare our startup script
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

# Expose Go app port and Postgres port
EXPOSE 80 5432

# Use our script as the entrypoint
ENTRYPOINT ["./entrypoint.sh"]