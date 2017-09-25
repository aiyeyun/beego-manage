CREATE TABLE IF NOT EXISTS `manage`.`cl_admin` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `username` VARCHAR(20) NOT NULL COMMENT '用户名',
  `nickname` VARCHAR(30) NOT NULL COMMENT '昵称',
  `email` VARCHAR(50) NOT NULL COMMENT '电子邮箱',
  `password` CHAR(32) NOT NULL COMMENT '密码',
  `portrait` VARCHAR(250) NOT NULL COMMENT '头像',
  `salt` CHAR(6) NOT NULL COMMENT '密码调味品',
  `role` TINYINT(3) NOT NULL COMMENT '角色',
  `last_ip` VARCHAR(20) NOT NULL COMMENT '最后登录IP地址',
  `created` DATETIME NOT NULL COMMENT '注册时间',
  `update` DATETIME NOT NULL COMMENT '最后登录时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `un_username` (`username` ASC)  COMMENT '唯一 用户名',
  UNIQUE INDEX `un_email` (`email` ASC)  COMMENT '唯一 邮箱',
  UNIQUE INDEX `un_nickname` (`nickname` ASC)  COMMENT '唯一 昵称')
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '管理员';

CREATE TABLE IF NOT EXISTS `manage`.`cl_menu` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `pid` INT(11) NOT NULL DEFAULT 0 COMMENT '父节点ID',
  `url` VARCHAR(50) NULL DEFAULT NULL COMMENT 'url\n父节点 没有URL\n子节点 有URL',
  `name` VARCHAR(30) NOT NULL COMMENT '菜单名',
  `icons` VARCHAR(30) NULL DEFAULT NULL COMMENT '图标 class 名\n父节点有 图标\n子节点没有 图标',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否开启显示\n0=>关闭, 1=>显示',
  `sort` INT(11) NOT NULL DEFAULT 1 COMMENT '排序',
  `created` DATETIME NOT NULL COMMENT '创建时间',
  `update` DATETIME NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `un_name` (`name` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '菜单';

CREATE TABLE IF NOT EXISTS `manage`.`cl_menu_auth` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `role` INT(11) NOT NULL COMMENT '角色',
  `menu_id` INT(11) NOT NULL COMMENT '菜单表 主键ID',
  `created` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `fk_cl_memu_auth_1_idx` (`menu_id` ASC),
  CONSTRAINT `fk_cl_memu_id`
    FOREIGN KEY (`menu_id`)
    REFERENCES `manage`.`cl_menu` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '菜单权限表';

INSERT INTO `cl_admin` VALUES (1,'admin','夜云','644362887@qq.com','f040f52ff90579e9a28807c1c4c3d4be','/static/assets/img/faces/avatar.jpg','!@#$%^',1,'127.0.0.1','2017-09-24 20:46:46','2017-09-24 21:13:50');

INSERT INTO `cl_menu` VALUES (1,0,'/menu','菜单栏','menu',1,1,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(2,0,'/admin','管理员','person',1,2,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(3,1,'/index','菜单栏列表',NULL,1,1,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(4,1,'/form','添加菜单栏',NULL,1,2,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(5,2,'/index','管理员列表',NULL,1,1,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(6,2,'/form','添加管理员',NULL,1,2,'2017-09-23 23:04:44','2017-09-23 23:04:44'),(7,2,'/role','管理员权限',NULL,1,3,'2017-09-23 23:04:44','2017-09-23 23:04:44');


