package svc

import (
	"testing"
	"math"
	"github.com/golang/go/src/fmt"
)

func TestBitmapSystem_GetHighbit(t *testing.T) {
	conf := NewBitmapConf()
	conf.HighbitsCount = 2
	bitmapSystem := NewBitmapSystem(conf)
	highbit := bitmapSystem.GetHighbit(uint32(math.Exp2(31)))
	if highbit != 2 {
		t.Errorf("bitmap system get high bit error, expect : %d, get : %d", 2, highbit)
	}
	fmt.Printf("2 exp 31 highbit is : %d\n", highbit)
	highbit = bitmapSystem.GetHighbit(uint32(math.Exp2(32) - 1))
	if highbit != 3 {
		t.Errorf("bitmap system get high bit error, expect : %d, get : %d", 3, highbit)
	}
	fmt.Printf("2 exp 32 - 1 highbit is : %d\n", highbit)
	highbit = bitmapSystem.GetHighbit(uint32(math.Exp2(30) - 1))
	if highbit != 0 {
		t.Errorf("bitmap system get high bit error, expect : %d, get : %d", 2, highbit)
	}
	fmt.Printf("2 exp 30 - 1 highbit is : %d\n", highbit)
}

func TestBitmapSystem_GetLowbit(t *testing.T) {
	conf := NewBitmapConf()
	conf.HighbitsCount = 2
	bitmapSystem := NewBitmapSystem(conf)
	guid := uint32(31)
	lowbit := bitmapSystem.GetLowbit(guid)
	if lowbit != guid {
		t.Errorf("bitmap system get lowbit error, expect : %d, get : %d", guid, lowbit)
	}
	fmt.Printf("bitmap system get lowbit : %d\n", lowbit)

	guid = uint32(math.Exp2(31) + 1)
	lowbit = bitmapSystem.GetLowbit(guid)
	if lowbit != 1 {
		t.Errorf("bitmap system get lowbit error, expect : %d, get : %d", 1, lowbit)
	}
	fmt.Printf("bitmap system get lowbit : %d\n", lowbit)
}

func TestBitmapSystem_AddGuid(t *testing.T) {
	conf := NewBitmapConf()
	conf.HighbitsCount = 16
	bitmapSystem := NewBitmapSystem(conf)
	guid := uint32(31)
	bitmapSystem.AddGuid(guid)
	keys := bitmapSystem.GetSectionArray()
	if len(keys) != 1 {
		t.Error("section array size : %d", len(keys))
	}
	bitmap := bitmapSystem.GetBitmap(keys[0])
	fmt.Printf("key[0] : %d, bitmap : %s\n", keys[0], bitmap.String())

	guid = uint32(math.Exp2(17))
	bitmapSystem.AddGuid(guid)
	keys = bitmapSystem.GetSectionArray()
	if len(keys) != 2 {
		t.Error("section array size : %d", len(keys))
	}
	bitmap = bitmapSystem.GetBitmap(keys[0])
	fmt.Printf("key[0] : %d, bitmap : %s\n", keys[0], bitmap.String())
	bitmap = bitmapSystem.GetBitmap(keys[1])
	fmt.Printf("key[1] : %d, bitmap : %s\n", keys[1], bitmap.String())
}
