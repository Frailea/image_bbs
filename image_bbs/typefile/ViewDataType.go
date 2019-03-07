package typefile

import "fmt"

//画面に表示するデータを格納する構造体
type ViewData struct {
	Image ImageData
	Tags  []TagData
}

//画面に表示するデータ格納する構造体をスライスに追加しハンドルに返す
func GenerateViewList(images []ImageData, tags []TagData) []ViewData {
	var views []ViewData
	for _, image := range images {
		////画面に表示するデータを格納する構造体を作成する
		view := initView(image, tags)
		views = append(views, view)
	}
	return views
}

//画面に表示するデータを格納する構造体を作成する
func initView(image ImageData, tags []TagData) ViewData {

	//画像データをスライスから取り出し、viewに追加する
	var view ViewData
	var imgtag []TagData
	view.Image = image
	//タグデータをスライスから取り出し、画像IDで画像に紐づける
	for _, tag := range tags {
		if view.Image.ImgId == tag.ImgId {
			imgtag = append(imgtag, tag)
			view.Tags = imgtag
		}
	}
	fmt.Println(view)
	return view
}
