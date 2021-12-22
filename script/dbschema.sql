create schema payment_bridge collate utf8_bin;
use payment_bridge;
create table if not exists dao_fetched_deal
(
	id        bigint auto_increment primary key,
	deal_id   bigint       not null,
	create_at varchar(128) not null
);

create table if not exists dao_info
(
	id          bigint auto_increment primary key,
	dao_name    varchar(255) not null,
	dao_address varchar(255) null,
	order_index int          not null,
	description varchar(255) null,
	create_at   varchar(64)  null
);

create table if not exists deal_file
(
	id                    bigint auto_increment primary key,
	deal_cid              varchar(255) default '' null,
	car_file_name         varchar(255) charset utf8 null,
	payload_cid           varchar(255) charset utf8 null,
	deal_id               bigint null,
	piece_cid             varchar(255) charset utf8 null,
	car_file_size         bigint null,
	miner_fid             varchar(128) null,
	source_file_path      varchar(255) charset utf8 null,
	car_file_path         varchar(255) charset utf8 null,
	car_md5               varchar(45) charset utf8 null,
	deal_status           varchar(64) charset utf8 null,
	pin_status            varchar(64) null,
	duration              int null,
	task_uuid             varchar(128) default '' null,
	cost                  varchar(255) null,
	client_wallet_address varchar(255) null,
	lock_payment_tx       varchar(255) null,
	lock_payment_status   varchar(32) default 'Pending' null,
	lock_payment_network  bigint null,
	dao_sign_status       varchar(32) null,
	send_deal_status      varchar(32) default '' null,
	verified              tinyint(1) default 0 null,
	create_at             varchar(128) null,
	delete_at             varchar(128) null,
	update_at             varchar(128) null,
	is_deleted            tinyint(1) default 0 null
);

create index deal_file_network_id_fk on deal_file (lock_payment_network);

create table if not exists network
(
	id           bigint auto_increment primary key,
	network_name varchar(255) not null,
	uuid         varchar(128) null,
	rpc_url      varchar(255) null,
	native_coin  varchar(128) null,
	description  varchar(255) null,
	constraint network_info_uuid_uindex	unique (uuid)
);

create table if not exists block_scan_record
(
	id         bigint auto_increment primary key,
	network_id bigint null,
	last_current_block_number bigint not null,
	update_at varchar(64) null,
	constraint number_UNIQUE
		unique (last_current_block_number),
	constraint block_scan_record_network_id_fk
		foreign key (network_id) references network (id)
);

create table if not exists coin
(
	id bigint auto_increment
		primary key,
	short_name varchar(255) not null,
	full_name varchar(255) null,
	cn_name varchar(255) null,
	coin_address varchar(255) null,
	uuid varchar(128) null,
	network_id bigint null,
	gas_price int default 0 null,
	gas_limit int default 0 null,
	description varchar(255) null,
	constraint coin_uuid_uindex
		unique (uuid),
	constraint coin_network_id_fk
		foreign key (network_id) references network (id)
);

create table if not exists event_dao_signature
(
	id bigint auto_increment
		primary key,
	tx_hash varchar(255) null,
	recipient varchar(255) null,
	payload_cid varchar(255) null,
	order_id varchar(255) null,
	deal_id bigint null,
	dao_pass_time varchar(64) null,
	block_no bigint null,
	coin_id bigint null,
	network_id bigint null,
	dao_address varchar(255) null,
	block_time varchar(64) null,
	signature_unlock_status varchar(8) default '0' null,
	status tinyint(1) null,
	tx_hash_unlock varchar(255) null,
	constraint dao_event_log_coin_id_fk
		foreign key (coin_id) references coin (id),
	constraint dao_event_log_network_id_fk
		foreign key (network_id) references network (id)
);

create index dao_event_log_order_id_index
	on event_dao_signature (order_id);

create index dao_event_log_payload_cid_index
	on event_dao_signature (payload_cid);

create index dao_event_log_tx_hash_index
	on event_dao_signature (tx_hash);

create table if not exists event_lock_payment
(
	id bigint auto_increment
		primary key,
	network_id bigint null,
	tx_hash varchar(255) not null,
	event_name text null,
	payload_cid varchar(256) null,
	token_address varchar(255) null,
	contract_address varchar(255) null,
	locked_fee varchar(255) null,
	min_payment varchar(255) null,
	deadline varchar(64) null,
	block_no bigint null,
	miner_address varchar(255) null,
	address_from varchar(255) null,
	address_to varchar(255) null,
	create_at varchar(64) null,
	lock_payment_time varchar(64) null,
	unlock_tx_hash varchar(255) null,
	unlock_tx_status varchar(32) null,
	unlock_time varchar(64) null,
	coin_id bigint null,
	constraint event_lock_payment_coin_info_id_fk
		foreign key (coin_id) references coin (id),
	constraint event_polygon_network_info_id_fk
		foreign key (network_id) references network (id)
);

create table if not exists event_unlock_payment
(
	id bigint auto_increment
		primary key,
	tx_hash varchar(255) null,
	payload_cid varchar(255) null,
	block_no varchar(128) null,
	token_address varchar(255) null,
	unlock_from_address varchar(255) null,
	unlock_to_user_address varchar(255) null,
	unlock_to_user_amount varchar(255) null,
	unlock_to_admin_address varchar(255) null,
	unlock_to_admin_amount varchar(255) null,
	unlock_time varchar(64) null,
	create_at varchar(64) null,
	network_id bigint null,
	coin_id bigint null,
	unlock_status varchar(32) null,
	constraint event_unlock_payment_coin_info_id_fk
		foreign key (coin_id) references coin (id),
	constraint event_unlock_payment_network_info_id_fk
		foreign key (network_id) references network (id)
);

create table if not exists source_file
(
	id bigint auto_increment
		primary key,
	resource_uri varchar(255) charset utf8 null,
	md5 varchar(45) charset utf8 null,
	uuid varchar(45) charset utf8 null,
	status varchar(45) charset utf8 null,
	dataset varchar(45) null,
	create_at varchar(128) null,
	file_size bigint null,
	file_name varchar(255) charset utf8 not null,
	delete_at varchar(128) null,
	user_id int null,
	ipfs_url varchar(255) null,
	pin_status varchar(32) null
);

create table if not exists source_file_deal_file_map
(
	source_file_id bigint not null,
	file_index int not null,
	deal_file_id bigint not null,
	create_at varchar(128) null,
	update_at varchar(128) null,
	primary key (source_file_id, deal_file_id),
	constraint fk_source_file_has_deal_file_deal_file1
		foreign key (deal_file_id) references deal_file (id),
	constraint fk_source_file_has_deal_file_source_file1
		foreign key (source_file_id) references source_file (id)
);

create index fk_source_file_has_deal_file_deal_file1_idx
	on source_file_deal_file_map (deal_file_id);

create index fk_source_file_has_deal_file_source_file1_idx
	on source_file_deal_file_map (source_file_id);

create table if not exists system_config_param
(
	id bigint auto_increment
		primary key,
	param_key varchar(255) null,
	param_value varchar(255) null,
	module varchar(128) null,
	`desc` varchar(255) null,
	constraint table_name_param_key_uindex
		unique (param_key)
);

