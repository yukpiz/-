<div style="font-size:120pt;">やさしいvim</div>


<img src="https://i.gyazo.com/ce86b03451c798c172d7498afe4f5c2f.png" style="border:none;background:none;width:500px;"/>

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

.reveal li code {
  font-size: 0.8em;
  margin: 0px 10px 0px 10px;
  padding: 5px 20px;
  background-color: #424242;
  color: #e2e2e2;
  border: 3px solid #8a8a8a;
  border-radius: 8px;
}

.reveal pre code {
  display: inline;
  margin: 0 2px;
  padding: 1px 3px;
  font-size: 2em;
  line-height: 1.3em;
}

.reveal pre.vim {
  background-color: #424242;
  border: 3px solid #8a8a8a;
  border-radius: 8px;
  padding: 20px 20px;
}

</style>


- - -

<div style="color:yellow;font-size:120pt;">WARNING</div>

<div style="font-size:40pt;">
できるだけ分かりやすく、実用的な面も含めてvimの良さをお伝えします。  
ただし、発言の保証が効かない場合があります。  
</div>

<br/>
<img src="https://i.gyazo.com/919022189384635187e6a1c3a545fd6a.jpg" style="width:900px;">

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
<div style="font-size:30pt;">VimConf 2018 Sponsor</div>

<img src="https://i.gyazo.com/8a02d4e82e446e3db3dbb302c7ddc7f2.png" style="background:none;width:1000px;"/>

<br/>
<div style="font-size:30pt;">採用もあるよ！</div>

- - -

<div style="font-size:100pt;">No Editor No Life.</div>
<div style="font-size:50pt">golang, java, ruby, python, nodejs, typescript, terraform... :-(</div>


- - -

<div class="margin-list20">
* 皆さんエディタ使ってますか？
* IDEでもエディタでも拘ってみるの面白いです
* 手に馴染んで、使いやすく、そしてモテモテに↑
* そんな魔法のツールをご紹介します
</div>

<img src="https://i.gyazo.com/728b68a385d068c0558e56f06f4dee6d.png" style="border:none;background:none;width:800px;"/>


- - -

<div style="font-size:100pt;">そもそもvimって</div>

<div class="margin-list20">
* viから派生した高機能テキストエディタ
* オランダ人プログラマーのブラムによって開発された(1988年)
* Vi IMproved(viの改良)である
* キーボードのみで操作される事が前提となっている
* 現在も活発に開発が行われている
* jobの非同期実行、ターミナル機能、ラムダ構文(VimL)などが追加
</div>

- - -

<div style="font-size:100pt;">Vimのすすめ(4 steps)</div>

- - -

<div style="font-size:100pt;">step++</div>

- - -

<div class="margin-list20">
* 多くのディストリビューションにプリインストールされています
* まずはターミナルで``$vim<Enter>``してみましょう
* 文字を入力しましょう [``i`` => ``aiueo``]
* 保存してみましょう [``<ESC>`` => ``:wq sample.txt``]
</div>

<img src="https://i.gyazo.com/02a620da015894b2d41120519ce76c38.gif" style="border:none;background:none;"/>

- - -

<div style="font-size:50pt;">
残念ながらこれだけでは効率よく開発ができません。<br/>
設定を書きましょう！
</div>

- - -

<div style="font-size:100pt;">step++</div>

- - -

<div class="margin-list20">
* ``${HOME}/.vimrc``は、vimの起動時に読み込まれる設定ファイル
* VimScriptと呼ばれる組み込みのスクリプト言語で記述します
* .vimrcで書けることはノーマルモードでもそのまま実行できます
* 行番号を表示してみましょう [``<ESC>`` => ``:set number``]
* 行番号を消してみましょう [``<ESC>`` => ``:set nonumber``]
* 設定の意味を調べよう [``<ESC>`` => ``:help number``]
</div>

<br/>
<div style="font-size:40pt;">0から書くのは大変、まずは誰かの.vimrcを参考にしよう！</div>
<img src="https://i.gyazo.com/4c32c90968e01c8cda0e666472c1f899.png" style="width:1400px;"/>

- - -

<div style="font-size:50pt;">
ここまで来ると思考のスピードに手が追いついてきます。<br/>
さらに良くするために、プラグインを入れてみましょう！
</div>

- - -

<div style="font-size:100pt;">step++</div>

- - -

<div class="margin-list20">
* プラグインを入れる為のパッケージマネージャーが必要
* 選択肢が多くありますが、拘りがなければ``dein.vim``が良い
* .vimrcに1行記述するだけで、プラグインの導入ができるようになる
* https://vimawesome.com/
</div>

<img src="https://i.gyazo.com/5d45685340e2010c7a0a69abf2ccdb38.png" style="width:1400px;"/>

- - -

<div style="font-size:50pt;">
プラグインが使えるようになると、<br/>
vimなしでは生きられなくなります、もう一息です<br/>
<br/>
次はプラグインを書いてみよう！
</div>

- - -

<div style="font-size:100pt;">step++</div>

- - -

<div class="margin-list20">
* VimScriptを使えば、独自のプラグインを作ったりできます
* 以下の関数を.vimrcに書いて、``:call FizzBuzz(100)<Enter>``しよう
</div>


```vim
function! FizzBuzz(num)
  for i in range(1, a:num)
    if i % 15 == 0
      echo "FizzBuzz!"
    elseif i % 5 == 0
      echo "Fizz"
    elseif i % 3 == 0
      echo "Buzz"
    else
      echo i
    endif
  endfor
endfunction
```

- - -
