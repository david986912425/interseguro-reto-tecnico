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

## 6. Endpoints (Curl)

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "dm7659746@gmail.com",
    "password": "123456"
}'
```

```bash
curl --location 'http://localhost:8080/api/matrix' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '[
  [1, 2],
  [3, 4],
  [5, 6]
]'
```

```bash
curl --location 'http://localhost:3000/api/matrix/analyze' \
--header 'Content-Type: application/json' \
--data '{
    "q": [
        [
            -0.16903085094570325,
            0.8970852271450604,
            0.4082482904638626
        ],
        [
            -0.50709255283711,
            0.2760262237369414,
            -0.816496580927726
        ],
        [
            -0.8451542547285166,
            -0.34503277967117696,
            0.4082482904638632
        ]
    ],
    "r": [
        [
            -5.916079783099615,
            -7.437357441610946
        ],
        [
            0,
            0.8280786712108259
        ],
        [
            0,
            0
        ]
    ]
}'
```