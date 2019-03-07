# 画像掲示板

画像をアップロードでき、その画像を表示し削除や、タグが付けられるプログラム

使用言語：Golang

データベース：postgreSQL

画像ストレージ：Google Cloud Storage

実行仮想マシン：Google Compute Engine


## 必要ライブラリなど
* フレームワーク　echo

jwt-go,echoをターミナルから以下のコマンドでインストールする。
`$　go get -u github.com/labstack/echo`
jwt-goをターミナルから以下のコマンドでインストールする。
`$　go get github.com/dgrijalva/jwt-go`

* gcsクライアントライブラリ

`$go get -u cloud.google.com/go/storage`

* GCPのサービスアカウントキーを環境変数に設定する

GCP Consoleでサービアカウントキーを取得します。

ターミナルで~/.bash_profileを開き以下を追記して設定する
[PATH] は、サービス アカウント キーが含まれる JSON ファイルのファイルパスです。

`export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"`


GCP サービスアカウントキーに関して参考サイト
https://cloud.google.com/docs/authentication/getting-started?hl=ja

* データベースのセットアップについて
postgreSQLをセットアップし、ユーザー`trainer`を準備する。

psqlに接続し、テーブルを作成する。
```
$ psql -d postgres;

CREATE DATABASE imagebbs;

\connect imagebbs;

CREATE TABLE images
( img_id            serial,
    img_path         varchar(255),
    create_date         date,
    update_date            date);

GRANT ALL ON images TO trainer;

CREATE TABLE tags
(
  tag_id serial,
  img_id int,
  tag_name varchar(30),
  create_date date,
  update_date date
  );


GRANT ALL ON tags TO trainer;

GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO trainer;
```

#### ブラウザから画像を添付して送信し、それをGCSに登録して表示するプログラム

ファイルのあるフォルダへ移動し
 `$ go run ImgRegister.go`でサーバーを起動する
 ブラウザでlocalhost:3000にアクセスする。
 その後、画像を添付し登録ボタンを押下することでGCSに画像がアップロードされ、それが表示される。

# 画像掲示板

ファイルのあるフォルダへ移動し
 `$ go run ImgBBS.go`でサーバーを起動する

## トップ画面にアクセスした時に、画像データが一覧で表示される機能

 ブラウザでlocalhost:3000にアクセスする。

## トップ画面にアクセスした時に、タグデータが一覧で表示される機能

  ブラウザでlocalhost:3000にアクセスする。

## ブラウザから画像をGCSに登録し、登録した画像のパスと日時をデータベースに登録する機能

 ブラウザでlocalhost:3000にアクセスする。
 その後、画像を添付し登録ボタンを押下することでGCSに画像がアップロードされ、アップロードされた画像のパスと日時がデータベースに登録される。

## ブラウザから削除ボタンを押し個別の画像がGCSとデータベースから削除される機能

 ブラウザでlocalhost:3000にアクセスする。
 その後、削除したい画像上の削除ボタンを押下することでGCSとデータベースからその画像に関するデータが削除される

## ブラウザから個別の画像にタグ登録する機能

  ブラウザでlocalhost:3000にアクセスする。
  その後、タグを追加したい画像のテキストボックスにタグ名を入力しタグ登録ボタンを押下することでデータベースにタグデータが登録される。

## ブラウザから個別のタグを削除する機能

  ブラウザでlocalhost:3000にアクセスする。
  その後、削除したいタグのタグ削除ボタンを押下することで対象のタグがデータベースから削除される。
