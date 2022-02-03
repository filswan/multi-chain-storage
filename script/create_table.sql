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

