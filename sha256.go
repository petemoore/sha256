package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	K = [64]uint32{
		0x428a2f98,
		0x71374491,
		0xb5c0fbcf,
		0xe9b5dba5,
		0x3956c25b,
		0x59f111f1,
		0x923f82a4,
		0xab1c5ed5,
		0xd807aa98,
		0x12835b01,
		0x243185be,
		0x550c7dc3,
		0x72be5d74,
		0x80deb1fe,
		0x9bdc06a7,
		0xc19bf174,
		0xe49b69c1,
		0xefbe4786,
		0x0fc19dc6,
		0x240ca1cc,
		0x2de92c6f,
		0x4a7484aa,
		0x5cb0a9dc,
		0x76f988da,
		0x983e5152,
		0xa831c66d,
		0xb00327c8,
		0xbf597fc7,
		0xc6e00bf3,
		0xd5a79147,
		0x06ca6351,
		0x14292967,
		0x27b70a85,
		0x2e1b2138,
		0x4d2c6dfc,
		0x53380d13,
		0x650a7354,
		0x766a0abb,
		0x81c2c92e,
		0x92722c85,
		0xa2bfe8a1,
		0xa81a664b,
		0xc24b8b70,
		0xc76c51a3,
		0xd192e819,
		0xd6990624,
		0xf40e3585,
		0x106aa070,
		0x19a4c116,
		0x1e376c08,
		0x2748774c,
		0x34b0bcb5,
		0x391c0cb3,
		0x4ed8aa4a,
		0x5b9cca4f,
		0x682e6ff3,
		0x748f82ee,
		0x78a5636f,
		0x84c87814,
		0x8cc70208,
		0x90befffa,
		0xa4506ceb,
		0xbef9a3f7,
		0xc67178f2,
	}
)

type (
	Block  [64]byte
	Padded []Block
	SHA256 [8]uint32
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	if len(os.Args) == 1 {
		ShowCalculation([]byte("abc"))
		ShowCalculation([]byte("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"))
	} else {
		file := os.Args[1]
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		ShowCalculation(data)
	}
}

func ShowCalculation(bytes []byte) *Padded {
	zeroBytePaddingCount := 63 - ((len(bytes) + 8) % 64)
	paddedLengthBytes := len(bytes) + 9 + zeroBytePaddingCount
	blockCount := paddedLengthBytes / 64
	b := make([]Block, blockCount, blockCount)
	for i := 0; i < blockCount; i++ {
		for j := 0; j < min(64, len(bytes)-i*64); j++ {
			b[i][j] = bytes[i*64+j]
		}
	}
	b[len(bytes)/64][len(bytes)%64] = 128
	b[blockCount-1][56] = byte(uint64(len(bytes)) >> 0x35 & 0xff)
	b[blockCount-1][57] = byte(uint64(len(bytes)) >> 0x2d & 0xff)
	b[blockCount-1][58] = byte(uint64(len(bytes)) >> 0x25 & 0xff)
	b[blockCount-1][59] = byte(uint64(len(bytes)) >> 0x1d & 0xff)
	b[blockCount-1][60] = byte(uint64(len(bytes)) >> 0x15 & 0xff)
	b[blockCount-1][61] = byte(uint64(len(bytes)) >> 0x0d & 0xff)
	b[blockCount-1][62] = byte(uint64(len(bytes)) >> 0x05 & 0xff)
	b[blockCount-1][63] = byte(uint64(len(bytes)) << 0x03 & 0xff)
	p := Padded(b)
	fmt.Printf("    %v.\n\n", p.CalculateSHA256())
	return &p
}

func (padded *Padded) Original() []byte {
	l := (*padded)[len(*padded)-1][56:]
	origLength := (uint64(l[7]) | uint64(l[6])<<8 | uint64(l[5])<<16 | uint64(l[4])<<24 | uint64(l[3])<<32 | uint64(l[2])<<40 | uint64(l[1])<<48 | uint64(l[0])<<56) / 8
	orig := make([]byte, origLength, origLength)
	for i := range orig {
		orig[i] = (*padded)[i/64][i%64]
	}
	return orig
}

func (padded *Padded) CalculateSHA256() *SHA256 {

	Ha := uint32(0x6a09e667)
	Hb := uint32(0xbb67ae85)
	Hc := uint32(0x3c6ef372)
	Hd := uint32(0xa54ff53a)
	He := uint32(0x510e527f)
	Hf := uint32(0x9b05688c)
	Hg := uint32(0x1f83d9ab)
	Hh := uint32(0x5be0cd19)

	W := [64]uint32{}

	for i, block := range *padded {

		fmt.Printf("\n# Hash of %q (block %v)\n\n", string(padded.Original()), i+1)
		fmt.Println("            a         b         c         d         e         f         g         h\n")
		fmt.Printf("init:   %08x  %08x  %08x  %08x  %08x  %08x  %08x  %08x\n", Ha, Hb, Hc, Hd, He, Hf, Hg, Hh)

		a, b, c, d, e, f, g, h := Ha, Hb, Hc, Hd, He, Hf, Hg, Hh
		for j := 0; j < 64; j++ {
			if j < 16 {
				W[j] = uint32(block[j*4])<<24 | uint32(block[j*4+1])<<16 | uint32(block[j*4+2])<<8 | uint32(block[j*4+3])
			} else {
				W[j] = σ1(W[j-2]) + W[j-7] + σ0(W[j-15]) + W[j-16]
			}
			T1 := h + Σ1(e) + Ch(e, f, g) + K[j] + W[j]
			T2 := Σ0(a) + Maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
			fmt.Printf("t = %2v  %08x  %08x  %08x  %08x  %08x  %08x  %08x  %08x\n", j, a, b, c, d, e, f, g, h)
		}

		fmt.Printf("\nBlock %v has been processed. The values of {Hi} are\n\n", i+1)

		fmt.Printf("H1 = %08x + %08x = %08x\n", Ha, a, Ha+a)
		fmt.Printf("H2 = %08x + %08x = %08x\n", Hb, b, Hb+b)
		fmt.Printf("H3 = %08x + %08x = %08x\n", Hc, c, Hc+c)
		fmt.Printf("H4 = %08x + %08x = %08x\n", Hd, d, Hd+d)
		fmt.Printf("H5 = %08x + %08x = %08x\n", He, e, He+e)
		fmt.Printf("H6 = %08x + %08x = %08x\n", Hf, f, Hf+f)
		fmt.Printf("H7 = %08x + %08x = %08x\n", Hg, g, Hg+g)
		fmt.Printf("H8 = %08x + %08x = %08x\n", Hh, h, Hh+h)

		Ha += a
		Hb += b
		Hc += c
		Hd += d
		He += e
		Hf += f
		Hg += g
		Hh += h
	}

	fmt.Println("\nThe message digest is\n")

	sha256 := SHA256([8]uint32{Ha, Hb, Hc, Hd, He, Hf, Hg, Hh})
	return &sha256
}

func (sha256 *SHA256) String() string {
	return fmt.Sprintf("%08x %08x %08x %08x %08x %08x %08x %08x", sha256[0], sha256[1], sha256[2], sha256[3], sha256[4], sha256[5], sha256[6], sha256[7])
}

func (b Block) String() string {
	text := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			text += fmt.Sprintf("%02x", b[i*8+j])
		}
		if i < 7 {
			text += " "
		}
	}
	return text
}

func (p *Padded) String() string {
	text := ""
	for i, j := range *p {
		text = text + fmt.Sprintf("%v: %v\n", i, j)
	}
	return text
}

func Ch(x, y, z uint32) uint32 {
	return (x & y) ^ (z &^ x)
}

func Maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func Σ0(x uint32) uint32 {
	return (x<<30 | x>>2) ^ (x<<19 | x>>13) ^ (x<<10 | x>>22)
}

func Σ1(x uint32) uint32 {
	return (x<<26 | x>>6) ^ (x<<21 | x>>11) ^ (x<<7 | x>>25)
}

func σ0(x uint32) uint32 {
	return (x<<25 | x>>7) ^ (x<<14 | x>>18) ^ x>>3
}

func σ1(x uint32) uint32 {
	return (x<<15 | x>>17) ^ (x<<13 | x>>19) ^ x>>10
}
