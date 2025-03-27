## **概要**

このプロジェクトは、Go 言語と **`excelize`** ライブラリを使って、**MySQL のデータを Excel ファイルとして出力する一連の流れ**を学べる実践的なサンプル集です。  
基本的な Excel 出力から、スタイルのカスタマイズ、大量データを高速に出力するためのストリーム処理まで、**業務レベルの Excel 出力に必要なテクニックを体系的に習得**できます。

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
