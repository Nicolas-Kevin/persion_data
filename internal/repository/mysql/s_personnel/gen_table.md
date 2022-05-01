#### go_gin_api.s_personnel 
人员表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint(8) unsigned | PRI | NO | auto_increment |  |
| 2 | user_name | 姓名 | varchar(10) |  | NO |  |  |
| 3 | user_age | 年龄 | tinyint(1) |  | NO |  | 0 |
| 4 | user_sex | 性别（0：女，1：男） | tinyint(1) |  | NO |  | 0 |
| 5 | user_date | 出生年月日 | varchar(10) |  | NO |  |  |
| 6 | political_out_look | 社会面貌：0 :普通公民，1：团员，2：党员 | varchar(1) |  | NO |  |  |
| 7 | company | 工作单位 | varchar(20) |  | NO |  |  |
| 8 | position | 职称/等级 | varchar(20) |  | NO |  |  |
| 9 | major | 专业 | varchar(20) |  | NO |  |  |
| 10 | user_level | 级别 | varchar(20) |  | NO |  |  |
| 11 | user_height | 身高cm | int(1) |  | NO |  |  |
| 12 | user_weight | 体重cm | int(1) |  | NO |  |  |
| 13 | shoe_size | 鞋号cm | int(1) |  | NO |  |  |
| 14 | head_num | 头围cm | int(1) |  | NO |  |  |
| 15 | blood_type | 血型cm | int(1) |  | YES |  |  |
| 16 | id_num | 身份证号 | varchar(18) |  | NO |  |  |
| 17 | bank_num | 银行卡号 | varchar(19) |  | YES |  |  |
| 18 | phone_num | 电话号 | varchar(11) |  | NO |  |  |
| 19 | user_team | 部队编号 | bigint(8) |  | NO |  | 0 |
| 20 | train_status | 演训情况 | varchar(255) |  | YES |  |  |
| 21 | reward_status | 奖励情况 | varchar(255) |  | YES |  |  |
| 22 | deleted |  | smallint(6) |  | YES |  | 0 |
