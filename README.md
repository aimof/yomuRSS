# yomuRSS: rss reader

自作のrssリーダーです。
暇なときに改造します。

## 使い方

```
go install github.com/aimof/yomuRSS/cmd/yomurss
```

yomuRSS専用のディレクトリを作成し、以下の設定を記載。

```sh
# shellの設定ファイルに、
export YOMUDIR=/path/to/dir
```

* `$YOMUDIR/articles/` のディレクトリを作成
* `$YOMUDIR/target.txt` を作成して、改行区切りで取得したいフィードのURLを記載。

```sh
yomurss get
```

このコマンドを入力すると `target.txt` で指定したURLから記事情報を取得して、 `$YOMUDIR/articles/` 内にjsonファイルが作成されます。

```sh
yomurss
```

このコマンドでTUIが起動し、記事を見られます。
矢印キーで上下移動、 enterで記事選択、qで終了します。