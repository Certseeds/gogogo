show databases;
create SCHEMA IF NOT EXISTS `background`
    DEFAULT CHARACTER SET utf8mb4 collate utf8mb4_unicode_ci;
USE `background`;

create TABLE IF NOT EXISTS `background`.`github_user`
(
    `id`               BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT ' 自增 id',
    `login`            VARCHAR(50)  NOT NULL UNIQUE COMMENT '用户名',
    `githubNodeId`     VARCHAR(128) NOT NULL UNIQUE COMMENT 'github的随机用户字符串ID',
    `githubId`         BIGINT(20)   NOT NULL UNIQUE COMMENT 'github的随机用户数字ID',
    `githubBlog`       VARCHAR(128) NOT NULL COMMENT 'github用户的blog链接',
    `twitterUsername` VARCHAR(128) NOT NULL COMMENT 'github用户的twitter链接',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC),
    UNIQUE INDEX `name_UNIQUE` (`login` ASC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create TABLE IF NOT EXISTS `background`.`github_follow`
(
    `id`  BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'follow关系的id',
    `src` BIGINT(20) NOT NULL COMMENT 'follow关系的关注者,为github_user内的id',
    `dst` BIGINT(20) NOT NULL COMMENT 'follow关系的关注者,为github_user内的id',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC),
    UNIQUE INDEX `connection_UNIQUE` (`src`, `dst`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create TABLE IF NOT EXISTS `background`.`github_follow_history`
(
    `id`  BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '历史follow关系的id',
    `src` BIGINT(20) NOT NULL COMMENT 'follow关系的关注者,为github_user内的id',
    `dst` BIGINT(20) NOT NULL COMMENT 'follow关系的关注者,为github_user内的id',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


show databases;
show tables;