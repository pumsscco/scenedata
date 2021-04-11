CREATE TABLE `GameObject` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `scene` varchar(4) NOT NULL COMMENT '场景',
 `section` varchar(8) NOT NULL COMMENT '区块',
 `kind` tinyint(2) NOT NULL COMMENT '种类',
 `gob_id` varchar(10) NOT NULL COMMENT '物品ID',
 `path` varchar(64) NOT NULL COMMENT '模型路径',
 `model` varchar(16) NOT NULL COMMENT '模型',
 `texture` varchar(16) NOT NULL COMMENT '贴图',
 `init_coor` varchar(32) NOT NULL COMMENT '初始坐标',
 `coor2` varchar(32) NOT NULL COMMENT '坐标2',
 `func` varchar(16) NOT NULL COMMENT '函数',
 `triple_int` varchar(12) NOT NULL COMMENT '三个用途不明的整数',
 `unk_float` float NOT NULL COMMENT '用途不明的浮点数',
 `attr` varchar(1024) NOT NULL COMMENT '详细信息',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='场景物体信息'