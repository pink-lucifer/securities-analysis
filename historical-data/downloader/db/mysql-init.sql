create database fin_security;

crearte user security@'%' identified by 'security8888';
grant all privileges on fin_security.* to security;
flush all privileges;

create table listed_symbol(
  `id` bigint primary key auto_increment comment 'auto incremental id',
  `uuid` varchar(64) not null comment 'uuid',
  `listed_market` INT(4) not null comment 'listed market',
  `symbol` varchar(6) not null comment 'listed symbol',
  `name` varchar(128) not null comment 'name of security',
  `ipo_year` year not null default '0000' comment 'the year of ipo',
  `sector` varchar(32) null comment 'sector of security',
  `industry` varchar(32) null comment 'industry of security',
  `summary_quote_url` varchar(256) null comment 'url for summary quote',
  `created_timestamp` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_updated_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
create unique index listed_symbol_uuid on listed_symbol(uuid) ;
create unique index listed_symbol_mkt_symbol on listed_symbol(listed_market, symbol);

insert into listed_symbol(uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url)
 values ('2b514f90-9fb4-41e5-8c0b-e56718a6cd05',9999,'TEST','TEST NAME',1985,'TEST SECTOR','TEST INDUSTRY','TEST SUMMARY QUOTE URL');
commit;