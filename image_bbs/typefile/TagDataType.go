package typefile

import "time"

// タグのデータを格納する構造体
type TagData struct {
	TagId      int
	ImgId      int
	TagName    string
	CreateDate string
	UpdateDate string
}

func InitTag(imgId int, tagName string) TagData {
	tag := new(TagData)
	//現在の時間を取得し日時に変更する
	t := time.Now()
	layout := "2006-01-02"
	str := t.Format(layout)
	tag.ImgId = imgId
	tag.TagName = tagName
	tag.CreateDate = str
	tag.UpdateDate = str
	return *tag
}
