package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

const dbURL = "user:password@tcp(localhost:3306)/dbname"

func exportToExcel() error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	f := excelize.NewFile()
	sheet := "Users"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// ヘッダー用のスタイル（フォントサイズ14、太字）
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 14,
			Bold: true,
		},
	})
	if err != nil {
		return err
	}

	// データ用のスタイル（フォントサイズ12）
	dataStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 12,
		},
	})
	if err != nil {
		return err
	}

	// ヘッダー
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "Name")
	f.SetCellValue(sheet, "C1", "Age")

	// ヘッダーにスタイル適用
	f.SetCellStyle(sheet, "A1", "C1", headerStyle)

	// 列幅を調整
	f.SetColWidth(sheet, "A", "A", 10) // ID列は狭く
	f.SetColWidth(sheet, "B", "B", 30) // 名前列は広く（日本語のため）
	f.SetColWidth(sheet, "C", "C", 10) // 年齢列も狭く

	rowNum := 2
	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)

		// データ行にスタイル適用
		f.SetCellStyle(sheet, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("C%d", rowNum), dataStyle)

		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNum), age)
		rowNum++
	}

	if err := f.SaveAs("users.xlsx"); err != nil {
		return err
	}

	fmt.Println("Excel ファイル作成完了")
	return nil
}

func main() {
	if err := exportToExcel(); err != nil {
		fmt.Printf("エラーが発生しました: %v\n", err)
	}
}
