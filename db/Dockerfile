FROM postgres:latest

ENV POSTGRES_PASSWORD docker

COPY *.sql /docker-entrypoint-initdb.d/
