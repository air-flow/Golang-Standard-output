# Golang-Standard-output

汎用的に使用できる golang の標準出力作成
Python logging を再現したい。
ロギングレベル設定
format 設定

## DEMO

output : [3]float64{10, 6, 3}

## Features

- formatter 整形
- モード機能
- 履歴機能

## Requirement

### application

- Go 1.15.2

## Usage

```bash
git clone https://github.com/air-flow/Golang-Standard-output.git
cd ./logging
go logging.go

[Running]
Run Time        => 2020-12-24 23:07:43.5887739 +0900 JST m=+0.001996401
Function Names  => main.main
Run Line number => 167
"debug test"
```

## Note

2020 年 12 月 23 日

- logging 基礎機能作成完了
- 出力関数呼び出し時刻出力機能実装
- 出力レベル設定機能実装

2020 年 12 月 24 日

- 出力関数呼び出し行数出力機能実装
- 出力関数呼び出しファイル名前出力機能実装

## Works Cited
