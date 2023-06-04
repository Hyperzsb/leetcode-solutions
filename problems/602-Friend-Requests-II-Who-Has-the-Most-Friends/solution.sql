 # Write your MySQL query statement below

select RA1.id, RA1.cnt + ifnull(RA2.cnt, 0) as num
from (
    select requester_id as id, count(accepter_id) as cnt
    from RequestAccepted
    group by requester_id
) as RA1
left outer join (
    select accepter_id as id, count(requester_id) as cnt
    from RequestAccepted
    group by accepter_id
) as RA2
on RA1.id = RA2.id
union
select RA2.id, ifnull(RA1.cnt, 0)  + RA2.cnt as num
from (
    select requester_id as id, count(accepter_id) as cnt
    from RequestAccepted
    group by requester_id
) as RA1
right outer join (
    select accepter_id as id, count(requester_id) as cnt
    from RequestAccepted
    group by accepter_id
) as RA2
on RA1.id = RA2.id
order by num desc
limit 1;

