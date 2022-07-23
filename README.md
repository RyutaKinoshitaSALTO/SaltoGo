# 変数
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

# 配列型

配列型の型名は[要素数]要素の型のように宣言する。型名に続けて{}で囲むことで初期値を設定することもできる。
``` 
var a = [3]string{"name", "from", "age"}
``` 

# 関数
関数はfunc 関数名(引数の定義) 戻り値型 { 関数本体 }のように定義する。
``` 
func num(x, y int) int {
    return x + y
}
``` 

# if文
``` 
if x == 0 {
    fmt.Println("It's zero")
} else if x > 0 {
    fmt.Println("It's positive")
} else {
    fmt.Println("It's negative")
}
``` 
# for文

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

# マップ map
map[キーの型]要素の型という書き方で定義する。

``` 
// 定義
var m map[int]string


// 代入
m[1] = "au"
`m[2] = "docomo"
m[3] = "softbank"
``` 
# ポインタ
ポインタとは値型(value type)に分類されるデータ構造のメモリ上のアドレスと型の情報であり、ポインタ型は`*string`のように、ポインタを使って参照・操作する型の前に`*`を置くことで定義できる。

``` 
var ip *int // ポインタ型
var fp *float64
var sp *[2]string

var i int // 値型（int型）
p := &i // ポインタ型
``` 

# 構造体　struct
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
https://qiita.com/katsuomi/items/d1e6625ae9a5b663e11f
https://qiita.com/shiei_kawa/items/eddf48287455380f618f
https://www.udemy.com/course/go-fintech/

