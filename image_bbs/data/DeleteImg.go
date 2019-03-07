package data

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	_ "github.com/lib/pq"
)

//画像をGCS上から削除する
func DeleteImg(imgPath string) {
	//コンテキスト生成
	ctx := context.Background()
	//ストレージを扱うクライアント取得
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//バケット名、オブジェクト名を取得
	bucket := client.Bucket("ishida-kadai6")
	object := bucket.Object(imgPath)
	//対象のオブジェクトを削除
	if err := object.Delete(ctx); err != nil {
		log.Fatal(err)
	}
}
