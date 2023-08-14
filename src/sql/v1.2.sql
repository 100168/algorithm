#-------------------------------采购订单模块-----------------------------
create table purchase_order
(
    id                     bigint unsigned auto_increment
        primary key,
    biz_type               varchar(16) null comment '业务类型',
    trade_no               varchar(32)                        not null comment '订单交易流水号',
    purchase_order_no      varchar(32)                        not null comment '采购订单号',
    sap_order_no           varchar(32) null comment 'SAP单号',
    purchase_order_attr    varchar(32) null comment '采购订单归属',
    purchase_order_type    varchar(32) null comment '采购订单类型',
    store_code             varchar(50) null comment '门店编码',
    commit_time            datetime default CURRENT_TIMESTAMP not null comment '订单提交时间',
    receive_warehouse_code varchar(200) null comment '收货仓库编码',
    purchase_order_amount  decimal(20, 2) null comment '订单总金额',
    purchase_order_status  varchar(32) null comment '订单状态',
    receiver               varchar(100) null comment '收件人',
    receiver_phone         varchar(32) null comment '收件人电话',
    receiver_province      varchar(50) null comment '收货省',
    receiver_city          varchar(50) null comment '收货城市',
    receiver_district      varchar(50) null comment '收货区',
    receiver_address       varchar(500) null comment '收货详细地址',
    vin                    varchar(32) null comment 'vin',
    description            varchar(1024) null comment '描述',
    remark                 varchar(1024) null comment '备注',
    sap_order_create_time      datetime  not null comment 'sap创建时间',
    is_del                 bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user            char(32) null comment '创建者',
    create_time            datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user            char(32) null comment '更新者',
    update_time            datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '采购订单';
create table purchase_order_detail
(
    id                  bigint unsigned auto_increment
        primary key,
    biz_type            varchar(16) null comment '业务类型',
    purchase_order_no   varchar(32)                        not null comment '采购订单号',
    sale_part_num       varchar(64) null comment '售后件号',
    discount_unit_price decimal(20, 2) null comment '折后单价',
    qty                 decimal(20, 3) null comment '数量',
    line_amount         decimal(20, 2) null comment '行金额',
    description         varchar(1024) null comment '描述',
    remark              varchar(1024) null comment '备注',
    is_del              bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user         char(32) null comment '创建者',
    create_time         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user         char(32) null comment '更新者',
    update_time         datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '采购订单明细';

create table purchase_order_operate_log
(
    id                bigint unsigned auto_increment
        primary key,
    biz_type          varchar(16) null comment '业务类型',

    purchase_order_no varchar(32)                        not null comment '采购订单号',
    action_desc       varchar(32) null comment '操作描述',
    modify_column     varchar(32) null comment '修改字段',
    old_value         varchar(32) null comment '原始值',
    new_value         varchar(32) null comment '修改值',
    description       varchar(1024) null comment '描述',
    remark            varchar(1024) null comment '备注',
    is_del            bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user       char(32) null comment '创建者',
    create_time       datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user       char(32) null comment '更新者',
    update_time       datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '采购订单操作记录';


#--------------------------销售订单模块-----------------------------
create table sale_order
(
    id                     bigint unsigned auto_increment
        primary key,
    biz_type               varchar(16) null comment '业务类型',
    sale_order_no          varchar(32)                        not null comment '销售单号',
    sap_order_no           varchar(32) null comment 'SAP单号',
    is_dfs                 bit      default b'0' null comment '是否dfs',
    purchase_order_no      varchar(32) null comment '采购订单号',
    deliver_warehouse_code varchar(200) null comment '发货仓库编码',
    dfs_status             varchar(32) null comment 'DFS状态',
    sale_order_status      varchar(32) null comment '销售订单状态',
    sale_order_type        varchar(32) null comment '销售订单类型',
    transfer_advice        varchar(32) null comment '运输建议',
    sync_status            bigint   default 0 comment '同步状态,0成功,非0需要同步',
    store_code             varchar(50) null comment '门店编码',
    receive_warehouse_code varchar(200) null comment '收货仓库编码',
    receiver_province      varchar(50) null comment '收货省',
    receiver_city          varchar(50) null comment '收货城市',
    receiver_district      varchar(50) null comment '收货区',
    receiver               varchar(100) null comment '收件人',
    receiver_phone         varchar(32) null comment '收件人电话',
    receiver_address       varchar(500) null comment '收货地址',

    description            varchar(1024) null comment '描述',
    remark                 varchar(1024) null comment '备注',
    is_del                 bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user            char(32) null comment '创建者',
    create_time            datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user            char(32) null comment '更新者',
    update_time            datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '销售订单';

create table sale_order_detail
(
    id            bigint unsigned auto_increment
        primary key,
    biz_type      varchar(16) null comment '业务类型',
    sale_order_no varchar(32)                        not null comment '销售单号',
    sale_part_num varchar(64) null comment '售后件号',
    deliver_qty   decimal(20, 3) null comment '发货数量',
    qty           decimal(20, 3) null comment '数量',
    receive_qty   decimal(20, 3) null comment '收货数量',
    description   varchar(1024) null comment '描述',
    remark        varchar(1024) null comment '备注',
    is_del        bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user   char(32) null comment '创建者',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user   char(32) null comment '更新者',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '采购订单明细';

create table sale_order_operate_log
(
    id            bigint unsigned auto_increment
        primary key,
    biz_type      varchar(16) null comment '业务类型',

    sale_order_no varchar(32)                        not null comment '销售单号',
    action_desc   varchar(32) null comment '操作描述',
    modify_column varchar(32) null comment '修改字段',
    old_value     varchar(32) null comment '原始值',
    new_value     varchar(32) null comment '修改值',
    description   varchar(1024) null comment '描述',
    remark        varchar(1024) null comment '备注',
    is_del        bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user   char(32) null comment '创建者',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user   char(32) null comment '更新者',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '销售订单操作记录';


#----------------------------缺件管理模块----------------------------
create table back_order
(
    id                bigint unsigned auto_increment
        primary key,
    biz_type          varchar(16) null comment '业务类型',
    back_order_no     varchar(32)                        not null comment '缺件单号',
    back_order_type   varchar(32) null comment '缺件订单类型',
    purchase_order_no varchar(32)                        not null comment '采购订单号',
    sale_part_num     varchar(64) null comment '售后件号',
    back_qty          decimal(20, 3) null comment '缺件总数',
    rest_back_qty     decimal(20, 3) null comment '剩余缺件总数',
    back_order_status varchar(32) null comment '缺件订单状态',
    store_code        varchar(50) null comment '门店编码',
    description       varchar(1024) null comment '描述',
    remark            varchar(1024) null comment '备注',
    is_del            bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user       char(32) null comment '创建者',
    create_time       datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user       char(32) null comment '更新者',
    update_time       datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '缺件订单';

create table back_order_operate_log
(
    id            bigint unsigned auto_increment
        primary key,
    biz_type      varchar(16) null comment '业务类型',

    back_order_no varchar(32)                        not null comment '缺件单号',
    action_desc   varchar(32) null comment '操作描述',
    modify_column varchar(32) null comment '修改字段',
    old_value     varchar(32) null comment '原始值',
    new_value     varchar(32) null comment '修改值',
    description   varchar(1024) null comment '描述',
    remark        varchar(1024) null comment '备注',
    is_del        bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user   char(32) null comment '创建者',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user   char(32) null comment '更新者',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '缺件订单操作记录';


create table back_sale_relation
(
    id            bigint unsigned auto_increment
        primary key,
    biz_type      varchar(16) null comment '业务类型',
    back_order_no varchar(32)                        not null comment '缺件单号',
    sale_order_no varchar(32)                        not null comment '销售单号',
    description   varchar(1024) null comment '描述',
    remark        varchar(1024) null comment '备注',
    is_del        bit      default b'0'              not null comment '是否删除 1是 0否 默认0',
    create_user   char(32) null comment '创建者',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_user   char(32) null comment '更新者',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '缺件订单跟销售订单的关系';
#---------索引--------------------------------------------------------------------------------------------

