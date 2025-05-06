package main

import (
	"net/http"

	"github.com/luisteixeira74/kubernetes-cicd-deploy-simulation/app"
)

func main() {
	http.HandleFunc("/", app.RootHandler)
	http.ListenAndServe(":8080", nil)
}
