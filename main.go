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
    client := recipes.NewClient(os.Getenv("STRIPE_SECRET_KEY"))
    client.Test()
    product_id := client.CreateProduct("test_product")
    plan_id := client.CreatePlan(2000, "test_plan", 0, product_id)
    fmt.Println(plan_id)
}
