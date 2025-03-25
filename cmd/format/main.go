package main

import (
	"database/sql"
	"fmt"
	"time"

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

	rows, err := db.Query("SELECT id, name, age, created_at FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	f := excelize.NewFile()
	sheet := "Users"
	f.SetSheetName(f.GetSheetName(0), sheet)

	f.SetColWidth(sheet, "A", "D", 20.0)

	dateStyle, _ := f.NewStyle(&excelize.Style{
		NumFmt: 22, // 2025-03-25 00:00:00 形式
	})

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F81BD"},
			Pattern: 1,
		},
	})
	cellStyle, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "all", Color: "#000000", Style: 1},
		},
	})

	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "Name")
	f.SetCellValue(sheet, "C1", "Age")
	f.SetCellValue(sheet, "D1", "Created At")
	f.SetCellStyle(sheet, "A1", "D1", headerStyle)

	rowNum := 2
	for rows.Next() {
		var id int
		var name string
		var age int
		var createdAtRaw []byte
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &age, &createdAtRaw); err != nil {
			fmt.Println("スキャンエラー:", err)
			continue
		}

		// 文字列に変換してパース
		createdAtStr := string(createdAtRaw)
		createdAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			fmt.Println("日付パースエラー:", err)
			continue
		}

		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNum), age)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNum), createdAt)
		f.SetCellStyle(sheet, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("D%d", rowNum), cellStyle)
		f.SetCellStyle(sheet, fmt.Sprintf("D%d", rowNum), fmt.Sprintf("D%d", rowNum), dateStyle)
		rowNum++
	}

	if err := f.SaveAs("users.xlsx"); err != nil {
		return err
	}

	fmt.Println("Excel ファイル作成完了")
	return nil
}

func main() {
	err := exportToExcel()
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
