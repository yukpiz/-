# CloudWatch LogsをリアルタイムにLogDNAに送信する

この記事は以下の順序で組み立てられています。

1. CloudWatch LogsのイベントをLambdaで受け取る
2. LogDNAにログ情報を送信する
3. (おまけ) S3にアーカイブを設定する

利用する技術スタックは以下になります。

- AWS(CloudWatch Logs, Lambda, S3)
- Serverless Framework
- LogDNA


## 1. CloudWatch LogsのイベントをLambdaで受け取る

まずはCloudWatch Logsのイベント時に発火するLambda関数を作成していきます。
今回AWSのリソースの作成にはServerless Frameworkを使っていきます。

まずはServerless Framework CLIをインストールします。

```bash
$ npm install -g serverless
$ serverless --version
Framework Core: 1.57.0
Plugin: 3.2.2
SDK: 2.2.1
Components Core: 1.1.2
Components CLI: 1.4.0
```







