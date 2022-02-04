alter table source_file add payload_cid varchar(100) not null default '';
alter table source_file modify file_size bigint;
alter table source_file modify create_at bigint;
alter table source_file add update_at bigint;

alter table event_lock_payment add vrf_rand    varchar(100) not null default '';


alter table block_scan_record modify  update_at bigint;
alter table dao_fetched_deal modify create_at bigint not null;
alter table dao_info modify create_at bigint;


alter table deal_file modify create_at bigint;

alter table deal_file modify lock_payment_status status;


SET SQL_SAFE_UPDATES = 0;
update deal_file set update_at=null where update_at='';
alter table deal_file modify update_at bigint;
#--alter table deal_file add refund_status_after_unlock varchar(45);

update deal_file set duration=duration/2880;


SET SQL_SAFE_UPDATES = 0;
update deal_file set delete_at=null where delete_at='';
alter table deal_file modify delete_at bigint;

alter table deal_file add max_price   double;


alter table event_lock_payment modify create_at bigint;

alter table event_unlock_payment modify create_at bigint;
alter table event_unlock_payment add deal_id bigint;
alter table event_unlock_payment add source_file_id bigint;
alter table event_unlock_payment add foreign key fk_event_unlock_payment_source_file_id(source_file_id) references source_file(id)



alter table event_lock_payment add source_file_id bigint;
alter table event_lock_payment add foreign key fk_event_lock_payment_source_file_id(source_file_id) references source_file(id)

alter table event_lock_payment modify locked_fee double;




SET SQL_SAFE_UPDATES = 0;
update event_lock_payment set lock_payment_time='0' where lock_payment_time='';
alter table event_lock_payment modify lock_payment_time bigint;


create unique index un_event_lock_payment on event_lock_payment(payload_cid, address_from);


alter table source_file_deal_file_map modify create_at bigint;

alter table source_file_deal_file_map modify update_at bigint;


create table offline_deal (
    id            bigint       not null auto_increment,
    deal_file_id  bigint       not null,
    deal_cid      varchar(100) not null,
    miner_fid     varchar(45)  not null,
    start_epoch   int          not null,
    sender_wallet varchar(200) not null,
    status        varchar(45)  not null,
	deal_id       bigint       not null,
    unlock_status varchar(45)  not null,
    note          text,
	create_at     bigint       not null,
	update_at     bigint       not null,
    unlock_at     bigint,
    primary key pk_ofline_deal(id),
    constraint fk_ofline_deal_deal_file_id foreign key (deal_file_id) references deal_file (id)
);


create index ind_offline_deal_deal_file_id on offline_deal(deal_file_id)


alter table source_file add file_type int;


create table source_file_upload_history (
    id             bigint       not null auto_increment,
    source_file_id bigint       not null,
    file_name      varchar(200) not null,
    wallet_address varchar(200) not null,
    status         varchar(45)  not null,
	create_at      bigint       not null,
	update_at      bigint       not null,
    primary key pk_source_file_upload_history(id),
    constraint fk_source_file_upload_history_source_file_id foreign key (source_file_id) references source_file (id)
)


alter table source_file drop column file_name;

alter table source_file drop column wallet_address;

