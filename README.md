# Plant Disease Diagnosis Service

## Requirements
* docker
* make
* go-migration

## Getting Start
1. `make up` で起動します
2. migrationが済んでいない場合は `make migrate-up` でmigrationを行います
3. `localhost` にアクセスします
4. Memberユーザーを作成するには *新規登録* から, Adminユーザーを作成するには `sign-up-admin.rest` のリクエストを送信してください
5. ログインIDとパスワードを入力してログインします
6. 識別モデルを作成するには `localhost/admin` から, もしくは `create-model.rest` のリクエストを送信してください
7. 葉表画像をアップロードして病害を診断できます

## 終了時
`make down` でアプリケーションを終了します

## リモート
* Google Compute Engine にデプロイしています
* テストユーザー `LoginID:"test" Password:"test1234"`
* http://plant-diagnosis.ddo.jp
