package main // import "github.com/serverwentdown/leet"

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Print("starting server")
	_ := grpc.NewServer()
}
