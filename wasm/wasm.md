# WebAssemblyって何？

**Wiki引用**  
Webブラウザのクライアントサイドスクリプトとして動作するプログラミング言語である  
ブラウザ上でバイナリフォーマット形式で実行されるのが特徴  

最初の目標はC, C++からのコンパイルサポートである  
CやC++で書かれたプログラムがWebブラウザで動くという事である  

Rustもサポートするらしい  
wasmはブラウザサポートが必要である  
2017年11月には全ての主要ブラウザが対応している(Chrome, Firefox, Safari, Edge)  


wasmには中間コードが存在する  
リニアアセンブリバイトコードと呼ばれる  


```
int factorial(int n) {
    if n == 0
        return 1;
    else
        return n * factorial(n-1);
}
```

```
get_local 0
i64.const 0
i64.eq
if i64
    i64.const 1
else
    get_local 0
    get_local 0
    i64.const 1
    i64.sub
    call 0
    i64.mul
end
```

```
20 00
42 00
51
04 7e
42 01
05
20 00
20 00
42 01
7d
10 00
7e
0b
```

wasmを使う、もしくは関連しているツールは様々ある  

**Emscripten**  
元々asm.js向けであったがwasmにも対応した  

Emscriptenのコンパイルでは、clang(|fastcomp-clang)、LLVM(|fastcomp)、binaryenを使用している  

**GCC asm.js backend**  
asm.js及びWebAssemblyに対応している  

**WebAssembly Studio**  
WebAssemblyを開発する為のIDEである  
C, Rustに対応している  
https://webassembly.studio/  

**Qt for WebAssembly**  
なんと、QtがWebAssemblyに対応している  
http://blog.qt.io/blog/2018/05/22/qt-for-webassembly/  


## MDN web docs

ここからはMDN web docsを参考にしつつ理解を深めていく  
まずはEmscriptenをインストールしてみる  
cmakeないとダメ``sudo apt-get install cmake``  


```
$ git clone https://github.com/juj/emsdk.git; cd emsdk
$ ./emsdk install --build=Release sdk-incoming-64bit binaryen-master-64bit
$ ./emsdk activate --global --build=Release sdk-incoming-64bit binaryen-master-64bit
$ source ./emsdk_env.sh
```

*hello.c*  

```c
#include <stdio.h>

int main(int argc, char ** argv) {
    printf("Hello WebAssembly!\n");
}
```

https://github.com/kripken/emscripten/issues/6012

```
$ emcc hello.c -s WASM=1 -o hello.html
$ emrun hello.html
```


## Qt for WebAssembly

```
$ git clone -b wip/webassembly https://github.com/qt/qtbase.git
$ git clone -b wip/webassembly https://code.qt.io/qt/qtdeclarative.git
$ git clone -b wip/webassembly https://code.qt.io/qt/qtwebsockets.git

$ cd qtbase
$ ./configure -xplatform wasm-emscripten -developer-build -release -static -no-thread -nomake tests -nomake examples  -no-dbus -no-headersclean -system-libpng -no-ssl -no-warnings-are-errors
$ sudo make
```

* rootでactivateし直す
* source ./emsdk_env.shする
* type em++できればOK









