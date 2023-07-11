package main

import "testing"

func Help() {
	goMain([]string{"./cwf", "--help"})
	// main cwf [オプション] <県名>

	// オプションをつけずに実行した場合は，現在の天気を出力する

	// オプション
	// 		-w, --week    このオプションはその週の天気予報を出力します．
	// 		-h, --help    このメッセージを出力して終了する
	// 		-v, --version バージョンを出力して終了する
	// 変数
	// 		県名    調べたい都道府県を指定する．
}

func Test_noOption_noCity(t *testing.T) {
	goMain([]string{"./cwf"})
	// main cwf [オプション] <県名>

	// オプションをつけずに実行した場合は，現在の天気を出力する

	// オプション
	// 		-w, --week    このオプションはその週の天気予報を出力します．
	// 		-h, --help    このメッセージを出力して終了する
	// 		-v, --version バージョンを出力して終了する
	// 変数
	// 		県名    調べたい都道府県を指定する．

}
