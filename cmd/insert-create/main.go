package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("insert_users.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 開始日付
	startDate, _ := time.Parse("2006-01-02", "2025-03-20")

	fmt.Fprintln(file, "INSERT INTO users (name, age, created_at) VALUES")

	for i := 1; i <= 100000; i++ {
		name := fmt.Sprintf("'User%d'", i)
		age := 20 + (i % 30)
		date := startDate.AddDate(0, 0, i%7).Format("2006-01-02 15:04:05")

		// 最後の行だけセミコロン
		if i == 100000 {
			fmt.Fprintf(file, "(%s, %d, '%s');\n", name, age, date)
		} else {
			fmt.Fprintf(file, "(%s, %d, '%s'),\n", name, age, date)
		}
	}

	fmt.Println("insert_users.sql を生成しました。")
}
