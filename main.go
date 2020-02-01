package main

import (
	"github.com/k-pozhidaev/mulciber.git/pkg/web"
	"log"
	"net/http"
)

func main() {


	web.InitContext(9096)

	web.CreateHandlers()

	log.Fatal(http.ListenAndServe(web.GetStringPort(), nil))

}

