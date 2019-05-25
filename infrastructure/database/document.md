# テーブル定義

### m_user
利用者のマスタ

| column             | type    | length  | NULL ABLE  | PK  |
|--------------------|---------|--------:|------------|-----|
| user_id            | int     | 3       |            | ◯  |
| user_name          | varchar | 50      |            |     |
| user_nickname      | varchar | 50      |            |     |
| chat_account_id    | int     | 8       |            |     |
| activ_flg          | int     | 1       |            |     |


### m_manage_servicetoken
外部サービスのトークン管理

| column             | type    | length  | NULL ABLE  | PK  |
|--------------------|---------|--------:|------------|-----|
| service_id         | char    | 2       |            | ◯  |
| service_token      | varchar | 100     |            |     |
| service_note       | varchar | 50      |            |     |
| activ_flg          | int     | 1       |            |     |


### m_chatwork_ignoreid
メッセージを無視するIDを管理

| column             | type    | length  | NULL ABLE  | PK  |
|--------------------|---------|--------:|------------|-----|
| chat_account_id    | int     | 8       |            |     |
