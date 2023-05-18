# Write your MySQL query statement below

select C.class
from (
  select class, count(student) as cnt
  from Courses
  group by class
) as C
where C.cnt >= 5;

