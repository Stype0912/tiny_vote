CREATE DATABASE tiny_vote;
use tiny_vote;
DROP TABLE IF EXISTS `t_user_votes`;
CREATE TABLE `t_user_votes`
(
    `id`    int(10)     NOT NULL AUTO_INCREMENT,
    `name`  varchar(32) NOT NULL,
    `votes` bigint DEFAULT 0,
    primary key (id),
    unique (name)
) engine = innoDB
  default charset = utf8;

INSERT INTO `t_user_votes` (name)
VALUES ('Alice'),
       ('Bob'),
       ('Carol'),
       ('Daniel'),
       ('Epsilon'),
       ('Frank'),
       ('George');