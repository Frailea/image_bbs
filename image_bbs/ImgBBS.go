package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"./data"
	"./typefile"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// レイアウト適用済のテンプレートを保存するmap
var templates map[string]*template.Template

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "layout.html", data)
}

//メイン　echoのインスタンスを作成
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// テンプレートを利用するためのRendererの設定
	t := &Template{}
	e.Renderer = t
	// 各ルーティングに対するハンドラを設定
	e.GET("/", RootHandler)
	e.POST("/form", ImgHandler)
	e.POST("/imgDelete", DeleteImageHandler)
	e.POST("/registerTag", RegisterTagDataHandler)
	e.POST("/deleteTag", DeleteTagDataHandler)
	// サーバーを開始
	e.Logger.Fatal(e.Start(":3000"))
}

// 初期化を行います。
func init() {
	loadTemplates()
}

// 各HTMLテンプレートに共通レイアウトを適用した結果を保存します
func loadTemplates() {
	var baseTemplate = "templates/layout.html"
	templates = make(map[string]*template.Template)
	templates["index"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/imgbase.html"))
}

//初期画面（トップ画面）を表示
func RootHandler(c echo.Context) error {
	//トップ画面を表示する
	error := RenderTopScreen(c)
	return error
}

//画像を登録して表示する。
func ImgHandler(c echo.Context) error {
	//現在の日時を取得
	t := time.Now()
	////フォームから画像を受け取る
	imgfile, err := c.FormFile("image")
	if err != nil {
		log.Fatal(err)
	}
	//入出力用にファイルを開く
	f, err := imgfile.Open()
	if err != nil {
		log.Fatal(err)
	}
	//登録日時と開いたマルチパートファイルを渡しGCSに登録する
	imagepath := data.ImgRegister(f, t)
	//遅延実行　fを閉じる
	defer f.Close()
	//登録日時と画像のパスを渡しデータベースに登録する
	data.ImgDataRegister(imagepath, t)
	//トップ画面を表示する
	error := RenderTopScreen(c)
	return error
}

//画像削除機能
func DeleteImageHandler(c echo.Context) error {
	//画像を識別する値を取得
	imgPath := c.FormValue("imgPath")
	//画像をGCSから削除する
	data.DeleteImg(imgPath)
	//画像をデータベースから削除する
	data.DeleteImgData(imgPath)
	//トップ画面を表示する
	error := RenderTopScreen(c)
	return error
}

//タグ登録機能
func RegisterTagDataHandler(c echo.Context) error {
	//フォームから受け取った画像ID、タグ名を変数に入れる
	imgId, _ := strconv.Atoi(c.FormValue("imgId"))
	tagname := c.FormValue("tagname")
	//画像ID、タグ名を渡しタグ構造体を作成する。
	tag := typefile.InitTag(imgId, tagname)
	//登録日時と画像のパスを渡しデータベースに登録する
	data.RegisterTag(tag)
	//トップ画面を表示する
	error := RenderTopScreen(c)
	return error
}

//対象のタグデータを削除する機能
func DeleteTagDataHandler(c echo.Context) error {
	//削除するタグのIDを取得
	tagId, _ := strconv.Atoi(c.FormValue("tagId"))
	//削除するタグのIDを渡しデータベースから削除する
	data.DeleteTagData(tagId)
	//トップ画面を表示する
	error := RenderTopScreen(c)
	return error
}

//トップ画面を表示する関数
func RenderTopScreen(c echo.Context) error {
	//タグデータ、画像データを取得し配列に格納
	images := data.GetAllImageData()
	tags := data.GetAllTagData()
	//画面データを格納する構造体を作成
	views := typefile.GenerateViewList(images, tags)
	//格納した画面データを渡し表示する
	return c.Render(http.StatusOK, "index", views)
}
