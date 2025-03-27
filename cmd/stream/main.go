package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

const dbURL = "user:password@tcp(localhost:3306)/dbname"

func exportLargeExcel() error {
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

	sw, _ := f.NewStreamWriter(sheet)

	// ヘッダー
	headers := []interface{}{"ID", "Name", "Age"}
	sw.SetRow("A1", headers)

	rowNum := 2
	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		row := []interface{}{id, name, age}
		cell := fmt.Sprintf("A%d", rowNum)
		sw.SetRow(cell, row)
		rowNum++
	}

	sw.Flush()

	if err := f.SaveAs("large_users.xlsx"); err != nil {
		return err
	}

	fmt.Printf("Excel ファイル作成完了（%d件）\n", rowNum-2)
	return nil
}

func main() {
	err := exportLargeExcel()
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
