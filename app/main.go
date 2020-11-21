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

	if err := backup(conf); err != nil {
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

func backup(conf Config) error {
	var srcFile = fmt.Sprintf("%s%s", conf.SourcePath(), conf.FileName())
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := src.Close(); err != nil{
			log.Fatal(err)
		}
	}()

	if !dirExists(conf.BackupPath()) {
		if err := os.MkdirAll(conf.BackupPath(), 0755); err != nil {
			return err
		}
	}

	var dstFile = fmt.Sprintf("%s%s", conf.BackupPath(), conf.FileName())
	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func(){
		if err := dst.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}