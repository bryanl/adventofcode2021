package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	// input is one row

	input := rows[0]
	fmt.Println(input)

	return nil
}

func HexToBinary(hex string) string {
	var sb strings.Builder

	for _, h := range strings.Split(hex, "") {
		ui, err := strconv.ParseUint(h, 16, 64)
		if err != nil {
			panic(err)
		}

		sb.WriteString(fmt.Sprintf("%04b", ui))

	}

	return sb.String()
}

func BinaryToInt(bits string) int {
	i, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

type Packet interface {
	Header() *Header
}

type ValuePacket struct {
	header *Header
	Value  int
}

func NewValuePacket(header *Header, payload string) *ValuePacket {
	packet := &ValuePacket{
		header: header,
		Value:  DecodeValue(payload),
	}

	return packet
}

var _ Packet = &ValuePacket{}

func (v *ValuePacket) Header() *Header {
	return v.header
}

type OperatorPacket struct {
	header *Header
}

type Header struct {
	Version int
	Type    int
}

func DecodePacket(bits string) Packet {
	header := &Header{
		// 0-2 are version
		Version: BinaryToInt(bits[:3]),
		// 3-5
		Type: BinaryToInt(bits[3:6]),
	}

	switch header.Type {
	case 4:
		return NewValuePacket(header, bits[6:])
	}

	return nil
}

func DecodeValue(bits string) int {
	var sb strings.Builder

	parts := ChunkString(bits, 5)
	for _, part := range parts {
		if len(part) == 5 {
			sb.WriteString(part[1:])
		}
	}

	return BinaryToInt(sb.String())
}

func ChunkString(s string, size int) []string {
	var out []string

	for i := 0; i < len(s); i += size {
		end := i + size

		if end > len(s) {
			end = len(s)
		}

		out = append(out, s[i:end])
	}

	return out
}
