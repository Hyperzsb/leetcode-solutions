# Write your MySQL query statement below

select tw.person_name
from (
    select person_name, @total_weight := @total_weight + weight as total_weight
    from Queue, (
        select @total_weight := 0
    ) as tmp 
    order by turn
) as tw
where tw.total_weight <= 1000
order by tw.total_weight desc
limit 1;

