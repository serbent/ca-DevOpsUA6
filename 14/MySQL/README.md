```
use mydb;
show tables;
create table test_1(id int);
show tables; 
describe table test_1 ;
alter table test_1 add name varchar(255);
escribe test_1;
insert into test_1 (id) values (1);
drop table test_1;
drop database mydb;
show databases;


 docker exec -it -p ${PWD}:/app -w /app mysql mysql -u root -p < /app/mysql.sql
select u.name as u_name, s.first_name as s_name, s.last_name as l_name, d.name as dormname from dorms d, students s, universities u where u
.id = s.university_id and d.id = s.dorm_id;
```