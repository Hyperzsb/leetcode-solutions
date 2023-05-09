# Write your MySQL query statement below

select *
from Cinema as C
where C.id % 2 = 1 and C.description <> 'boring'
order by C.rating desc;

