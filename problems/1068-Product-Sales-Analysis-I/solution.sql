# Write your MySQL query statement below

#select P.product_name, S.year, S.price from Sales as S, Product as P where S.product_id = P.product_id;

select P.product_name, S.year, S.price from Sales as S inner join Product as P using(product_id);

