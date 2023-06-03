# Write your MySQL query statement below

select CA.visited_on as visited_on, sum(CB.day_sum) as amount, round(avg(CB.day_sum), 2) as average_amount
from (
    select visited_on, sum(amount) as day_sum
    from Customer
    group by visited_on
) as CA, (
    select visited_on, sum(amount) as day_sum
    from Customer
    group by visited_on
) as CB
where datediff(CA.visited_on, CB.visited_on) between 0 and 6
group by CA.visited_on
having count(CB.visited_on) = 7;

