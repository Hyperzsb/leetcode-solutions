# Write your MySQL query statement below

select S1.product_id, S1.year as first_year, S1.quantity, S1.price
from Sales as S1
inner join (
    select product_id, min(year) as year
    from Sales
    group by product_id
) as S2
on S1.product_id = S2.product_id and S1.year = S2.year;

