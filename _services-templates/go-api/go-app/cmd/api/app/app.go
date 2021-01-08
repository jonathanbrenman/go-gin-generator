package app

import (
	"go-app/cmd/api/router"
)

func Start() {

	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}
