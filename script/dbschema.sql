use payment_bridge;


create table offline_deal_log (
	id              bigint        not null auto_increment,
	offline_deal_id int           not null,
	deal_id         bigint        not null,
	status          varchar(45)   not null,
	message         text,
	create_at       varchar(45)   not null,
    primary key pk_offline_deal_log(id),
    constraint fk_offline_deal_log_offline_deal_id foreign key (offline_deal_id) references offline_deals(id)
);


