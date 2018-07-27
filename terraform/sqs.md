
*{env}/env.tf*

```tf
module "sqs"
    source = "../modules/sqs"
    app = "${var.app}"
    env = "${var.env}"
```

*modules/variables.tf*

```tf
variable "app" { description = "app name" }
variable "env" { description = "dev, staging, production, etc." }
```

name: キューの名称  
name_prefix: 接頭辞  
visibility_timeout_seconds: キューの可視性タイムアウトの秒数 0-43200(12時間)  
message_retention_seconds: SQSがメッセージを保持する秒数 60-1209600  
max_message_size: メッセージのサイズ上限 1024-262144(byte)  
delay_seconds: キュー内の全てのメッセージが遅延する秒数 0-900  
receive_wait_time_seconds: ロングポーリング時の待機時間  
policy: SQSキューのJSONポリシー  
redrive_policy: デッドレターキューを設定するJSONポリシー  
fifo_queue: FIFOキューを指定する場合の真偽値(true/false)  
content_based_deduplication: FIFOキューの重複除外機能  
kms_master_key_id: なんかのキーID  
kms_data_key_reuse_period_seconds: KMSで暗号化するときに使う時間設定  
tags: キューに割り当てるタグ  




