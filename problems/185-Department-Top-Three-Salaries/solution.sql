# Write your MySQL query statement below

select D1.name as Department, E1.name as Employee, E1.salary as Salary
from Employee as E1
inner join Department as D1
on E1.departmentId = D1.id
where (
    select count(distinct(E2.salary))
    from Employee as E2
    where E2.departmentId = E1.departmentId and E2.salary > E1.salary
) < 3;

