alter table source_file add payload_cid varchar(100) not null default '';

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


alter table source_file modify create_at bigint;

alter table source_file_deal_file_map modify create_at bigint;

alter table source_file_deal_file_map modify update_at bigint;



