package svc

import "github.com/RoaringBitmap/roaring"

type BitmapSystem struct {
	Conf *BitmapConfig
	BitmapSectionMap map[uint32]*roaring.Bitmap
}

type BitmapConfig struct {
	HighbitsCount uint32
}

func NewBitmapConf() *BitmapConfig {
	return &BitmapConfig{
		HighbitsCount: 16,
	}
}

func NewBitmapSystem(conf *BitmapConfig) *BitmapSystem {
	return &BitmapSystem{
		Conf: conf,
		BitmapSectionMap: make(map[uint32]*roaring.Bitmap),
	}
}

const MAX_UINT32 = ^uint32(0)

func (bs *BitmapSystem) GetHighbit(guid uint32) uint32 {
	return guid >> (32 - bs.Conf.HighbitsCount)
}

func (bs *BitmapSystem) GetLowbit(guid uint32) uint32 {
	return (guid << bs.Conf.HighbitsCount) >> bs.Conf.HighbitsCount
}

func (bs *BitmapSystem) GetGuid(highbit uint32, lowbit uint32) uint32 {
	return (highbit << bs.Conf.HighbitsCount) + lowbit
}

func (bs *BitmapSystem) AddGuid(guid uint32) {
	highbit := bs.GetHighbit(guid)
	lowbit := bs.GetLowbit(guid)
	if bs.BitmapSectionMap[highbit] == nil {
		bs.BitmapSectionMap[highbit] = new(roaring.Bitmap)
	}
	bs.BitmapSectionMap[highbit].Add(lowbit)
}

func (bs *BitmapSystem) GetBitmap(sectionIndex uint32) *roaring.Bitmap {
	return bs.BitmapSectionMap[sectionIndex]
}

func (bs *BitmapSystem) SetBitmap(sectionIndex uint32, bitmap *roaring.Bitmap) {
	bs.BitmapSectionMap[sectionIndex] = bitmap
}

func (bs *BitmapSystem) GetSectionArray() []uint32 {
	keys := make([]uint32, 0, len(bs.BitmapSectionMap))
	for k := range bs.BitmapSectionMap{
		keys = append(keys, k)
	}
	return keys
}
