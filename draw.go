package main // import "github.com/serverwentdown/leet"

import (
	"log"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// TODO: full-sized grid with multiple channels mapped onto the grid

type Drawer struct {
	layers [][]uint32
	length int
	device *ws2811.WS2811
}

func NewDrawer(length int) (*Drawer, error) {
	opt := ws2811.DefaultOptions
	opt.Channels[0].LedCount = length

	d := &Drawer{}
	d.length = length
	d.layers = make([][]uint32, 100)
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
	d.layers[layer] = dots
}

func (d *Drawer) SetLayerOrFill(layer int32, dots []uint32, fill uint32) {
	if dots == nil {
		fillDots := make([]uint32, d.length)
		for i := 0; i < d.length; i++ {
			fillDots[i] = fill
		}
		d.layers[layer] = fillDots
	} else {
		d.SetLayer(layer, dots)
	}
}

func (d *Drawer) Draw() error {
	for i := 0; i < len(d.device.Leds(0)); i++ {
		dot := mix(d.layers, i)
		d.device.Leds(0)[i] = dot
	}

	if err := d.device.Render(); err != nil {
		return err
	}
	return nil
}

func mix(layers [][]uint32, j int) uint32 {
	var base uint32
	for i := 0; i < len(layers); i++ {
		base = mixColors(base, layers[i][j])
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
	log.Println(aa, ar, ag, ab, ba, br, bg, bb)
	oa := uint32(ba*(255-aa)/255 + aa)
	or := uint32(br*ba*(255-aa)/(255*255) + ar*aa/255)
	og := uint32(bg*ba*(255-aa)/(255*255) + ag*aa/255)
	ob := uint32(bb*ba*(255-aa)/(255*255) + ab*aa/255)
	return (oa << 24) | (or << 16) | (og << 8) | ob
}
