# Write your MySQL query statement below

select P.product_name, sum(O.unit) as unit
from Products as P
inner join (
    select product_id, unit
    from Orders
    where order_date >= '2020-02-01' and order_date <= '2020-02-29'
) as O
on P.product_id = O.product_id
group by P.product_id
having sum(O.unit) >= 100;

