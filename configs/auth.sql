use `blog_service`;

create table `blog_auth` (
    `id` int(10) unsigned not null auto_increment,
    `app_key` varchar(20) default '' comment 'Key',
    `app_secret` varchar(50) default '' comment 'Secret',

     #######################################公共字段

    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default  '0' comment '修改时间',
    `modified_by` varchar(100) default  '' comment  '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除0为未删除，1为已删除',
    ######################################

    primary key (`id`) using btree
) engine=InnoDB default charset=utf8mb4 comment='认证管理';