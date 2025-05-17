FROM mongo:latest

ENV MONGO_INITDB_ROOT_USERNAME=root
ENV MONGO_INITDB_ROOT_PASSWORD=secret
ENV MONGO_INITDB_DATABASE=intersegurodb


ADD init-mongo.js /docker-entrypoint-initdb.d/