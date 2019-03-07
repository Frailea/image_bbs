package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//個別のタグデータを削除する
func DeleteTagData(tagId int) {
	//データベースと接続
	db, err := sql.Open("postgres", "user=trainer password=1111 dbname=imagebbs sslmode=disable")
	if err != nil {
		panic(err)
	}
	//タグデータを削除するSQL文をセット
	stmt, err := db.Prepare("delete from tags where tag_id =$1")
	if err != nil {
		panic(err)
	}
	//対象のタグのIDを渡しsqlを実行
	stmt.Exec(tagId)
	return
}
