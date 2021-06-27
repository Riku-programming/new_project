実装するリスト  

csvデータを分けてReplace関数のReplaceに流し込めるようにする
ユーザーに質問する機能  


処理の流れ

main関数が実行  
↓  
readHtml.Replace関数が実行  
↓  
readCsvToObject関数が呼ばれる(どのデータを変数に入れるかを選択してる)  
↓  
ReplaceToHtmlがhtmlの変換したい文字列と変換する値(csvに書かれたデータ)に変換する  
↓  
ConvertHtmlToPdfは変更しない(ほぼ完成)