package main

import (
		"io"
		"log"
		"net/http"
    	"fmt"

)
	

func main (){

		http.HandleFunc("/hello", func(w http.ResponseWriter, r * http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "text/plain")
				io.WriteString(w, "Hello World !\n")
				go hello ("cuongnp2")
				
		})

		err := http.ListenAndServe(":8080", nil)

		if err !=nil {
				log.Fatalf("Clould not start server: %s\n", err.Error())
		}
}

func hello(param string){
	fmt.Printf (param)
}
