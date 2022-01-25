alter table source_file add payload_cid varchar(100) not null default '';
alter table source_file modify file_size bigint;
alter table source_file modify create_at bigint;

alter table event_lock_payment add vrf_rand    varchar(100) not null default '';


alter table block_scan_record modify  update_at bigint;
alter table dao_fetched_deal modify create_at bigint not null;
alter table dao_info modify create_at bigint;


alter table deal_file modify create_at bigint;
SET SQL_SAFE_UPDATES = 0;
update deal_file set update_at=null where update_at='';
alter table deal_file modify update_at bigint;


SET SQL_SAFE_UPDATES = 0;
update deal_file set delete_at=null where delete_at='';
alter table deal_file modify delete_at bigint;

alter table deal_file add max_price   double;

alter table event_lock_payment modify create_at bigint;

alter table event_unlock_payment modify create_at bigint;


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
	create_at     bigint       not null,
	update_at     bigint       not null,
    primary key pk_ofline_deal(id),
    constraint fk_ofline_deal_deal_file_id foreign key (deal_file_id) references deal_file (id)
);





