package main

import (
	"net/http"
)

func addRoutes() {

	siteRouter.Handler(http.MethodGet, "/", contentHandler)

}
