# mruby-embed

## go-mruby

* `go get`ではインストールできないので、自分でGOPAHT以下にcloneしてくる
* makeを実行すればvendor以下にmrubyを落としてセットアップ完了
```
$ cd ${GOPATH}/src/github.com/mitchellh/go-mruby
$ make
```

## クロスビルド

 * `https://github.com/matsumotory/mruby-cross-compile-on-mac-osx`を使ってWindows用ライブラリ等を作成してみたけどビルドできなかった・・・。
```
$ GOOS=windows GOARCH=amd64 go build -o out/test.exe
# github.com/mitchellh/go-mruby
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:46:33: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:62:40: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:64:42: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:113:46: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:126:47: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:137:45: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:155:51: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:209:45: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:286:45: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:300:47: undefined: MrbValue
../../gocode/src/github.com/mitchellh/go-mruby/decode.go:300:47: too many errors
```


## mruby用インスタンス生成

* GO側でmrubyのHashインスタンスを生成したい
* go-mruby側にはHashインスタンスを生成する実装はなさそう
* hashを生成するだけのmrubyスクリプトを呼ぶのがよさそう
