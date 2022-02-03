create database mcp_v2_1;
use mcp_v2_1;

drop table network;

create table network (
    id             bigint       not null auto_increment,
    name           varchar(60)  not null,
    rpc_url        varchar(200) not null,
    native_coin    varchar(60)  not null,
    description    varchar(1000),
	create_at      bigint       not null,
	update_at      bigint       not null,
    primary key pk_network(id),
    unique key un_network_name(name)
);

set @curUtcMilliSec=UNIX_TIMESTAMP()*1000;

insert into network(name,rpc_url,native_coin,create_at,update_at) values('polygon','https://polygon-mumbai.g.alchemy.com/v2/86HeefA3O9EF22t2NTLbmcpfN0hb9vlv','MATIC',@curUtcMilliSec,@curUtcMilliSec);
insert into network(name,rpc_url,native_coin,create_at,update_at) values('goerli','https://goerli.infura.io/v3/a30f13ea65fe406a86783fa912982906','GOERLI',@curUtcMilliSec,@curUtcMilliSec);
insert into network(name,rpc_url,native_coin,create_at,update_at) values('nbai','https://api.nbai.io/','NBAI',@curUtcMilliSec,@curUtcMilliSec);
insert into network(name,rpc_url,native_coin,create_at,update_at) values('bsc','https://data-seed-prebsc-1-s1.binance.org:8545/','BNB',@curUtcMilliSec,@curUtcMilliSec);

create table coin (
    id           bigint        not null auto_increment,
    name         varchar(60)   not null,
    address      varchar(200)  not null,
    network_id   bigint        not null,
    gas_price    decimal(20,0) not null,
    gas_limit    decimal(20,0) not null,
    description  varchar(1000),
    create_at    bigint        not null,
    update_at    bigint        not null,
    primary key pk_coin(id),
    unique key un_coin_name(name),
    unique key un_coin_address(address),
    constraint fk_coin_network_id foreign key (network_id) references network (id)
)

insert into coin(name,address,network_id,gas_price,gas_limit,create_at,update_at) values('USDC','0xe11A86849d99F524cAC3E7A0Ec1241828e332C62',1,0,0,@curUtcMilliSec,@curUtcMilliSec);
insert into coin(name,address,network_id,gas_price,gas_limit,create_at,update_at) values('WFIL','0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055',1,0,0,@curUtcMilliSec,@curUtcMilliSec);


create table dao_signer (
	id           bigint        not null auto_increment,
	name         varchar(60)   not null,
	address      varchar(200)  not null,
	order        int           not null,
    description  varchar(1000),
    create_at    bigint        not null,
    update_at    bigint        not null,
    primary key pk_dao_signer(id),
    unique key un_dao_signer_name(name),
    unique key un_dao_signer_address(address)
);

insert into dao_signer (name,address,order,create_at,update_at) values('Dao1','0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d',1,@curUtcMilliSec,@curUtcMilliSec);
insert into dao_signer (name,address,order,create_at,update_at) values('Dao2','0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1',2,@curUtcMilliSec,@curUtcMilliSec);
insert into dao_signer (name,address,order,create_at,update_at) values('Dao3','0x800210CfB747992790245eA878D32F188d01a03A',3,@curUtcMilliSec,@curUtcMilliSec);



create table dao_fetched_deal (
	id        bigint not null auto_increment,
	deal_id   bigint not null,
	create_at bigint not null,
    primary key pk_dao_fetched_deal(id)
);

create table car_file (
	id                    bigint not null auto_increment,
	name                  varchar(60)   not null,
	payload_cid           varchar(60)   not null,
	piece_cid             varchar(60)   not null,
	file_size             bigint        not null,
	pin_status            varchar(60)   not null,
	file_path             varchar(200)  not null,
	car_md_5              varchar(60)   not null,
	duration              bigint        not null,
	task_uuid             varchar(60)   not null,
	status                varchar(60)   not null,
	client_wallet_address varchar(60)   not null,
	max_price             decimal(20,0) not null,
	create_at             bigint        not null,
	update_at             bigint        not null,
    primary key pk_car_file(id),
    unique key un_car_file_payload_cid(payload_cid)
);


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
    primary key pk_ofline_deal(id),
    constraint fk_ofline_deal_deal_file_id foreign key (deal_file_id) references deal_file (id)
);

create table source_file (
	id           bigint        not null auto_increment,
	resource_uri varchar(1000) not null,
	status       varchar(60)   not null,
	file_size    bigint        not null,
	dataset      varchar(200)  not null,
	ipfs_url     varchar(1000) not null,
	pin_status   varchar(60)   not null,
	payload_cid  varchar(200)  not null,
	nft_tx_hash  varchar(200)  not null,
	token_id     varchar(60)   not null,
	mint_address varchar(200)  not null,
	file_type    int           not null,
    car_file_id  bigint,
	create_at    bigint        not null,
	update_at    bigint        not null,
    primary key pk_source_file_id(id),
    unique key un_source_file_payload_cid(payload_cid),
    constraint fk_source_file_car_file_id foreign key (car_file_id) references car_file (id)
);

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
);


create table event_dao_signature (
	id                      bigint       not null auto_increment,
	tx_hash                 varchar(200) not null,
	recipient               varchar(200) not null,
	payload_cid             varchar(200) not null,
	deal_id                 bigint       not null,
	dao_pass_time           varchar(200) not null,
	block_no                bigint       not null,
	block_time              varchar(200) not null,
	status                  boolean      not null,
	network_id              bigint       not null,
	coin_id                 bigint       not null,
	dao_address             varchar(200) not null,
	signature_unlock_status varchar(200) not null,
	tx_hash_unlock          varchar(200) not null,
	create_at               bigint       not null,
	update_at               bigint       not null,
    primary key pk_event_dao_signature(id),
);


