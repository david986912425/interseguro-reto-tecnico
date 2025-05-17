package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-backend/types"
	"go-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func connectMongo() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "MONGO_URI no definido en el entorno")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client, err
}

func Login(c *fiber.Ctx) error {
	var input types.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Formato inválido")
	}

	client, err := connectMongo()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error de conexión a la base de datos")
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("intersegurodb").Collection("users")

	var user types.User
	err = collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Credenciales incorrectas")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Credenciales incorrectas")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error al generar el token")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
