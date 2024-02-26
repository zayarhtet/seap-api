create database seap_db;

use seap_db;

create table credential (
	credential_id int not null auto_increment primary key,
    password varchar(255) not null
);
ALTER TABLE credential
ADD UNIQUE (credential_id);

show tables;

create table role (
	role_id int primary key not null,
    name varchar(20) not null
);

ALTER TABLE role
ADD UNIQUE (role_id);

insert into role (`role_id`, `name`) values ("0", "admin");
insert into role (`role_id`, `name`) values ("1", "tutor");
insert into role (`role_id`, `name`) values ("2", "tutee");

create table member (
	member_id int not null primary key auto_increment unique,
    first_name varchar(255),
    last_name varchar(255),
    username varchar(20) unique not null,
    email varchar(50) unique not null,
    credential_id int not null unique,
    role_id int not null,
	CONSTRAINT FK_memeber_credential FOREIGN KEY (credential_id)
    REFERENCES credential(credential_id),
	CONSTRAINT FK_memeber_role FOREIGN KEY (role_id)
    REFERENCES role(role_id)
);

create table family (
	family_id int not null unique primary key auto_increment,
    family_name varchar(255) not null,
    family_info varchar(1000) not null,
    family_icon varchar(300)
);

show tables;

alter table member
add column created_at timestamp default current_timestamp;
alter table member
add column modified_at timestamp on update current_timestamp;

alter table family
add column created_at timestamp default current_timestamp;
alter table family
add column modified_at timestamp on update current_timestamp;

alter table credential
add column created_at timestamp default current_timestamp;
alter table credential
add column modified_at timestamp on update current_timestamp;

alter table role
add column created_at timestamp default current_timestamp;
alter table role
add column modified_at timestamp on update current_timestamp;

select * from member;
show tables;

use seap_db;
insert into credential (password) values ("$2a$10$JviVgZsoGjhFXCNwwWE8EO6wm.dVhCruceCasyKtw/y7UHZkpNhru");
select * from credential;
SELECT LAST_INSERT_ID();

select * from member;
alter table member modify credential_id varchar(255) not null;
update member set credential_id = "2b8b842f-79c0-4818-b35c-5f18062e7d5f" where username = "chenqingbao";
insert into credential (credential_id, password)
values
("2b8b842f-79c0-4818-b35c-5f18062e7d5f", "$2a$10$JviVgZsoGjhFXCNwwWE8EO6wm.dVhCruceCasyKtw/y7UHZkpNhru");

ALTER TABLE credential MODIFY credential_id varchar(255) NOT NULL;

alter table credential drop primary key;

describe credential;

delete from credential where credential_id = "5";

insert into member (
first_name, last_name, username, email, credential_id, role_id)
values
("Qing Bao", "Chen", "chenqingbao", "chen.qingbao22@gmail.com", 4, 2);

describe member;
alter table member add constraint fk_member_credential foreign key (credential_id) references credential (credential_id);
alter table member drop column member_id;
describe member;
alter table member add constraint pk_member primary key (username);
select * from family;