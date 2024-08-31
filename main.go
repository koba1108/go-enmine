package main

import (
	"fmt"
	"github.com/jhillyerd/enmime"
	"os"
)

func main() {
	// メール電文を読み込む
	file, err := os.Open("statics/emails/attachment.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	// メール電文を解析する
	envelope, err := enmime.ReadEnvelope(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	// 送信元の取得
	fmt.Printf("From: %v\n", envelope.GetHeader("From"))

	for _, at := range envelope.Attachments {
		// ファイル名
		fmt.Printf("FileName: %s\n", at.FileName)
		// ファイルの種類
		fmt.Printf("ContentType: %s\n", at.ContentType)
		// ファイルの保存
		f, err := os.Create("statics/attachments/" + at.FileName)
		if err != nil {
			fmt.Print(err)
			return
		}
		if _, err = f.Write(at.Content); err != nil {
			fmt.Print(err)
			return
		}
		if err = f.Close(); err != nil {
			fmt.Print(err)
			return
		}
	}
}
