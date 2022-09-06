#-- default database name is mcs_v2_1
#-- you can change to another one, in that case please modify database name below and in config file accordingly
#-- create database  mcs_v2_1;
#-- use  mcs_v2_1;

create table network (
    id                             bigint        not null auto_increment,
    name                           varchar(100)  not null,
    last_scan_block_number_payment bigint,
    last_scan_block_number_dao     bigint,
    description                    text,
    create_at                      bigint        not null,
    update_at                      bigint        not null,
    primary key pk_network(id),
    constraint un_network_name unique(name)
);

insert into network(name,create_at,update_at) values('polygon.mumbai',unix_timestamp(),unix_timestamp());
set @network_id_polygon_mumbai:=@@identity;

insert into network(name,create_at,update_at) values('bsc.testnet',unix_timestamp(),unix_timestamp());
set @network_id_bsc_testnet:=@@identity;

create table token (
    id            bigint       not null auto_increment,
    name          varchar(100) not null,
    address       varchar(100) not null,
    network_id    bigint       not null,
    description   text,
    create_at     bigint       not null,
    update_at     bigint       not null,
    primary key pk_token(id),
    constraint un_token_name_network_id unique(name,network_id),
    constraint un_token_address unique(address),
    constraint fk_token_network_id foreign key (network_id) references network(id)
);

insert into token(name,address,network_id,create_at,update_at) values('USDC','0xe11A86849d99F524cAC3E7A0Ec1241828e332C62',@network_id_polygon_mumbai,unix_timestamp(),unix_timestamp());
insert into token(name,address,network_id,create_at,update_at) values('USDC','0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7',@network_id_bsc_testnet,unix_timestamp(),unix_timestamp());

create table wallet (
    id            bigint       not null auto_increment,
    type          int          not null, #--0:metamask, 1:filecoin
    address       varchar(100) not null,
    is_dao        boolean,
    nonce         varchar(64),
    create_at     bigint       not null,
    update_at     bigint       not null,
    primary key pk_wallet(id),
    constraint un_wallet_address_type unique(address,type)
);

create index ind_wallet_is_dao on wallet(is_dao);

create table miner (
    id            bigint       not null auto_increment,
    fid           varchar(100) not null,
    create_at     bigint       not null,
    primary key pk_miner(id)
);

create table source_file (
    id            bigint        not null auto_increment,
    payload_cid   varchar(100)  not null,
    resource_uri  varchar(750)  not null,
    ipfs_url      varchar(1000) not null,
    file_size     bigint        not null,
    dataset       varchar(100),
    pin_status    varchar(100)  not null,
    create_at     bigint        not null,
    update_at     bigint        not null,
    primary key pk_source_file(id),
    constraint un_source_file_payload_cid unique(payload_cid),
    constraint un_source_file_resource_uri unique(resource_uri)
);

create table source_file_upload (
    id             bigint        not null auto_increment,
    source_file_id bigint        not null,
    file_type      int           not null,  #--0:normal file, 1:mint file
    file_name      varchar(200)  not null,
    uuid           varchar(100)  not null,
    wallet_id      bigint        not null,
    status         varchar(100)  not null,
    duration       int           not null,  #--unit:day
    pin_status     varchar(100)  not null,
    is_free        boolean       not null,
    create_at      bigint        not null,
    update_at      bigint        not null,
    primary key pk_source_file_upload(id),
    constraint un_source_file_upload unique(source_file_id,uuid),
    constraint fk_source_file_upload_source_file_id foreign key (source_file_id) references source_file(id),
    constraint fk_source_file_upload_wallet_id foreign key (wallet_id) references wallet(id)
);


create table source_file_mint (
    id                    bigint        not null auto_increment,
    source_file_upload_id bigint        not null,
    nft_tx_hash           varchar(100)  not null,
    mint_address          varchar(100)  not null,
    token_id              bigint        not null,
    create_at             bigint        not null,
    update_at             bigint        not null,
    primary key pk_source_file_mint(id),
    constraint fk_source_file_mint_source_file_upload_id foreign key (source_file_upload_id) references source_file_upload(id)
);

create table car_file (
    id                 bigint        not null auto_increment,
    car_file_name      varchar(200)  not null,
    payload_cid        varchar(200)  not null,
    piece_cid          varchar(200)  not null,
    car_file_size      bigint        not null,
    car_file_path      varchar(1000) not null,
    duration           int           not null,
    task_uuid          varchar(100)  not null,
    max_price          varchar(100)  not null,
    status             varchar(100)  not null,
    is_free            boolean       not null,
    create_at          bigint        not null,
    update_at          bigint        not null,
    primary key pk_car_file(id)
);

create table car_file_source (
    id                    bigint        not null auto_increment,
    car_file_id           bigint        not null,
    source_file_upload_id bigint        not null,
    create_at             bigint        not null,
    primary key pk_car_file_source(id),
    constraint un_car_file_source unique(car_file_id,source_file_upload_id),
    constraint fk_car_file_source_car_file_id foreign key (car_file_id) references car_file(id),
    constraint fk_car_file_source_source_file_upload_id foreign key (source_file_upload_id) references source_file_upload(id)
);

create table offline_deal (
    id               bigint        not null auto_increment,
    car_file_id      bigint        not null,
    deal_cid         varchar(100)  not null,
    miner_id         bigint        not null,
    verified         boolean       not null,
    start_epoch      int           not null,
    sender_wallet_id bigint        not null,
    status           varchar(100)  not null,
    deal_id          bigint,
    on_chain_status  varchar(100),
    unlock_tx_hash   varchar(100),
    unlock_at        bigint,
    note             text,
    create_at        bigint        not null,
    update_at        bigint        not null,
    primary key pk_offline_deal(id),
    constraint un_offline_deal_deal_id unique(deal_id),
    constraint fk_offline_deal_car_file_id foreign key (car_file_id) references car_file(id),
    constraint fk_offline_deal_miner_id foreign key (miner_id) references miner(id),
    constraint fk_offline_deal_sender_wallet_id foreign key (sender_wallet_id) references wallet(id)
);

create table offline_deal_log (
    id               bigint        not null auto_increment,
    offline_deal_id  bigint        not null,
    on_chain_status  varchar(100)  not null,
    on_chain_message text,
    create_at        bigint        not null,
    primary key pk_offline_deal_log(id),
    constraint fk_offline_deal_log_offline_deal_id foreign key (offline_deal_id) references offline_deal(id)
);

create table transaction (
    id                           bigint        not null auto_increment,
    source_file_upload_id        bigint        not null,
    network_id                   bigint        not null,
    token_id                     bigint        not null,
    wallet_id_pay                bigint        not null,
    wallet_id_recipient          bigint        not null,
    wallet_id_contract           bigint        not null,
    pay_tx_hash                  varchar(100)  not null,
    pay_amount                   varchar(100)  not null,
    pay_at                       bigint        not null,
    deadline                     bigint        not null,
    unlock_amount                varchar(100),
    last_unlock_at               bigint,
    refund_tx_hash               varchar(100),
    refund_amount                varchar(100),
    refund_at                    bigint,
    refund_by_wallet_id          bigint,
    create_at                    bigint        not null,
    update_at                    bigint        not null,
    primary key pk_transaction(id),
    constraint un_transaction unique(source_file_upload_id),
    constraint fk_transaction_source_file_upload_id foreign key (source_file_upload_id) references source_file_upload(id),
    constraint fk_transaction_network_id foreign key (network_id) references network(id),
    constraint fk_transaction_token_id foreign key (token_id) references token(id),
    constraint fk_transaction_wallet_id_pay foreign key (wallet_id_pay) references wallet(id),
    constraint fk_transaction_wallet_id_recipient foreign key (wallet_id_recipient) references wallet(id),
    constraint fk_transaction_wallet_id_contract foreign key (wallet_id_contract) references wallet(id),
    constraint fk_transaction_refund_by_wallet_id foreign key (refund_by_wallet_id) references wallet(id)
);

create table dao_pre_sign (
    id                           bigint        not null auto_increment,
    offline_deal_id              bigint        not null,
    batch_count                  int           not null,
    batch_size_max               int           not null,
    source_file_upload_cnt_total int           not null,
    source_file_upload_cnt_sign  int           not null,
    network_id                   bigint        not null,
    wallet_id_signer             bigint        not null,
    wallet_id_recipient          bigint        not null,
    wallet_id_contract           bigint        not null,
    tx_hash                      varchar(100)  not null,
    status                       varchar(100)  not null,
    create_at                    bigint        not null,
    update_at                    bigint        not null,
    primary key pk_dao_pre_sign(id),
    constraint un_dao_pre_sign unique(offline_deal_id,wallet_id_signer),
    constraint fk_dao_pre_sign_offline_deal_id foreign key (offline_deal_id) references offline_deal(id)
);

create table dao_signature (
    id                           bigint        not null auto_increment,
    offline_deal_id              bigint        not null,
    batch_no                     int,
    network_id                   bigint        not null,
    wallet_id_signer             bigint        not null,
    wallet_id_recipient          bigint,
    wallet_id_contract           bigint        not null,
    tx_hash                      varchar(100)  not null,
    status                       varchar(100)  not null,
    signed_by_hash               boolean       not null,
    create_at                    bigint        not null,
    update_at                    bigint        not null,
    primary key pk_dao_signature(id),
    constraint un_dao_signature unique(offline_deal_id,tx_hash),
    constraint fk_dao_signature_network_id foreign key (network_id) references network(id),
    constraint fk_dao_signature_offline_deal_id foreign key (offline_deal_id) references offline_deal(id),
    constraint fk_dao_signature_wallet_id_signer foreign key (wallet_id_signer) references wallet(id),
    constraint fk_dao_signature_wallet_id_recipient foreign key (wallet_id_recipient) references wallet(id),
    constraint fk_dao_signature_wallet_id_contract foreign key (wallet_id_contract) references wallet(id)
);

create table dao_signature_source_file_upload (
    id                           bigint        not null auto_increment,
    dao_signature_id             bigint        not null,
    source_file_upload_id        bigint        not null,
    create_at                    bigint        not null,
    update_at                    bigint        not null,
    primary key pk_dao_signature_source_file_upload(id),
    constraint fk_dao_signature_source_file_upload_dao_signature_id foreign key (dao_signature_id) references dao_signature(id),
    constraint fk_dao_signature_source_file_upload_source_file_upload_id foreign key (source_file_upload_id) references source_file_upload(id)
);



#--2022.09.06
alter table network add last_scan_block_number_dao bigint;
alter table network change last_scan_block_number last_scan_block_number_payment bigint;
alter table car_file drop column deal_success;

alter table transaction change refund_after_unlock_amount refund_amount varchar(100);
alter table transaction change refund_after_unlock_at refund_at bigint;
alter table transaction change refund_after_unlock_tx_hash refund_tx_hash varchar(100);


SET SQL_SAFE_UPDATES = 0;
update transaction set refund_amount=refund_after_expired_amount where refund_amount is null and refund_after_expired_amount is not null;
update transaction set refund_at=refund_after_expired_at where refund_at is null and refund_after_expired_at is not null;
update transaction set refund_tx_hash=refund_after_expired_tx_hash is null and refund_after_expired_tx_hash is not null;
alter table transaction drop column refund_after_expired_tx_hash;
alter table transaction drop column refund_after_expired_amount;
alter table transaction drop column refund_after_expired_at;
SET SQL_SAFE_UPDATES = 1;

alter table offline_deal add note             text;


alter table source_file_mint modify token_id bigint  not null;

alter table dao_signature add signed_by_hash               boolean       not null;

alter table source_file_upload add pin_status    varchar(100)  not null;
SET SQL_SAFE_UPDATES = 0;
update source_file_upload a set pin_status=(select pin_status from source_file b where a.source_file_id=b.id);
SET SQL_SAFE_UPDATES = 1;
alter table dao_signature add batch_no int;
alter table dao_signature modify wallet_id_recipient          bigint;

alter table source_file_upload add is_free        boolean       not null;
alter table car_file add is_free        boolean       not null;

alter table transaction add refund_by_wallet_id          bigint;
alter table transaction add constraint fk_transaction_refund_by_wallet_id foreign key (refund_by_wallet_id) references wallet(id);
alter table token drop key un_token_name;
alter table token add  constraint un_token_name_network_id unique(name,network_id);
insert into network(name,create_at,update_at) values('bsc.testnet',unix_timestamp(),unix_timestamp());
set @network_id_bsc:=@@identity;
insert into token(name,address,network_id,create_at,update_at) values('USDC','0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7',@network_id_bsc,unix_timestamp(),unix_timestamp());


create table dao_pre_sign (
    id                           bigint        not null auto_increment,
    offline_deal_id              bigint        not null,
    batch_count                  int           not null,
    batch_size_max               int           not null,
    source_file_upload_cnt_total int           not null,
    source_file_upload_cnt_sign  int           not null,
    network_id                   bigint        not null,
    wallet_id_signer             bigint        not null,
    wallet_id_recipient          bigint        not null,
    wallet_id_contract           bigint        not null,
    tx_hash                      varchar(100)  not null,
    status                       varchar(100)  not null,
    create_at                    bigint        not null,
    update_at                    bigint        not null,
    primary key pk_dao_pre_sign(id),
    constraint un_dao_pre_sign unique(offline_deal_id,wallet_id_signer),
    constraint fk_dao_pre_sign_offline_deal_id foreign key (offline_deal_id) references offline_deal(id)
);

