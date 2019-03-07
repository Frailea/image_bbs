package data

import (
	"database/sql"

	"../typefile"
	_ "github.com/lib/pq"
)

//画像データを全て取得し構造体に入れそれを配列にする
func GetAllImageData() []typefile.ImageData {
	//データベースと接続
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	//データベースからデータを全件取得する
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select img_id,img_path,to_char(create_date,'YYYY年MM月DD日'),to_char(update_date,'YYYY年MM月DD日') from images;")
	if err != nil {
		panic(err)
	}
	//データベースとの接続を閉じる
	db.Close()
	//取得したデータを入れるための配列を用意
	var images []typefile.ImageData
	//構造体に取得したデータを入れる
	for rows.Next() {
		var image typefile.ImageData
		err := rows.Scan(
			&image.ImgId,
			&image.ImgPath,
			&image.CreateDate,
			&image.UpdateDate)
		if err != nil {
			panic(err)
		}
		//構造体に入れたデータを配列に格納する
		images = append(images, image)
	}
	//データを取り込んだrowsを閉じる
	rows.Close()
	//画像データを格納した配列を呼び出し元に渡す
	return images
}
