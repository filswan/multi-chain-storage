alter table source_file add payload_cid varchar(100) not null default '';
alter table source_file add vrf_rand    varchar(100) not null default '';

alter table event_lock_payment add vrf_rand    varchar(100) not null default '';

