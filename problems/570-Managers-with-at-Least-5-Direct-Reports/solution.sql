-- Write your MySQL query statement below

select Result.name
from (
  select E1.name, count(*) as cnt
  from Employee as E1
  right outer join (
    select * from
    Employee
    where Employee.managerId is not null
  ) as E2
  on E1.id = E2.managerId
  group by E1.name
) as Result
where Result.name is not null and Result.cnt >= 5;

