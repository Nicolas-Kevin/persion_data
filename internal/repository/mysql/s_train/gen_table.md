#### go_gin_api.s_train 
演训表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | train_name | 演训名称 | varchar(10) |  | NO |  |  |
| 3 | train_data | 时间年月日 | varchar(10) |  | NO |  |  |
| 4 | train_org | 地点（市） | varchar(10) |  | NO |  |  |
| 5 | train_num | 参赛人员数量 | int(11) |  | NO |  |  |
| 6 | train_project | 科目 | varchar(10) |  | NO |  |  |
| 7 | train_place | 组织方 | varchar(10) |  | NO |  |  |
| 8 | train_time | 宣传图片，视频 | text |  | YES |  |  |
| 9 | deleted |  | tinyint(1) |  | NO |  | 0 |
