package main

import (
	"os"

	"github.com/mjmhtjain/stackit/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
