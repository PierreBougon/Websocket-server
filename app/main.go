package app

import (
	"fmt"
	"github.com/PierreBougon/Websocket-server/app/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println(port)

	// API
	api := router.PathPrefix("/api").Subrouter()

	//Connect Websocket
	api.HandleFunc("/ws", controllers.ConnectWebSocket).Methods("GET")

	//Launch the app, visit localhost:443/api
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
