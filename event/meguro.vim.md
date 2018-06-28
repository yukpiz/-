# vital.vim(Web.HTTP, Web.JSON)を使ってみる


### 導入

``:call dein#add('vim-jp/vital.vim')<CR>``して``:call dein#install()<CR>``するだけ  

### とりあえず叩いてみる

```vimscript

let vital = vital#of('vital')
let Http = vital.import('Web.HTTP')
let Json = vital.import('Web.JSON')

let response = Http.get('{URL}')
let json_string = response['content']
" https://github.com/thinca/vim-prettyprint
PP Json.decode(json_string)

```

簡単すぎた


### お天気を取ってみる

```vimscript

let appid = 'xxxxxxxxxxxxxxxxxxxxxxxx'

let params = {'q': 'Meguro,jp', 'appid': appid}
let response = Http.get('http://api.openweathermap.org/data/2.5/weather', params)
let json_string = response['content']
PP Json.decode(json_string)

```
