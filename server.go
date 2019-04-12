package main // import "github.com/serverwentdown/leet"

//go:generate protoc -I framebuffer framebuffer/framebuffer.proto --go_out=plugins=grpc:framebuffer

import (
	"context"
	"log"

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
	s.drawer.SetLayerOrFill(int32(in.Layer), in.Frame.Dots, in.Frame.Fill)
	err := s.drawer.Draw()
	if err != nil {
		return nil, err
	}
	return &framebuffer.DrawResponse{}, nil
}

func (s *Server) DrawFrames(ctx context.Context, in *framebuffer.FrameSequence) (*framebuffer.DrawResponse, error) {
	log.Print("Received FrameSequence, but not implemented")
	return &framebuffer.DrawResponse{}, nil
}
