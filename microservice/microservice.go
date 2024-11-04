package microservice

import (
	"log"
	"net/http"
)

func Start(server *http.Server) error {

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("error with ListenAndServe: %v", err)
		}

	}()

	log.Printf("server started at %v\n", server)

	select {}

}
