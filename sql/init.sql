CREATE DATABASE /*!32312 IF NOT EXISTS*/`message_bot` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `message_bot`;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `tg_message` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `chat_id` bigint(20) NOT NULL,
                           `u_id` bigint(20) NOT NULL COMMENT '用户id',
                           `is_command` tinyint(1) NOT NULL COMMENT '是否是命令',
                           `command` varchar(128) NOT NULL COMMENT '命令类型',
                           `text_content` text COMMENT '文本内容',
                           `receive_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `send_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '消息发送时间',
                           `template_id` bigint(20) NOT NULL,
                           `reply_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `reply_text` text,
                           `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           PRIMARY KEY (`id`),
                           KEY `u_id` (`u_id`),
                           KEY `chat_id` (`receive_time`),
                           KEY `create_time` (`send_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `message` */

/*Table structure for table `template` */

DROP TABLE IF EXISTS `template`;

CREATE TABLE `template` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `command` varchar(128) NOT NULL,
                            `reg` varchar(128) NOT NULL COMMENT '正则表达式',
                            `reply` text NOT NULL,
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                            PRIMARY KEY (`id`),
                            KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;