# Write your MySQL query statement below

select V.customer_id as customer_id, COUNT(V.customer_id) as count_no_trans from Visits as V left join Transactions as T using(visit_id) where transaction_id is null group by V.customer_id;

