-- 用户表
create table if not exists user
(
    id            bigint auto_increment comment 'id' primary key,
    user_account  varchar(256)                           not null comment '账号',
    user_password varchar(512)                           not null comment '密码',
    union_id      varchar(256)                           null comment '微信开放平台id',
    open_id       varchar(256)                           null comment '公众号openId',
    username      varchar(256)                           null comment '用户昵称',
    user_avatar   varchar(1024)                          null comment '用户头像',
    user_profile  varchar(512)                           null comment '用户简介',
    user_role     varchar(256) default 'user'            not null comment '用户角色：user/admin/ban',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete     int          default 0                 not null comment '是否删除',
    index idx_unionId (union_id)
) comment '用户';


-- 题目表
create table if not exists question
(
    id           bigint auto_increment primary key,
    title        varchar(512)                       null comment '标题',
    content      text                               null comment '内容',
    tags         varchar(1024)                      null comment '标签列表（json 数组）',
    answer       text                               null comment '题解',
    submit_num   int      default 0                 not null comment '题目提交数',
    accepted_num int      default 0                 not null comment '题目通过数',
    judge_case   text                               null comment '判题用例（json 数组）',
    judge_config text                               null comment '判题配置（json 对象）',
    thumb_num    int      default 0                 not null comment '点赞数',
    favour_num   int      default 0                 not null comment '收藏数',
    user_id      bigint                             not null comment '创建用户 id',
    create_time  datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time  datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete    int      default 0                 not null comment '是否删除',
    index idx_userId (user_id)
) comment '题目';

-- 题目提交表
create table if not exists question_submit
(
    id            bigint auto_increment comment 'id' primary key,
    code_language varchar(128)                       not null comment '编程语言',
    code          text                               not null comment '用户代码',
    judge_info    text                               null comment '判题信息（json 对象）',
    status        int      default 0                 not null comment '判题状态（0 - 待判题、1 - 判题中、2 - 判题结束）',
    question_id   bigint                             not null comment '题目 id',
    user_id       bigint                             not null comment '创建用户 id',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete     int      default 0                 not null comment '是否删除',
    index idx_questionId (question_id),
    index idx_userId (user_id)
) comment '题目提交';