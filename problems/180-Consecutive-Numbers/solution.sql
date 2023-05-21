# Write your MySQL query statement below

select L1.num as ConsecutiveNums
from Logs as L1
inner join Logs as L2
on L1.id + 1 = L2.id and L1.num = L2.num
inner join Logs as L3
on L1.id + 2 = L3.id and L1.num = L3.num
group by L1.num;

