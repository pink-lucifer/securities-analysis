create database fin_security;

crearte user security@'%' identified by 'security8888';
grant all privileges on fin_security.* to security;
flush all privileges;

create table listed_symbol(
  `id` bigint primary key auto_increment comment 'auto incremental id',
  `uuid` varchar(64) not null comment 'uuid',
  `listed_market` INT(4) not null comment 'listed market',
  `symbol` varchar(32) not null comment 'listed symbol',
  `name` varchar(256) not null comment 'name of security',
  `ipo_year` year not null default '0000' comment 'the year of ipo',
  `sector` varchar(256) null comment 'sector of security',
  `industry` varchar(256) null comment 'industry of security',
  `summary_quote_url` varchar(256) null comment 'url for summary quote',
  `created_timestamp` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_updated_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
create unique index listed_symbol_uuid on listed_symbol(uuid) ;
create unique index listed_symbol_mkt_symbol on listed_symbol(listed_market, symbol);

insert into listed_symbol(uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url)
 values ('2b514f90-9fb4-41e5-8c0b-e56718a6cd05',9999,''TEST','TEST NAME',1985,'TEST SECTOR','TEST INDUSTRY','TEST SUMMARY QUOTE URL');


create table eod_price(
  `id` bigint primary key auto_increment comment 'auto incremental id',
  `uuid` varchar(64) not null comment 'uuid',
  `listed_market` INT(4) not null comment 'listed market',
  `symbol` varchar(32) not null comment 'listed symbol',
  `trade_date` DATE not null comment 'trade date',
  `open_price` decimal(18,6) not null default 0 comment 'open price for the listed symbol',
  `close_price` decimal(18,6) not null default 0 comment 'close price for the listed symbol',
  `highest_price` decimal(18,6) not null default 0 comment 'highest price for the listed symbol',
  `lowest_price` decimal(18,6) not null default 0 comment 'lowest price for the listed symbol',
  `volume` decimal(30,6) null comment 'url for summary quote',
  `created_timestamp` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_updated_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

create unique index eod_price_uuid on eod_price(uuid) ;
create unique index eod_price_date_mkt_symbol on eod_price(trade_date, listed_market, symbol);

insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume)
VALUES ('345e5f3e-e87c-429c-975f-6da7d6033f6f',9999,'TEST','2018-01-01',10.00, 10.89, 12.25, 9.89, 10000000);
insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume)
VALUES ('f1d37e74-0e52-4c20-86e8-26cbfe6c8b01',9999,'TEST','2018-01-02',11.00, 11.89, 13.25, 10.89, 11000000);
insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume)
VALUES ('fe66aebc-0f46-46da-8a03-1ca4985c77cc',9999,'TEST','2018-01-03',12.00, 12.89, 14.25, 11.89, 12000000);
insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume)
VALUES ('d48eec2b-2e44-42ad-a9d4-6f450715c356',9999,'TEST','2018-01-04',13.00, 13.89, 15.25, 12.89, 13000000);
insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume)
VALUES ('20d48fff-024f-4c61-8f9b-584de4e67012',9999,'TEST','2018-01-05',14.00, 14.89, 16.25, 13.89, 14000000);
commit;