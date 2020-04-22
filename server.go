package main // import "github.com/serverwentdown/leet"

//go:generate protoc -I framebuffer framebuffer/framebuffer.proto --go_out=plugins=grpc:framebuffer

import (
	"context"
	"log"
	"time"

	"github.com/serverwentdown/leet/framebuffer"
)

type Server struct {
	drawer *Drawer
}

func NewServer(drawer *Drawer) *Server {
	s := &Server{
		drawer: drawer,
	}
	return s
}

func (s *Server) DrawFrame(ctx context.Context, in *framebuffer.FrameBuffer) (*framebuffer.DrawResponse, error) {
	s.drawer.SetLayer(int32(in.Layer), in.Frame.Dots)
	timeDrawStart := time.Now()
	err := s.drawer.Draw()
	timeDrawEnd := time.Now()
	if err != nil {
		return nil, err
	}
	log.Printf("Draw took %d milliseconds", timeDrawEnd.Sub(timeDrawStart)/1000/1000)
	return &framebuffer.DrawResponse{}, nil
}

func (s *Server) DrawFrames(ctx context.Context, in *framebuffer.FrameSequence) (*framebuffer.DrawResponse, error) {
	log.Print("Received FrameSequence, but not implemented")
	return &framebuffer.DrawResponse{}, nil
}

func (s *Server) GetLength(ctx context.Context, in *framebuffer.LengthRequest) (*framebuffer.LengthResponse, error) {
	length := uint32(s.drawer.Length)
	return &framebuffer.LengthResponse{
		Length: length,
	}, nil
}
