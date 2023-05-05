# Write your MySQL query statement below

select E.name, B.bonus
from Employee as E
left outer join Bonus as B
on E.empId = B.empId
where B.bonus < 1000 or B.bonus is null;

