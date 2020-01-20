
mysql> create database test;
Query OK, 1 row affected (0.05 sec)

mysql> use test;


create table users(
  id int auto_increment primary key,
  name varchar(100) not null,
  email varchar(100) not null unique,
  password varchar(100) not null,
  createAt timestamp default current_timestamp()
)engine=InnoDB default charset=utf8;
