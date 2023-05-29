# Write your MySQL query statement below

(
  select U.name as results
  from MovieRating as MR
  inner join Users as U
  on MR.user_id = U.user_id
  group by MR.user_id
  order by count(MR.movie_id) desc, U.name asc
  limit 1
) union all (
  select M.title as results
  from MovieRating as MR
  inner join Movies as M
  on MR.movie_id = M.movie_id
  where MR.created_at > '2020-01-31' and MR.created_at < '2020-03-01'
  group by MR.movie_id
  order by avg(MR.rating) desc, M.title asc
  limit 1
);

