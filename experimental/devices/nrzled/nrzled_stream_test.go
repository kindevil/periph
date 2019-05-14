// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package nrzled

import (
	"bytes"
	"image"
	"image/color"
	"testing"

	"github.com/kindevil/periph/conn/gpio/gpiostream"
	"github.com/kindevil/periph/conn/gpio/gpiostream/gpiostreamtest"
	"github.com/kindevil/periph/conn/physic"
)

func TestStream_NewBits_3(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		N: "Yo",
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92,
					0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49,
					0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92,
					0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49,
					0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	if s := d.String(); s != "nrzled{Yo}" {
		t.Fatal(s)
	}
	if c := d.ColorModel(); c != color.NRGBAModel {
		t.Fatal(c)
	}
	if r := d.Bounds(); r != image.Rect(0, 0, 10, 1) {
		t.Fatal(r)
	}
	if err = d.Halt(); err != nil {
		t.Fatal(err)
	}
	if err = g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_New_fail(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{}
	opts := DefaultOpts
	opts.Freq = 0
	if _, err := NewStream(&g, &opts); err == nil {
		t.Fatal("hz == 0")
	}
	opts = DefaultOpts
	opts.Channels = 2
	if _, err := NewStream(&g, &opts); err == nil {
		t.Fatal("channels == 2")
	}
}

func TestStream_Draw_NRGBA_3(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d,
					0xb6, 0x92, 0x49, 0xb6, 0x92, 0x49, 0xb4, 0x92, 0x4d, 0x24,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	img := image.NewNRGBA(d.Bounds())
	copy(img.Pix, getRGBW())
	if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
		t.Fatal(err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Draw_RGBA_3(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d,
					0xa6, 0xda, 0x49, 0xb6, 0xd3, 0x4d, 0x34, 0xdb, 0x49, 0x36,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(d.Bounds())
	copy(img.Pix, getRGBW())
	if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
		t.Fatal(err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Draw_RGBA_4(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92,
					0x49, 0x24, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6,
					0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6,
					0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xa6, 0xd2, 0x49, 0x24, 0xda, 0x49, 0xb6, 0xd3,
					0x4d, 0x34, 0xdb, 0x49, 0x36, 0x92, 0x4d, 0x26,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	opts.Channels = 4
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(d.Bounds())
	copy(img.Pix, getRGBW())
	if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
		t.Fatal(err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Draw_Limits(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d,
					0xa6, 0xda, 0x49, 0xb6, 0xd3, 0x4d, 0x34, 0xdb, 0x49, 0x36,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(image.Rect(-1, -1, 20, 20))
	copy(img.Pix, getRGBW())
	if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
		t.Fatal(err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Write_3(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{
		Ops: []gpiostream.Stream{
			&gpiostream.BitStream{
				Bits: []byte{
					0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49,
					0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0x92, 0x49, 0x24,
					0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0x24, 0xdb,
					0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb6, 0x92, 0x49, 0xa4, 0x92, 0x49, 0x36, 0x92, 0x49,
					0xa6, 0x92, 0x49, 0xb6, 0x92, 0x49, 0xb4, 0x92, 0x4d, 0x24,
					0x00, 0x00, 0x00,
				},
				Freq: 800 * physic.KiloHertz,
				LSBF: false,
			},
		},
	}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	if n, err := d.Write(getRGB()); n != 30 || err != nil {
		t.Fatal(n, err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Write_fail(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{DontPanic: true}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	if n, err := d.Write([]byte{1}); n != 0 || err == nil {
		t.Fatal(n, err)
	}
	if n, err := d.Write([]byte{1, 2, 3}); n != 0 || err == nil {
		t.Fatal(n, err)
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Halt_fail(t *testing.T) {
	g := gpiostreamtest.PinOutPlayback{DontPanic: true}
	opts := DefaultOpts
	opts.NumPixels = 10
	d, err := NewStream(&g, &opts)
	if err != nil {
		t.Fatal(err)
	}
	if d.Halt() == nil {
		t.Fatal("expected failure")
	}
	if err := g.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestStream_Raster_3_3(t *testing.T) {
	data := []byte{
		// 24 bits per pixel in RGB
		0, 1, 2,
		0xFD, 0xFE, 0xFF,
	}
	expected := []byte{
		// 72 bits per pixel in GRB
		0x92, 0x49, 0x26, 0x92, 0x49, 0x24, 0x92, 0x49, 0x34,
		0xdb, 0x6d, 0xb4, 0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xb6,
	}
	actual := make([]byte, len(expected))
	rasterBits(actual, data, 3, 3)
	if !bytes.Equal(expected, actual) {
		t.Fatalf("\nexpected %#v\n  actual %#v", expected, actual)
	}
}

func TestStream_Raster_4_4(t *testing.T) {
	data := []byte{
		// 32 bits per pixel in RGBW
		0, 1, 2, 3,
		0xFC, 0xFD, 0xFE, 0xFF,
	}
	expected := []byte{
		// 96 bits per pixel in GRBW
		0x92, 0x49, 0x26, 0x92, 0x49, 0x24, 0x92, 0x49, 0x34, 0x92, 0x49, 0x36,
		0xdb, 0x6d, 0xa6, 0xdb, 0x6d, 0xa4, 0xdb, 0x6d, 0xb4, 0xdb, 0x6d, 0xb6,
	}
	actual := make([]byte, len(expected))
	rasterBits(actual, data, 4, 4)
	if !bytes.Equal(expected, actual) {
		t.Fatalf("\nexpected %#v\n  actual %#v", expected, actual)
	}
}

//

// getRGB returns a buffer of 10 RGB pixels.
func getRGB() []byte {
	return []byte{
		0x00, 0x00, 0x00,
		0x00, 0x00, 0xFF,
		0x00, 0xFF, 0x00,
		0x00, 0xFF, 0xFF,
		0xFF, 0x00, 0x00,
		0xFF, 0x00, 0xFF,
		0xFF, 0xFF, 0x00,
		0xFF, 0xFF, 0xFF,
		3, 4, 5,
		6, 7, 8,
	}
}

// getRGBW returns a buffer of 10 RGB pixels.
func getRGBW() []byte {
	return []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xFF, 0xFF,
		0x00, 0xFF, 0x00, 0xFF,
		0x00, 0xFF, 0xFF, 0xFF,
		0xFF, 0x00, 0x00, 0xFF,
		0xFF, 0x00, 0xFF, 0xFF,
		0xFF, 0xFF, 0x00, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0x80,
		6, 7, 8, 9,
	}
}
