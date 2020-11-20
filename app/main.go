package main

import (
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/BurntSushi/toml"
)

func main() {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	var srcPath = conf.SourcePath()
	var dstPath = conf.BackupPath()

	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	fmt.Println(conf.FileName())

	if _, err = io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

	// f := excelize.NewFile()
	// // ワークシートを作成する
	// index := f.NewSheet("Sheet2")
	// // セルの値を設定
	// f.SetCellValue("Sheet2", "A1", "こんにちはせかい")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // ワークブックのデフォルトワークシートを設定します
	// f.SetActiveSheet(index)
	// // 指定されたパスに従ってファイルを保存します
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}
