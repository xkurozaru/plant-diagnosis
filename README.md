# Plant Disease Diagnosis Service

## Requirements
* docker
* make
* go-migration

## Getting Start
1. `make up` で起動します
2. `localhost` にアクセスします
3. Memberユーザーを作成するには *新規登録* から, Adminユーザーを作成するには `sign-up-admin.rest` のリクエストを送信してください
4. ログインIDとパスワードを入力してログインします
5. 識別モデルを作成するには `localhost/admin` から, もしくは `create-model.rest` のリクエストを送信してください
6. 葉表画像をアップロードして病害を診断できます

## 終了時
`make down` でアプリケーションを終了します

## リモート
* Google Compute Engine にデプロイしています
* テストユーザー `LoginID:"test" Password:"test1234"`
