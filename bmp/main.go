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

func main() {
	dat, err := ioutil.ReadFile("file/480-360-sample.bmp")
	header := NewHeader(dat, 0)
	fmt.Println(header)
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
