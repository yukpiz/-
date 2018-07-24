# Firebase Realtime Database

* https://github.com/firebase/firebase-admin-go

## 秘密鍵ファイルを用意する

Firebaseコンソールのプロジェクト設定からダウンロード可能  

## データベースに接続する

TODO  

## データベースを操作してみる

1. NewRefにてデータベース参照先をパスで指定する
2. 順序制約によってソートされる
3. フィルタリング制約によって抽出される

#### まずはtype Refを学ぶ

* Child
単純に指定されたパスへの子要素への参照を返す  

データベースはこのようになっているとする  

```
- root
  - messages
    - KngFsloJ1jsNF3N6sNM
      - KngFuJsnfhlIK84HyBE
        * message: "(」・ω・)」うー!"
        * sent_at: 1532404800
        * type: 1
```

連結して潜っていけます  

```go
ref := client.NewRef("root/messages/KngFsloJ1jsNF3N6sNM/KngFuJsnfhlIK84HyBE/message")
```

連結せずにメソッドチェーン  

```go
ref := client.NewRef("root").Child("messages").Child("KngFsloJ1jsNF3N6sNM").Child("KngFuJsnfhlIK84HyBE").Child("message")
```

存在しないキーや階層でも存在するときと同じように参照が返ってきます  

```go
ref := client.NewRef("uhh----nya-----")
```




```go
var firebaseMessages FirebaseMessages

//interface{}を介さず直接構造体にぶっこむ
client.NewRef("messages/u9_mu61").Get(context.Background(), &firebaseMessages)

//子要素で並び替えて取得(昇順)
query := client.NewRef("messages/u9_mu61").OrderByChild("sent_at")
results, _ := query.GetOrdered(context.Background())
fmt.Printf("len: %d\n", len(results))
for _, result := range results {
    var firebaseMessage FirebaseMessage
    result.Unmarshal(&firebaseMessage)
    fmt.Printf("%s: %d\n", result.Key(), firebaseMessage.SentAt)
}
```

