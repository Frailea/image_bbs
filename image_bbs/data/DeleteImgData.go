package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//データベースのデータを一括削除する
func DeleteImgData(imgPath string) {
	//データベースと接続
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	if err != nil {
		panic(err)
	}
	//削除するsql文をセット
	stmt, err := db.Prepare("delete from images where img_path =$1")
	if err != nil {
		panic(err)
	}
	//対象の画像パスを渡しsqlを実行
	stmt.Exec(imgPath)

	return
}
