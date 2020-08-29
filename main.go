package main

import(
    "fmt"
    "os"
    "log"
	"github.com/joho/godotenv"
	recipes "golang-stripe-recipe/pkg/recipes"
)

// LoadEnv loading .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
    LoadEnv()
    fmt.Println("test")
    client := recipes.NewClient(os.Getenv("STRIPE_SECRET_KEY"))
    client.Test()
}
