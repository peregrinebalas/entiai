package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

func main() {
	os.Setenv("DATABASE_URL", "postgresql://localhost:5432")

	// for _, e := range os.Environ() {
	// 	pair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(pair[0])
	// }

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// var greeting string
	// err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)

	// s := &http.Server{
	// 	Addr: ":8080",
	// 	// Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// log.Fatal(s.ListenAndServe())

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)

	ctx, _ := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	log.Fatal(serverOne.ListenAndServe())
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request to / received\n")
	io.WriteString(w, "This is my website!\n")
}
