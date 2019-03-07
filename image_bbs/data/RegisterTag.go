package data

import (
	"database/sql"

	"../typefile"

	_ "github.com/lib/pq"
)

//フォームから受け取ったタグデータを登録する
func RegisterTag(tag typefile.TagData) {
	//データベースと接続する
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	if err != nil {
		panic(err)
	}
	//タグデータを登録するSQL文をセットする
	stmt, err := db.Prepare("INSERT INTO tags(img_id,tag_name,create_date,update_date) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err)
	}
	//タグデータをセットして実行
	stmt.Exec(tag.ImgId, tag.TagName, tag.CreateDate, tag.UpdateDate)
	//DBとのコネクションを切断
	db.Close()
	return
}
