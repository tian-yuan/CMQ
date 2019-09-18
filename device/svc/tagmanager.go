package svc

import "github.com/RoaringBitmap/roaring"

type TagConfg struct {
	TagName string
}

type TagManager struct {
	Conf *TagConfg
	Bitmap *BitmapSystem
}

func NewTagManager(tagName string, highbitCount uint32) *TagManager {
	conf := new(TagConfg)
	conf.TagName = tagName
	bitmapConf := NewBitmapConf()
	bitmapConf.HighbitsCount = highbitCount
	bitmap := NewBitmapSystem(bitmapConf)
	return &TagManager{
		Conf: conf,
		Bitmap: bitmap,
	}
}

func (tm *TagManager) AddGuid(guid uint32) {
	tm.Bitmap.AddGuid(guid)
}

func (tm *TagManager) GetBitmap(sectionIndex uint32) *roaring.Bitmap  {
	return tm.Bitmap.GetBitmap(sectionIndex)
}

func (tm *TagManager) GetSectionArray() []uint32 {
	return tm.Bitmap.GetSectionArray()
}

func (tm *TagManager) GetGuidArray(sectionIndex uint32) []uint32 {
	bitmap := tm.Bitmap.GetBitmap(sectionIndex)
	if bitmap == nil {
		return nil
	}
	guidArr := bitmap.ToArray()
	for i, guid := range guidArr {
		guidArr[i] = tm.Bitmap.GetGuid(sectionIndex, guid)
	}
	return guidArr
}
