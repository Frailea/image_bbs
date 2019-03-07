package data

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

func ImgRegister(f multipart.File, t time.Time) string {

	//コンテキスト(リクエストを渡す)生成
	ctx := context.Background()
	//ストレージを扱うクライアント取得
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//日時を取得
	layout := "2006-01-02 15:04:05"
	str := t.Format(layout)
	//バケット名、オブジェクト名（登録日時）
	bucket := client.Bucket("ishida-kadai6")
	object := bucket.Object("images/" + str + ".jpg")
	//コンテキストを書き込みモードでオープン
	writer := object.NewWriter(ctx)
	//タイプjpg
	writer.ObjectAttrs.ContentType = "image/jpg"
	//アクセス制御リストを公開設定
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		storage.ACLRule{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}
	//受け取って開いた画像をwriterにコピー
	if _, err := io.Copy(writer, f); err != nil {
		log.Fatal(err)
	}
	//writerをクローズする。
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}
	//バケットの属性取得
	attr := writer.Attrs()
	// アップロードした画像パスをハンドルに返す
	return attr.Name
}
