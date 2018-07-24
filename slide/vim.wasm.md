<div style="font-size:120pt;">vim.wasmがすごい</div>
<br/>
<div style="font-size:60pt;">(WebAssemblyを勉強してみた！)</div>

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

.reveal pre code { display: inline; margin: 0 2px; padding: 1px 3px; }
</style>


- - -

<div style="font-size:50pt;">PROFILE</div>

![](https://pbs.twimg.com/profile_images/1009090611108982785/s13PS89H_400x400.jpg)

<div class="margin-list20">

* <span style="color:#4cd94c;">@yukpiz</span>(Twitter/Github)
* redish Inc(Engineer)
* \#<span class="cl-dg">golang</span> #<span class="cl-dg">vim</span> #<span class="cl-dg">crypko</span>
* \#<span class="cl-dg">AWS</span> #<span class="cl-dg">Rails</span> #<span class="cl-dg">Angular</span> #<span class="cl-dg">Android</span>
* \#<span class="cl-dg">Fishing</span> #<span class="cl-dg">Anime</span> #<span class="cl-dg">小型船舶操縦士1級</span>

</div>

- - -

<div style="font-size:80pt;"><b>redish Inc</b></div>

<img src="https://s3-ap-northeast-1.amazonaws.com/assets.redish.jp/redish/redish_logo_rgb2.png" style="border:none;background:none;width:210px;height:210px;margin-left:-15px;"/>
<img src="https://s3-ap-northeast-1.amazonaws.com/assets.redish.jp/redish/redish_qr.png" style="border:none;background:none;width:200px;height:200px;margin-left:30px;"/>

<div style="font-size:35pt;"><b>レストラン予約代行のコンシェルジュサービス</b></div>

<br/>
<div style="font-size:30pt;">Service Vision: レストラン体験を豊かに</div>
<div style="font-size:30pt;">Company Vision: 全ての繋がりを価値あるモノに</div>

<br/>
<div style="font-size:30pt;">採用もあるよ！</div>

- - -

<div style="font-size:80pt;">vim.wasm?</div>

<div class="margin-list20">

* <u>https://rhysd.github.io/vim.wasm/</u>
* vim/vimの実験的なWebAssemblyコンパイル
* テキストエディタがブラウザで動いている
* vim.jsではない、なにこれ？

</div>

<img src="https://i.gyazo.com/3a6540ef65c85a08a8afa13d2267f501.gif" style="border:none;background:none;">

- - -

<div style="font-size:80pt;">WebAssembly?</div>

<div class="margin-list20">

* 2015年6月に発表された
* ブラウザ上で実行可能なバイナリフォーマット
* C/C++, Rust, Go, C# → wasm
* 2017年11月に主要ブラウザが対応している

</div>

- - -

<div style="font-size:80pt;">C/C++ → emcc/em++ → wasm/js/html</div>

<div style="font-size:60pt;">
```c
#include <stdio.h>

int main(int argc, char ** argv) {
    printf("(」・ω ・)」うー!\n");
    printf("(/・ω ・)/にゃー!\n");
}
```
</div>

<div style="font-size:50pt;">↓</div>

<div style="font-size:60pt;">
```bash
$ emcc hello.c -s WASM=1 -o hello.html
$ emrun hello.html
```
</div>

<div style="font-size:50pt;">↓</div>

<img src="https://i.gyazo.com/5d5e35815f95f01b267b176bc84120fb.png" style="border:none;background:none;width:1200px;">

- - -

<div style="font-size:40pt;">Cで書いた関数をjavascriptで呼び出してみる</div>

<div style="font-size:40pt;">

```c
#include <stdio.h>
#include <emscripten/emscripten.h>

#ifdef __cplusplus
extern "C" {
#endif

int EMSCRIPTEN_KEEPALIVE Uhh() {
    printf("(」・ω ・)」うー!\n");
    return 1;
}

char EMSCRIPTEN_KEEPALIVE *Nya() {
    return "(/・ω ・)/にゃー!\n";
}

int EMSCRIPTEN_KEEPALIVE add(int a, int b) {
    return a + b;
}

#ifdef __cplusplus
}
#endif
```

</div>

<div style="font-size:40pt;">

```javascript
Module.ccall('Uhh', null, null, null);
// (」・ω・)」うー!

var result2 = Module.ccall('Nya', 'string', null, null);
console.log(result2); // (/・ω・)/にゃー!

var result3 = Module.ccall('add', 'number', ['number', 'number'], [3, 5]);
console.log(result3); // 8
```

</div>


- - -

<div style="font-size:80pt;">Golang for WebAssembly(Go 1.11+)</div>

Cだとちょっとつらいけど、GoだとライブラリのWebAssembly化捗るよね！  

<div style="font-size: 60pt;">
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello WebAssembly!")
}
```
</div>

<div style="font-size: 60pt;">
```bash
$ GOOS=js GOARCH=wasm go1.11beta1 build -o main.wasm main.go
```
</div>


- - -

<div style="font-size:80pt;">Qt for WebAssembly</div>

<div style="font-size:40pt;">qt/qtbase(wip/webassembly)</div>
<div style="font-size:35pt;">qmake環境構築間に合わず・・</div>

<img src="https://i.gyazo.com/ed97d4581c53f2ff58cd834e792148ba.gif" style="border:none;background:none;">


- - -

<div style="font-size:80pt;">まだまだこれから</div>

<div class="margin-list20">

* WebAssemblyがJavaScriptを置き換えるか？ == No
* WebAssemblyが何を目指しているのか、ドキュメントにしっかり書いてあった
* ちゃんと使ってみて勉強してみたら色々わかってきた
* これまでパフォーマンス面でJavaScriptで実現の大変だったところをサポート
* 新しい技術の1つとして色んな言語やライブラリのサポートが楽しみ
* なんかおもろい

</div>

- - -

<div style="font-size:80pt;">期待</div>

<div class="margin-list20">

* Qt(for WebAssembly)
* Vim(vim.wasm)
* OpenCV(Build cv-wasm.wasm)
* Golang(WebAssembly ("wasm") support)

</div>

- - -

<div style="font-size:80pt;">ご静聴ありがとうございました！</div>

<div style="font-size:45pt;">@yukpiz</div>

