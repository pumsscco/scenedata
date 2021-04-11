CREATE TABLE `NPCInfo` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `scene` varchar(4) NOT NULL COMMENT '场景',
 `section` varchar(8) NOT NULL COMMENT '区块',
 `npc_id` varchar(10) NOT NULL COMMENT 'NPC ID',
 `model` varchar(16) NOT NULL COMMENT '模型',
 `func` varchar(16) NOT NULL COMMENT '函数调用',
 `init_coor` varchar(32) NOT NULL COMMENT '初始坐标',
 `coor2` varchar(32) NOT NULL COMMENT '坐标2',
 `two_int` varchar(8) NOT NULL COMMENT '两个用途不明的整数',
 `unk_float` float NOT NULL COMMENT '用途不明的浮点数',
 `attr` varchar(1024) NOT NULL COMMENT '详细信息',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='场景NPC信息'