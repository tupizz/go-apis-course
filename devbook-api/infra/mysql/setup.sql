CREATE OR REPLACE TABLE usuarios
(
    id            int auto_increment primary key,
    name          varchar(100) not null,
    nick          varchar(100) not null unique,
    email         varchar(100) not null unique,
    password      varchar(20) not null,
    created_at    timestamp default current_timestamp()
) ENGINE=INNODB;