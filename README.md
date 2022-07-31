# Go言語とは
Go言語とは、検索エンジンとして馴染み深いGoogleが開発した静的型付け言語である。

#### 静的型付け言語とは
変数などのデータ型の宣言があらかじめ必要なプログラミング言語
例えば、数値型(int)であったり文字列型(string)といったデータ型の宣言を
行わなければならない言語である

#### 他の代表的な静的型付け言語
　C言語
　Java
　Scala
　Swift
 
 # Go言語を使用するメリット
#### 初学者でも理解しやすい
CやJavaとも構文が似ているため、これらの言語について知っていると習得しやすい言語になっている
また少ないコード実装できるため、初学者も理解しやすい
例:Go言語、C言語でHello Worldを出力 

##### Go言語
``` 
package main
import "fmt"
func main() {
　fmt.Printf("Hello World\n")
}
```
##### C言語
``` 
#include <stdio.h>
 int main(void) {
    printf("Hello, world.\n")
    return 0;
}
``` 

#### 処理速度が速い
この事情に関しては様々な要因があるが大きな要因として挙げられるのが並行処理について挙げられることが多い
並行処理を容易にするためにGo言語はコードの種類が極端に制限されており、内部の処理を簡易化しています
また、Go言語はプログラマーが意識せずとも自動的に複数処理を同時に行うよう設計されている

#### 安全性が高い
Go言語はコードがシンプルなので、複数のエンジニアが開発にかかわっても記述の仕方がぶれにくく、ミスが起きにくい仕様となっている
Go言語にはエラーが発生しやすい「ポインタ演算」の機能がそもそも備わっていないため、メモリの安全性が高いのもメリットである
Go言語は他の言語に比べて機能が少ない分、扱いやすく安全性が高いという特徴がある

# Go言語の得意分野について
#### Webサーバー・Webサービスの構築
Googleサービスとの親和性が高い
Linux、Mac、Windowsなど多くのOSに対応
大規模で高負荷なサービスの実装も得意
特にWebメディアやショッピングサイトなど、自社でプラットフォームを持っている企業で多く導入されている

#### スマートフォンアプリ開発
Go言語はエラーの検出性が高いため改修がしやすく、スマートフォン向けのアプリ開発にも向いている
「Go mobile」「Go Cloud」などのツールやWeb Application Frameworkを使えば、Web・アプリ用とコードを書き換える必要がなく、手軽にスマートフォンアプリの開発ができる
Androidはもちろん、iOSにも対応しているのが人気の理由

#### ドローン・IoT
ロボティクス・IoTフレームワーク「Gobot」が公開されており、ドローンのような高度な組み込み開発にも利用可能
Gobotにはネットワーク上のデバイスや複数のデバイス間で相互通信できる機能があるため、外出先から遠隔操作したり、センサーが感知した情報を他デバイスに伝えたりするIoT開発にも使用しやすいという特徴がある


# 基本的な書き方

###  変数
var [変数名] [変数の型]のように定義する。

``` 
var n int
n = 5

// int型の変数x,y,zを定義・代入<br>
var x, y, z int

x, y, z = 1, 3, 5

// int型の変数x,yとstring型のnameを定義<br>
var (
    x, y int
    name string
) 
```

型指定の必要がない為、以下の形で定義するのが一般的である。
``` 
var i = 1
var b = true
var f = 3.14
var s = "git"
``` 

### 配列型

配列型の型名は[要素数]要素の型のように宣言する。型名に続けて{}で囲むことで初期値を設定することもできる。
``` 
var a = [3]string{"name", "from", "age"}
``` 

### 関数
関数はfunc 関数名(引数の定義) 戻り値型 { 関数本体 }のように定義する。
``` 
func num(x, y int) int {
    return x + y
}
``` 

###  if文
``` 
if x == 0 {
    fmt.Println("It's zero")
} else if x > 0 {
    fmt.Println("It's positive")
} else {
    fmt.Println("It's negative")
}
``` 
###  for文

``` 
// 基本
for i := 0; i < 10; i++ {
    fmt.Println(i)
    i++
}


// break文
i := 0
for {
    fmt.Println(i)
    i++
    if i == 70 {
        break
    }
}


// 条件付き
i := 0
for i < 50 {
    fmt.Println(i)
    i++
}


// continue文
for i := 0; i < 100; i++ {
    if (i % 2 == 1) {
        continue
    }
    fmt.Println(i)
    i++
}
``` 

###  マップ map
map[キーの型]要素の型という書き方で定義する。

``` 
// 定義
var m map[int]string


// 代入
m[1] = "au"
`m[2] = "docomo"
m[3] = "softbank"
``` 
### ポインタ
ポインタとは値型(value type)に分類されるデータ構造のメモリ上のアドレスと型の情報であり、ポインタ型は`*string`のように、ポインタを使って参照・操作する型の前に`*`を置くことで定義できる。

``` 
var ip *int // ポインタ型
var fp *float64
var sp *[2]string

var i int // 値型（int型）
p := &i // ポインタ型
``` 

###  構造体　struct
構造体を使用するには一般的にtypeと組み合わせて新しい型を定義する。構造体は、structで定義された構造体に、typeを使って新しい型名を与えるという順序で定義する。
``` 
type Point struct {
    X int
    Y int
}

var pt Point
pt.X = 10
pt.X // == 10
``` 
# 参考文献

golangの基礎(初級編) → https://go-tour-jp.appspot.com/list
golangの基礎(中級編) → https://gopherdojo.org/
golangの応用編(上級編) → https://www.udemy.com/course/go-fintech/
ginフレームワークの基礎(初級編) → https://qiita.com/shiei_kawa/items/eddf48287455380f618f
ginフレームワークの基礎(中級編) → https://qiita.com/katsuomi/items/d1e6625ae9a5b663e11f
golangを環境関係なく実行 → https://go.dev/play/



