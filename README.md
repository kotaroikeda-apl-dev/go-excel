## **概要**

## **実行方法**

```sh
docker compose up -d # データベース起動
go run cmd/basic/main.go # データベースからExcelを出力
go run cmd/format/main.go # Excel のフォーマット調整
go run cmd/insert-create/main.go # データベースにデータを挿入するSQLを作成
go run cmd/stream/main.go # 大容量のデータをストリームでExcelに出力
docker compose down # データベース停止
```

## **学習ポイント**

1. **`excelize`** を使って Excel ファイルを作成し、**`SetCellValue()`** でデータをセルに書き込む。
2. **`SetCellStyle()`** でセルのスタイルを設定する。
3. **`SetColWidth()`** で列の幅を設定する。
4. **`excelize.Font`** でフォントのスタイルを設定する。
5. **`excelize.Fill`** でセルの背景色を設定する。
6. **`excelize.Border`** でセルの枠線を設定する。
7. **`excelize.StreamWriter`** で大容量のデータをストリームで Excel に出力する。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
