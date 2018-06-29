# AngularにSentryを導入する


**環境**  

* Angular5(@angular/* 5.2.0)
* sentry(https://sentry.io)


## 1. 何ができるか？

* 発生した例外を検知できる
* 例外は管理コンソールで確認できる
* OSやブラウザ等の環境情報が取れる
* スタックトレース
* issue化、statusやassign管理
* Slack通知

## 2. 導入

```
$ npm install raven-js --save
```


## 3. 初期化

*app.module.ts*  

```
import { NgModule, ErrorHandler } from '@angular/core';
import * as Raven from 'raven-js';

Raven.config('https://7452f16fe012490588736825100a374e@sentry.io/1234111').install();
export class RavenErrorHandler implements ErrorHandler {
  handleError(err:any) : void {
    Raven.captureException(err);
  }
};

@NgModule({
  providers: [
    { provide: ErrorHandler, useClass: RavenErrorHandler },
  ],
})
```

## 4. 環境を設定する

初期化までやってしまえば、例外のキャプチャは開始される  
適当な例外を発生させると、https://sentry.io に表示される  


ただ、このままだとローカルでの``ng serve``時などの例外も全て検知されてしまいます  
environmentsを定義して、stagingやproductionのみ検知する、等の対策が必要です  

#### 4-1. environemtnsを定義する

Angularではsrc/environments/environment.{env}.tsのように設定ファイルを置いておくと、  
``ng serve --stg``などで設定を切り替える事ができます  

今回は、dev/prodの2環境を用意して、prodのみエラー検知ができるようにします  

*src/environments/environemtn.prod.ts*  

```
export const environment = {
  production: false,
  development: true,
};
```

*src/environments/environemtn.dev.ts*  

```
export const environment = {
  production: false,
  development: true,
};
```

#### 4-2. ng serveやng buildで環境を切り替える

prod環境で実行、ビルドする場合  

```
$ ng serve --prod
$ ng build --prod
```

dev環境で実行、ビルドする場合  

```
$ ng serve --dev
$ ng build --dev
```

もちろんステージング等の環境を増やしていく事もできます  
オプションをつけない、つまりenvironment.tsをdev環境にしてしまう手もあります  

## べんり

本番環境で起きた不具合や例外をいち早く検知できるのは非常にありがたいです  
特にデータに依存するような例外だと、再現性がなくデバッグに非常にコストがかかります  

そういった例外を詳細に確認できたり通知できたりするのでかなり便利です  

