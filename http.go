package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var builder strings.Builder
	builder.Grow(1048577)
	for i := 0; i < 1048577; i++ {
		builder.WriteByte(0)
	}
	beefcake := builder.String()

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		fmt.Fprintf(w, "I'm %s\n beefcake:%s", hostname, beefcake)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
