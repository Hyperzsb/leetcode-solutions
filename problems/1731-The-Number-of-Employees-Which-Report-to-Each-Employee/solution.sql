# Write your MySQL query statement below

select E1.employee_id, E1.name, E2.reports_count, E2.average_age
from Employees as E1
inner join (
    select reports_to, count(employee_id) as reports_count, round(avg(age), 0) as average_age
    from Employees
    group by reports_to
) as E2
on E1.employee_id = E2.reports_to
order by E1.employee_id;

