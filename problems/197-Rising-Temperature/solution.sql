# Write your MySQL query statement below

select W1.id
from Weather as W1 
inner join Weather as W2
on W1.recordDate = DATE_ADD(W2.recordDate, interval 1 day)
where W1.temperature > W2.temperature;

