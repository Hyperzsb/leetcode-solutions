# Write your MySQL query statement below

select S.user_id, round(ifnull(CCnt.cnt, 0) / TCnt.cnt, 2) as confirmation_rate 
from Signups as S
left outer join (
  select S.user_id, count(*) as cnt
  from Signups as S
  left outer join Confirmations as C
  on S.user_id = C.user_id
  where C.action='confirmed'
  group by S.user_id
) as CCnt
on S.user_id = CCnt.user_id
left outer join (
  select S.user_id, count(*) as cnt
  from Signups as S
  left outer join Confirmations as C
  on S.user_id = C.user_id
  group by S.user_id
) as TCnt
on S.user_id = TCnt.user_id;

