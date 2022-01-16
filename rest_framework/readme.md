# 技术选型

## web 框架

- gin
- beego
- echo
- Iris
- http【标准库】

## orm 框架

- mysql-driver
- sqlx
- gorm
- beego-orm

```sql
create table student
(
    id         int unsigned auto_increment
        primary key,
    created_at int unsigned null comment '创建时间戳',
    updated_at int unsigned null comment '更新时间戳',
    deleted_at int unsigned null comment '删除时间戳',
    name       varchar(64)  null comment '名字',
    age        tinyint      null comment '年龄',
    class      varchar(64)  null comment '班级'
)
    comment '学生';
```