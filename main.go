package main

import (
	"TugasMID2/configs"
	"TugasMID2/routes"
	"os"
)

func main() {
	configs.InitDB()

	// e := routes.New()
	// e.Logger.Fatal(e.Start(":8000"))
	e := routes.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start("0.0.0.0:" + port))

}
