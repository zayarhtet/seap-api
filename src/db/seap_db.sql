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
alter table family modify family_id varchar(255) not null;
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
use seap_db;
describe member;
describe role;
select * from member;
insert into member(first_name, last_name, username, email, credential_id, role_id)
values ("Zayar", "Htet", "admin", "zayarhtet797@gmail.com", "e0f5a784-aa4b-4523-a3bd-b4c01a6ca7e6", 99);
select * from role;
insert into role(role_id, name) values (99, "admin");
delete from `seap_db`.`role` where (`role_id` = 0);

DELETE FROM `seap_db`.`member` WHERE (`username` = 'admin');
select * from credential;
select * from member;
describe member;

show tables;
select * from family;
describe family;

create table family_member (
	username varchar(255) not null,
    family_id varchar(255) not null,
    role_id int not null,
    primary key (username, family_id),
	CONSTRAINT FK_familymember_member FOREIGN KEY (username)
    REFERENCES member(username) on delete cascade,
	CONSTRAINT FK_familymember_family FOREIGN KEY (family_id)
    REFERENCES family(family_id) on delete cascade,
    constraint FK_familymember_role foreign key (role_id)
    references role(role_id)
);

select * from family_member;
select * from member;
select * from family; describe family;
insert into family(family_id, family_name, family_info, family_icon)
values ("9dc1b896-4384-4cc8-bbcc-aaa773067153", "Object-oriented Programming", "23/24/2 Group 1", "/fp.png");

insert into family_member(username, family_id, role_id) values
("HELLO1", "803360bc-71f4-4b10-a119-ed93de707650", 2);

insert into family_member(username, family_id, role_id) values
("miyuki", "803360bc-71f4-4b10-a119-ed93de707650", 1);

insert into family_member(username, family_id, role_id) values
("HELLO2", "803360bc-71f4-4b10-a119-ed93de707650", 2);

insert into family_member(username, family_id, role_id) values
("HELLO3", "9dc1b896-4384-4cc8-bbcc-aaa773067153", 2);

insert into family_member(username, family_id, role_id) values
("HELLO4", "9dc1b896-4384-4cc8-bbcc-aaa773067153", 2);

insert into family_member(username, family_id, role_id) values
("chenqingbao", "9dc1b896-4384-4cc8-bbcc-aaa773067153", 1);

insert into family_member(username, family_id, role_id) values
("miyuki", "ff716cbb-501f-471b-b84c-fdc1b6cd6f16", 2);
insert into family_member(username, family_id, role_id) values
("chenqingbao", "ff716cbb-501f-471b-b84c-fdc1b6cd6f16", 1);
insert into family_member(username, family_id, role_id) values
("HELLO3", "ff716cbb-501f-471b-b84c-fdc1b6cd6f16", 1);

alter table family_member
add column created_at timestamp default current_timestamp;
alter table family_member
add column modified_at timestamp on update current_timestamp;
select * from family_member;
ALTER TABLE family_member
  DROP FOREIGN KEY FK_familymember_family;
  
select * from family_member;
select * from member;
delete from member where username = 'HELLO3';
delete from family_member where username = 'HELLO3';
alter table family_member add constraint FK_familymember_family foreign key (family_id) references family (family_id) on delete cascade;

SELECT *
FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
WHERE TABLE_NAME = 'family_member';

use seap_db;
show tables;
describe credential;

create table duty (
	duty_id varchar(255) primary key not null,
    title varchar(255) not null,
    instruction varchar(1000),
    publishing_date timestamp not null,
    deadline_date timestamp not null,
    closing_date timestamp not null,
    family_id varchar(255) not null,
    point_system bool default true,
    possible_points double,
    multipleSubmission bool default true,
	CONSTRAINT FK_duty_family FOREIGN KEY (family_id)
    REFERENCES family(family_id) on delete cascade
);

select * from duty;

create table grading (
	username varchar(255) not null,
    duty_id varchar(255) not null,
    family_id varchar(255) not null,
    submitted bool default false,
    points double,
    is_late bool default false,
    is_passed bool default false,
    grade_comment varchar(1000),
    execution_comment varchar(2000),
    primary key (username, duty_id, family_id),
	CONSTRAINT FK_grading_member FOREIGN KEY (username)
    REFERENCES member(username) on delete cascade,
	CONSTRAINT FK_grading_family FOREIGN KEY (family_id)
    REFERENCES family(family_id) on delete cascade,
    constraint FK_grading_duty foreign key (duty_id)
    references duty(duty_id)
);
describe grading;
alter table grading drop primary key;

alter table grading drop constraint FK_grading_member;
alter table grading drop constraint FK_grading_duty;
alter table grading add constraint FK_grading_member FOREIGN KEY (username)
    REFERENCES member(username) on delete cascade;
alter table grading drop constraint FK_grading_family;
alter table grading add constraint FK_grading_family FOREIGN KEY (family_id)
    REFERENCES family(family_id) on delete cascade;
alter table grading drop constraint FK_grading_duty;
alter table grading add constraint FK_grading_duty foreign key (duty_id)
    references duty(duty_id) on delete cascade;

alter table grading add column grading_id varchar(255) primary key not null;

select * from grading;
alter table grading add column submitted_at timestamp;

describe duty;
select * from duty;
delete from duty where duty_id = "cd8cc9bb-0724-424d-bd94-999195e3cc84";
insert into duty (
duty_id, family_id, title, publishing_date, deadline_date, closing_date, possible_points, instruction ) values (
"44062536-0d5c-422d-9dd2-b4b03fb7b3df", "ff716cbb-501f-471b-b84c-fdc1b6cd6f16", "PT1",
TIMESTAMP('2024-03-16 14:30:45'), TIMESTAMP('2024-03-17 14:30:45'), TIMESTAMP('2024-03-17 14:30:45'), 100, "download the file and solve.");

insert into duty (
duty_id, family_id, title, publishing_date, deadline_date, closing_date, possible_points, instruction ) values (
"7cb57570-7e36-4967-8d54-b5b65dca9531", "9dc1b896-4384-4cc8-bbcc-aaa773067153", "Assignment II.",
TIMESTAMP('2024-03-16 14:30:45'), TIMESTAMP('2024-03-23 14:30:45'), TIMESTAMP('2024-03-23 14:30:45'), 100, "download the file and solve.");

insert into duty (
duty_id, family_id, title, publishing_date, deadline_date, closing_date, possible_points, instruction ) values (
"f21feab1-914b-4107-b19c-07be74ebd7f6", "803360bc-71f4-4b10-a119-ed93de707650", "Quiz I.",
TIMESTAMP('2024-03-16 14:30:45'), TIMESTAMP('2024-03-23 14:30:45'), TIMESTAMP('2024-03-23 14:30:45'), 100, "download the file and solve.");

create table submitted_file (
	file_id varchar(255) primary key not null,
    grading_id varchar(255) not null,
    file_path varchar(1000) not null,
	submitted_at timestamp default current_timestamp,
	constraint submitted_file_grading foreign key (grading_id) references grading(grading_id) on delete cascade
);

create table given_file (
	file_id varchar(255) primary key not null,
    duty_id varchar(255) not null,
    file_path varchar(1000) not null,
    constraint given_file_duty foreign key (duty_id) references duty(duty_id) on delete cascade
);
insert into given_file (file_id, duty_id, file_path)
values ("9ef621b5-3186-4f3e-acd9-69c6406ab24f", "cd8cc9bb-0724-424d-bd94-999195e3cc84", "./brnyr/HW1.icl");

insert into given_file (file_id, duty_id, file_path)
values ("dd0d11e9-c2b5-401b-bcd8-980bf7661604", "cd8cc9bb-0724-424d-bd94-999195e3cc84", "./brnyr/HW12.icl");

insert into given_file (file_id, duty_id, file_path)
values ("de36545b-b0bc-4fb7-91c1-dbbc62e16ce5", "dc80e248-6afe-4638-88b0-a67f3c527c7b", "./brnyr/HW2.icl");
use seap_db;
select * from grading;
describe grading;
insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"a4b27bd4-7924-4805-92bb-1d90d92eb4b5",
    "HELLO1",
    "f21feab1-914b-4107-b19c-07be74ebd7f6",
    "803360bc-71f4-4b10-a119-ed93de707650",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"ab8f55af-4213-49a1-8a6a-a2f316f71b99",
    "HELLO2",
    "f21feab1-914b-4107-b19c-07be74ebd7f6",
    "803360bc-71f4-4b10-a119-ed93de707650",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"0374545a-f686-43e0-bd2c-6fcc628d7f96",
    "HELLO3",
    "bed5c9f8-4813-4e02-b903-0b6b0a942503",
    "9dc1b896-4384-4cc8-bbcc-aaa773067153",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"c6ea6524-ce79-4583-82c6-8e77975cfca4",
    "HELLO3",
    "7cb57570-7e36-4967-8d54-b5b65dca9531",
    "9dc1b896-4384-4cc8-bbcc-aaa773067153",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"c3dba5cb-9910-4321-af3c-f14f97483af1",
    "miyuki",
    "44062536-0d5c-422d-9dd2-b4b03fb7b3df",
    "ff716cbb-501f-471b-b84c-fdc1b6cd6f16",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"0ec8d8e4-8885-46ad-bfbd-8f9e89e1eaa2",
    "miyuki",
    "cd8cc9bb-0724-424d-bd94-999195e3cc84",
    "ff716cbb-501f-471b-b84c-fdc1b6cd6f16",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);

insert into grading (
grading_id, username, duty_id, family_id, submitted, points, is_passed, grade_comment, execution_comment, submitted_at
) values (
	"8741cedc-8fae-463b-8447-f5ac1023a2c4",
    "miyuki",
    "dc80e248-6afe-4638-88b0-a67f3c527c7b",
    "ff716cbb-501f-471b-b84c-fdc1b6cd6f16",
    true,
    100,
    true,
    "GOOD",
    "good",
     TIMESTAMP(current_timestamp)
);
use seap_db;
describe grading;
describe duty;

ALTER TABLE duty 
RENAME COLUMN multipleSubmission TO multiple_submission;

use seap_db;
show tables;
select * from grading;
select * from family_member;
describe family_member;
select * from member;
select * from family;
select * from duty;
select * from member;
alter table duty add column plugin_name varchar(255);

update duty set plugin_name="fpclean" where plugin_name = null;

select * from credential where credential_id = "47186c71-d626-4ccf-aa68-db183464f661";
show tables; describe given_file; describe submitted_file;
describe grading;
alter table grading add column hasGraded bool default false;
update credential set password = "$2a$10$d4HXxxg61rAzDQ9KsPdLUucHNf6qhtBOiQ1nc3QRwJfV4gLKvdvjm" where credential_id = "e0f5a784-aa4b-4523-a3bd-b4c01a6ca7e6";

select * from duty where duty_id = "40736a34-d2f2-470c-a5ce-51f9f5804e99";
describe duty;

select * from given_file;

select * from submitted_file;
member (member_id, profilepath)

describe given_file;

file_table (file_id, path)
given_file_table (dutyId, file_id)
submitted_file_table (grading_id, file_id, submitted_at)
member (member_id, file_id);

alter table grading rename column hasGraded to has_graded;

DELETE FROM `duty` WHERE `duty`.`duty_id` = 'ff1bff55-f8cc-4316-b046-7f6d2d05f68f';
select * from credential where credential_id = "f72d1529-327c-4f2d-8fe1-6bf777b7fd25";