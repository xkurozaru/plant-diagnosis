# Plant Disease Diagnosis Service
## Getting Start
1. `make up` で起動します
2. migrationが済んでいない場合は `make migrate-up` でマイグレーションします
3. `localhost` にアクセスします
4. localで立ち上げた場合にはユーザーが居ないので、右上の **新規登録** からユーザーを登録します
5. 葉の画像から病害を識別できます

## 終了時
`make down` でアプリケーションを終了します

## リモート
* Google Compute Engine にデプロイしています
* テストユーザー `LoginID:"test" Password:"test1234"`
