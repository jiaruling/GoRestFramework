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
)comment '学生';

INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639541712, 1639541712, 123, '张三', 18, '大二');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639541743, 1639541743, 123, '李四', 19, '大三');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639541763, 1639541763, null, '王五', 18, '大四');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639551828, 1639551828, null, '赵六', 21, '研一');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639551865, 1639551865, null, '钱七', 22, '研二');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639551876, 1639551876, null, '孙八', 23, '研三');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639551945, 1639552088, null, '赵十', 29, '博三');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639552665, 1639552665, null, '张三丰', 25, '社会人');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639552688, 1639552688, null, '张无忌', 27, '武当');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639556192, 1639556192, null, '张无忌', 27, '武当');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639646516, 1639646516, 123, null, 99, '养老院');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639646635, 1639646635, 123, null, 88, '养老院');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639646830, 1639646830, 123, null, 77, '养老院');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639647037, 1639647037, null, null, 66, '养老院');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639647097, 1639647097, null, null, 66, '养老院');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639647597, 1639647597, null, null, 26, '养老院123');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639647697, 1639647697, null, null, 26, '养老院123');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639649868, 1639649868, null, '1999', 31, '华山');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639719977, 1639719977, null, '曹操', 108, '三国');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639719990, 1639719990, null, '张飞', 108, '三国');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639719998, 1639719998, null, '关羽', 108, '三国');
INSERT INTO imooc.student (created_at, updated_at, deleted_at, name, age, class) VALUES (1639719998, 1639720314, null, '刘字长', 108, '三国');