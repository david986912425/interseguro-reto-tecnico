# Reto Técnico Interseguro - Backend

## 🚀 Tecnologías utilizadas

- Docker
- Go
- Express

## 1. 📦 Instalación

```bash
git clone git@github.com:david986912425/interseguro-reto-tecnico.git
```

## 2. 🗝️ Modificar los archivos .env
go-backend/.env
```bash
MONGO_URI=mongodb://david:secret@mongo:27017/intersegurodb
JWT_KEY=jsonwebtoken
EXPRESS_URL=http://express-backend:3000
```

express-backend/.env
```bash
MONGO_URI=mongodb://david:secret@mongo:27017/intersegurodb
JWT_KEY=jsonwebtoken
```

## 3.Construir y levantar los contenedores
```bash
docker-compose up --build
```

## 4. ✅ Verificar que todo esté corriendo
Express: http://localhost:3000

Go: http://localhost:8080

## 5. ✅ Conectarte a Mongo
![alt text](image.png)