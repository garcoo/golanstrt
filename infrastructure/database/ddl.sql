/** 利用者のマスタ*/
CREATE TABLE m_user(
    user_id int(3) NOT NULL PRIMARY KEY
    ,user_name varchar(50) NOT NULL
    ,user_nickname varchar(50) NOT NULL
    ,chat_account_id int(8) NOT NULL
    ,activ_flg int(1) NOT NULL
);

/** 外部サービスのトークン管理*/
CREATE TABLE m_manage_servicetoken(
    service_id char(2) NOT NULL PRIMARY KEY
    ,service_token varchar(100) NOT NULL
    ,service_note varchar(50) NOT NULL
    ,activ_flg int(1) NOT NULL
);

/** メッセージを無視するIDを管理*/
CREATE TABLE m_chatwork_ignoreid(
    chat_account_id int(8) NOT NULL
);
