# Write your MySQL query statement below

select P1.product_id, ifnull(P2.price, 10) as price
from Products as P1
left outer join (
    select P3.product_id, new_price as price
    from Products as P3
    inner join (
        select product_id, max(change_date) as last_date
        from Products
        where change_date <= '2019-08-16'
        group by product_id
    ) as P4
    on P3.product_id = P4.product_id and P3.change_date = P4.last_date
) as P2
on P1.product_id = P2.product_id
group by P1.product_id;

