FROM postgres:15-alpine

ENV POSTGRES_PASSWORD=pismo4321
ENV POSTGRES_DB=pismo
ENV POSTGRES_USER=pismo


WORKDIR /app

COPY transaction-service .
COPY app.yaml .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 80 5432

ENTRYPOINT ["./entrypoint.sh"]