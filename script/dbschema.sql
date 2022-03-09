use payment_bridge;


create table offline_deal_log (
	id              bigint        not null auto_increment,
	deal_cid        varchar(100)  not null,
	status          varchar(45)   not null,
	message         text,
	create_at       varchar(45)   not null,
    primary key pk_offline_deal_log(id)
);

create index idx_offline_deal_log on offline_deal_log(deal_cid);
