package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		/*
		*http.response & bool
		 */
		flusher, ok := w.(http.Flusher) // type assertion

		if !ok {
			http.Error(w, "Streaming not supported!", http.StatusInternalServerError)
			return
		}

		for i := 1; i <= 10; i++ {
			fmt.Fprintf(w, "number: %d\n", i)
			flusher.Flush()         // Forçando a escrita dos dados para o cliente > Flush sends any buffered data to the client.
			time.Sleep(time.Second) // Pausa de 1 segundo entre os números
		}
	})

	http.ListenAndServe(":8080", nil)
}
