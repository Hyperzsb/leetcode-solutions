# Write your MySQL query statement below

select P.project_id, round(sum(E.experience_years) / count(E.employee_id), 2) as average_years
from Project as P
left outer join Employee as E
on P.employee_id = E.employee_id
group by P.project_id;

