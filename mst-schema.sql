CREATE TABLE `MstInfo` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `scene` varchar(4) NOT NULL COMMENT '场景',
 `section` varchar(8) NOT NULL COMMENT '区块',
 `mst_id` varchar(10) NOT NULL COMMENT '敌人队伍ID',
 `model` varchar(4) NOT NULL COMMENT '模型',
 `init_coor` varchar(32) NOT NULL COMMENT '初始坐标',
 `coor2` varchar(32) NOT NULL COMMENT '坐标2',
 `fix1` tinyint(1) NOT NULL COMMENT '第一个固定的1',
 `coor3` varchar(32) NOT NULL COMMENT '坐标3',
 `mst_num` tinyint(1) NOT NULL COMMENT '敌人数量',
 `mst_list` varchar(32) NOT NULL COMMENT '敌人详表',
 `attr` varchar(1024) NOT NULL COMMENT '详细信息',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='场景敌人信息'