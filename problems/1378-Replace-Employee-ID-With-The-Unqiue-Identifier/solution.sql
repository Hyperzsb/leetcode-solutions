# Write your MySQL query statement below

select EUNI.unique_id, E.name from Employees as E left join EmployeeUNI as EUNI on E.id=EUNI.id;

