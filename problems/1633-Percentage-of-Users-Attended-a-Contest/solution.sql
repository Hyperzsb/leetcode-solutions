# Write your MySQL query statement below

select R.contest_id, round(count(U.user_name) / (select count(*) from Users) * 100, 2) as percentage
from Register as R
left outer join Users as U
on R.user_id = U.user_id
group by R.contest_id
order by percentage desc, R.contest_id asc;

