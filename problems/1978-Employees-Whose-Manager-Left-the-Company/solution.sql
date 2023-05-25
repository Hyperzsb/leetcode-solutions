# Write your MySQL query statement below

select employee_id
from Employees
where salary < 30000 and manager_id not in (
    select employee_id
    from EMployees
    group by employee_id
)
order by employee_id;

