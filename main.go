package main // import "github.com/serverwentdown/leet"

import (
	"log"
	"net"

	"github.com/serverwentdown/leet/framebuffer"
	"google.golang.org/grpc"
)

const listen = ":5000"
const length = 287

func main() {
	log.Print("starting server")

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	drawer, err := NewDrawer(287)
	if err != nil {
		log.Fatalf("failed to setup WS281x library: %v", err)
	}

	framebuffer.RegisterDrawerServer(s, NewServer(drawer))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
