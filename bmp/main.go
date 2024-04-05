package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
)

type Header struct {
	signature  uint16
	fileSize   uint32
	reserved   uint32
	dataOffset uint32
}

func NewHeader(data []byte, offset int) Header {
	return Header{
		signature:  binary.LittleEndian.Uint16(data[offset:]),
		fileSize:   binary.LittleEndian.Uint32(data[offset+2:]),
		reserved:   binary.LittleEndian.Uint32(data[offset+6:]),
		dataOffset: binary.LittleEndian.Uint32(data[offset+10:]),
	}
}

type InfoHeader struct {
	size            uint32
	width           uint32
	height          uint32
	planes          uint16
	bpp             uint16
	compression     uint32
	imageSize       uint32
	xpixels         uint32
	ypixels         uint32
	colors          uint32
	importantColors uint32
}

func NewInfoHeader(data []byte, offset int) InfoHeader {
	return InfoHeader{
		size:            binary.LittleEndian.Uint32(data[offset+0:]),
		width:           binary.LittleEndian.Uint32(data[offset+4:]),
		height:          binary.LittleEndian.Uint32(data[offset+8:]),
		planes:          binary.LittleEndian.Uint16(data[offset+12:]),
		bpp:             binary.LittleEndian.Uint16(data[offset+14:]),
		compression:     binary.LittleEndian.Uint32(data[offset+16:]),
		imageSize:       binary.LittleEndian.Uint32(data[offset+20:]),
		xpixels:         binary.LittleEndian.Uint32(data[offset+24:]),
		ypixels:         binary.LittleEndian.Uint32(data[offset+28:]),
		colors:          binary.LittleEndian.Uint32(data[offset+32:]),
		importantColors: binary.LittleEndian.Uint32(data[offset+36:]),
	}
}

type Pixel struct {
	r uint8
	g uint8
	b uint8
}

func NewPixel(data []byte, offset int) Pixel {
	return Pixel{
		b: data[offset+0],
		g: data[offset+1],
		r: data[offset+2],
	}
}

type BMP24 struct {
	header Header
	info   InfoHeader
	data   []Pixel
}

func (b BMP24) DataSize() uint32 {
	return b.info.width * b.info.height * 3
}

func main() {
	dat, err := ioutil.ReadFile("file/64-64.bmp")
	img := BMP24{}
	img.header = NewHeader(dat, 0)
	img.info = NewInfoHeader(dat, 14)

	fmt.Println(img)
	if err != nil {
		log.Fatal("nope")
	}
	for i := 0; i < len(dat); i++ {
		fmt.Printf("%03d ", dat[i])
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}
