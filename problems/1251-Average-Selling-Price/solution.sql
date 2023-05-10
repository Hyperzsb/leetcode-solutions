# Write your MySQL query statement below

select P.product_id, round(sum(P.price * US.units) / sum(US.units), 2) as average_price
from Prices as P
left outer join UnitsSold as US
on P.product_id = US.product_id and P.start_date <= US.purchase_date and P.end_date >= US.purchase_date
group by P.product_id;

