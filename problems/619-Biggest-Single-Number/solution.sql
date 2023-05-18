# Write your MySQL query statement below

select max((
    select num
    from (
        select num, count(*) as cnt
        from MyNumbers
        group by num
    ) as N
    where N.cnt = 1
    order by N.num desc
    limit 1
)) as num;

