# Write your MySQL query statement below

select Q1.query_name, round(sum(Q2.quality) / count(Q1.result), 2) as quality, round((select count(*) from Queries as Q3 where Q1.query_name = Q3.query_name and Q3.rating < 3) / count(Q1.result) * 100, 2) as poor_query_percentage
from Queries as Q1
inner join (
    select query_name, result, (rating / position) as quality
    from Queries
) as Q2
on Q1.query_name = Q2.query_name and Q1.result = Q2.result
group by Q1.query_name

