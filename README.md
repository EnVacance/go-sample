# go-sample
Goの開発環境セットアップ

■Install Go Lang
公式はバイナリを公式サイトからダウンロードする方法を推奨している。

インストール後にcommand lineで “go version”を叩いてセットアップ完了を確認。

“go env”でpath等の設定を確認
ex. GOPATH="/Users/polerbear/go"


■Workspace setup
Go path配下に公式が定めている(?)workspaceの構造のdirectory構造を作成。

polerbear/go
  - bin
  - src
    - github.com
      - envacance(my github user name)
        - go_crash_course(an example project)
        - list of repos goes here…

■Go getのテスト
ライブラリのインストールコマンドを動作させてテスト。
pip install的な。

コマンドラインのどこでも良いので、以下のコマンドを叩く。
“go get github.com/aws/aws-sdk-go/aws”

先ほど作成したWorkspaceのgo配下にpkg dirが作成され、その配下に
色々作成されていることを確認。


■単純なプロジェクト&コードの実行
以下のdirにmain.goを作成
src/github.com/<my github user name>/sample_project_name/01_hello/main.go

適当なコードを書く

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}


01_helloへコマンドラインで移動して”go run main.go”
Hello Worldが表示されることを書くに。
