package data

import (
	"database/sql"
	"fmt"
	"time"

	"../typefile"

	_ "github.com/lib/pq"
)

//データベースに受け取った文字列を登録する
func ImgDataRegister(imagepath string, t time.Time) {
	//登録日時のフォーマットを変更
	layout := "2006-01-02"
	str := t.Format(layout)
	//画像パス、登録日時、更新日時を構造体に格納
	var image typefile.ImageData
	image.ImgPath = imagepath
	image.CreateDate = str
	image.UpdateDate = str
	fmt.Println(image)
	//データベースと接続
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	//エラーチェック
	//sql文をセットする
	stmt, err := db.Prepare("INSERT INTO images(img_path,create_date,update_date) VALUES($1,$2,$3)")
	if err != nil {
		panic(err)
	}
	//構造体に格納してある画像データをデータベースに登録
	stmt.Exec(image.ImgPath, image.CreateDate, image.UpdateDate)
	//DBとのコネクションを切断
	db.Close()
	return
}
