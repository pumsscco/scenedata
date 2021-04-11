CREATE TABLE `PathWay` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `scene` varchar(4) NOT NULL COMMENT '场景',
 `section` varchar(8) NOT NULL COMMENT '区块',
 `pwy_id` varchar(10) NOT NULL COMMENT '路径ID',
  `fix1` tinyint(1) NOT NULL COMMENT '第一个固定的1',
`unk_float` float NOT NULL COMMENT '用途不明的浮点数',
 `coor_num` tinyint(2) NOT NULL COMMENT '坐标数量',
 `coor_list` varchar(1024) NOT NULL COMMENT '坐标列表',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='场景路径详情'