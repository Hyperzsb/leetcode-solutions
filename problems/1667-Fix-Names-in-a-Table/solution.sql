# Write your MySQL query statement below

select user_id, concat(ucase(substr(name, 1, 1)), lcase(substr(name, 2, length(name)-1))) as name
from Users
order by user_id

