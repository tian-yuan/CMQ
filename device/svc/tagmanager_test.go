package svc

import (
	"testing"
	"math"
	"github.com/golang/go/src/fmt"
)

func TestTagManager_AddGuid(t *testing.T) {
	tagManager := NewTagManager("android", 16)
	guid := uint32(31)
	fmt.Printf("add guid : %d\n", guid)
	tagManager.AddGuid(guid)
	guid = uint32(101)
	fmt.Printf("add guid : %d\n", guid)
	tagManager.AddGuid(guid)
	guid = uint32(math.Exp2(17))
	fmt.Printf("add guid : %d\n", guid)
	tagManager.AddGuid(guid)
	sectionArr := tagManager.GetSectionArray()
	for _, section := range sectionArr {
		guidArr := tagManager.GetGuidArray(section)
		fmt.Println(guidArr)
	}
}

func TestTagManager_Broadcast(t *testing.T) {
	tagAndroid := NewTagManager("android", 16)
	tagIos := NewTagManager("ios", 16)
	guid := uint32(31)
	tagAndroid.AddGuid(guid)
	tagIos.AddGuid(guid)
	tagAndroid.AddGuid(uint32(32))
	tagAndroid.AddGuid(uint32(math.Exp2(17)))
	tagIos.AddGuid(uint32(math.Exp2(18) + 1))
	androidSectionArr := tagAndroid.GetSectionArray()
	fmt.Printf("android section arr : ")
	fmt.Print(androidSectionArr)
	fmt.Println()
	iosSectionArr := tagIos.GetSectionArray()
	fmt.Printf("ios section arr : ")
	fmt.Print(iosSectionArr)
	fmt.Println()
	androidBitmap := tagAndroid.GetBitmap(0)
	iosBitmap := tagIos.GetBitmap(0)
	androidBitmap.And(iosBitmap)
	fmt.Println(androidBitmap)
	fmt.Println(math.Exp2(16))
}
