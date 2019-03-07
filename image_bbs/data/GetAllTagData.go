package data

import (
	"database/sql"

	"../typefile"
	_ "github.com/lib/pq"
)

//画像データを全て取得し構造体に入れそれを配列にする
func GetAllTagData() []typefile.TagData {
	//データベースと接続
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	//データベースからデータを全件取得する
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select img_id,tag_id,tag_name from tags;")
	if err != nil {
		panic(err)
	}
	//データベースとの接続を閉じる
	db.Close()
	//取得したデータを入れるための配列を用意
	var tags []typefile.TagData
	//構造体に取得したデータを入れる
	for rows.Next() {
		var tag typefile.TagData
		err := rows.Scan(
			&tag.ImgId,
			&tag.TagId,
			&tag.TagName,
		)
		if err != nil {
			panic(err)
		}
		//構造体に入れたデータを配列に格納する
		tags = append(tags, tag)
	}
	//データを取り込んだrowsを閉じる
	rows.Close()
	//画像データを格納した配列を呼び出し元に渡す
	return tags
}
