package main

import (
	"TugasMID2/configs"
	"TugasMID2/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))

}
