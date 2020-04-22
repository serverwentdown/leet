package main // import "github.com/serverwentdown/leet"

import (
	"log"
	"time"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// TODO: full-sized grid with multiple channels mapped onto the grid

const (
	MAX_LAYERS = 10
)

type Drawer struct {
	Length int
	layers [][]uint32
	device *ws2811.WS2811
}

func NewDrawer(length int) (*Drawer, error) {
	opt := ws2811.DefaultOptions
	opt.Channels[0].LedCount = length
	opt.Channels[0].Brightness = 255

	d := &Drawer{}
	d.Length = length
	d.layers = make([][]uint32, MAX_LAYERS)
	for i := range d.layers {
		d.layers[i] = make([]uint32, length)
	}
	for i := range d.layers[0] {
		d.layers[0][i] = 0xff000000
	}
	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		return nil, err
	}
	err = dev.Init()
	if err != nil {
		return nil, err
	}
	d.device = dev

	return d, nil
}

func (d *Drawer) SetLayer(layer int32, dots []uint32) {
	for i, dot := range dots {
		if i > d.Length {
			break
		}
		d.layers[layer][i] = dot
	}
}

func (d *Drawer) Draw() error {
	timeMixStart := time.Now()
	for i := 0; i < d.Length; i++ {
		dot := mix(d.layers, i)
		d.device.Leds(0)[i] = dot
	}
	timeMixEnd := time.Now()

	if err := d.device.Render(); err != nil {
		return err
	}

	log.Printf("mix took %d milliseconds", timeMixEnd.Sub(timeMixStart)/1000/1000)
	return nil
}

func mix(layers [][]uint32, i int) uint32 {
	var base uint32
	for _, layer := range layers {
		if i >= len(layer) {
			continue
		}
		base = mixColors(layer[i], base)
	}
	return base
}

func mixColors(a uint32, b uint32) uint32 {
	// Extract channels
	aa, ba := uint32(a>>24), uint32(b>>24)
	ar, br := a&uint32(0x00FF0000)>>16, b&uint32(0x00FF0000)>>16
	ag, bg := a&uint32(0x0000FF00)>>8, b&uint32(0x0000FF00)>>8
	ab, bb := a&uint32(0x000000FF), b&uint32(0x000000FF)
	// Apply alpha computation to each channel
	oa := uint32(ba*(255-aa)/255 + aa)
	or := uint32(br*ba*(255-aa)/(255*255) + ar*aa/255)
	og := uint32(bg*ba*(255-aa)/(255*255) + ag*aa/255)
	ob := uint32(bb*ba*(255-aa)/(255*255) + ab*aa/255)
	return (oa << 24) | (or << 16) | (og << 8) | ob
}
