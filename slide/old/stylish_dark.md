<style>
.margin-list20 li {
  margin-top: 20px;
  font-size: 45pt;
}

.cl-dg {
  color: darkgray;
}

.slide {
  margin-top: 20px;
}

.reveal pre code {
  display: inline;
  margin: 0 2px;
  padding: 1px 3px;
}
</style>

<div class="title1">これはタイトル</div>
<div class="title2">これはサブタイトル</div>
<div class="title3">これはサブサブタイトル</div>

- - -

<div class="title1">長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル</div>

- - -

カレーの作り方  
野菜を切って炒めます。  
鍋に野菜を入れ、火が通ったら肉を入れます。  

全体に火が通ったら、水を入れ沸騰するまで待ちます。  
最後に適量のカレールーを入れて、溶かして完成です。  

How to make curry.  
Cut the vegetables and stir it.  
Put the vegetables in the pot and put the meat when the fire passes.  

When the whole fire is on, add water and wait until it boils.  
Finally put an appropriate amount of curry roux,  
it is completed by melting.  

- - -

* Go言語
* Python
* Ruby
* Java

- - -

|#|Name|Age|Like|Check|
|:---|:---|---:|:---|:---:|
|1|yukpiz|28|vim|o|
|2|yukpiz|29|golang|o|
|3|yukpiz|30|aws|o|
|4|yukpiz|31|java|x|
|5|yukpiz|32|ruby|x|

- - -

十字路を``右``に曲がると、``コンビニ``がありますので、  
コンビニを過ぎたら``左``に曲がってください。  

そのまま進むと``信号``があるので、信号の``左側``の``白い建物の2F``です。  

- - -

```go
package main

import(
	"fmt"
)

func main() {
	n := 150
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz!")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
```




