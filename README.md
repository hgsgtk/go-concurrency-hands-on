# 概要
書籍『達人プログラマ』では、知識のポートフォリオに投資することの重要性、その方法の一つとして一年に一個新しいプログラミング言語を学ぶことを推奨しています。
我々は知らず知らずのうちにプログラミング言語の背景にあるパラダイム・言語機能/仕様に影響を受け、その知識ポートフォリオの幅の発想力をもって、課題解決に日々取り組んでいると言えると思います。
手札は大いに越したことはありません。一枚しか手札を持っていない人と、複数枚手札をもって選べるのでは、課題解決の腕力は一段変わってくるはずです。

今回は、goroutine/channelについて入門する会です。GoはGoogleがマルチコア時代のコンピュータにおけるスケール問題に対する解決策として開発された言語であり、その結果がgoroutineやGCの実装だったりします。

一つの手札を増やす時間にしましょう。

# お品書き
- 並行処理と並列処理
- goroutineとは
- channelによるデータ競合回避
- 更に手を動かして理解する

# 並行(処理)と並列(処理)
まず、並行処理・並列処理の違いは何でしょうか？これらは同じものではありません。この違いを認識することが、ある種「お前はGopherなのか？」というリトマス試験紙的Questionです。

- 並行: Concurrency 同時にいくつかの**質の異なる**ことを扱うこと
- 並列: Parallelism 同時にいくつかの**質の同じ**ことを扱うこと

ref: [Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism) by Rob Pike

## 並列と並行の違い
TODO Rob Pikeのスライドから図を引用、解説

# goroutine（ゴールーチン）とは
## 概要
- Concurrencyを実現するためのgoroutine
- 複数のgoroutine一つ一つに役割を与えて分業
- 軽量なスレッドみたいなもの、Linux・Unixのスレッドよりコストが低い
- 一つのスレッドの上で複数のgoroutineが動く

TODO Linux・Unixのスレッドとの違い

## Try! use goroutine
`go`キーワードをつけて関数を呼び出すことでgoroutineを作ることができます。

TODO 出力する文字列でいぇいいぇい、時間差とかつける

# channelによるデータ競合回避
## データ競合
TODO 競合状態 race condition について
順序が保証されない
変更・参照が競合するという問題

## goroutine間でのデータのやり取り
### 共有変数でやる場合
TODO 共有変数名でやる場合、つらい

### channelを使った通信
`Do not communicate by sharing memory; instead, share memory by communicating`

TODO 出自

TODO Channel図

その他にもロックを取る方法がある

## channelとは
- 送受信できる型を定義する
- バッファあり・無し
- 送受信時の処理ブロック

## channel基本文法
- 初期化
- 送信
- 受信

ファーストクラスオブジェクトのため、変数に入れる・引数・戻り値ができる

TODO コード例

### Try! use channel

TODO

 ### 複数のchannelから同時に受信したい
`select`キーワードを使う

### Try! use select

TODO

### 標準パッケージ内での実例
- time.Afterの例

TODO 

### 単方向チャネル
- 型（受信専用・送信専用）

## Try! 単方向チャネル

TODO 

### concurrencyを実現する
- for-selectパターン
   - goroutineごとに無限ループを作る
   - メインのgoroutineはselectで結果を受信

#### Try! for-selectパターン

# syncでのデータ競合回避
## usecase
- どういうときに使う？

## sync.Mutexによるロック

### Try! ロック
コード例

## 書き込み・読み込みロック

### Try! 書き込み・読み込みロック
コード例

## 複数のgoroutine待機

# Context

## Contextによるキャンセル処理
goroutineをまたいだ処理のキャンセル

## タイムアウト処理


### Try! Context

TODO

# 更に手を動かして理解する
## 題材: 平行なディレクトリ走査
TODO

## 題材: TCPサーバー
TODO

## 題材: チャットサーバ
TODO
