# Write your MySQL query statement below

select round(sum(T.tiv_2016), 2) as tiv_2016
from (
  select I1.tiv_2016
  from Insurance as I1
  inner join Insurance as I2
  on I1.pid <> I2.pid and I1.tiv_2015 = I2.tiv_2015
  where I1.pid not in (
    select I3.pid
    from Insurance as I3
    inner join Insurance as I4
    on I3.pid <> I4.pid and I3.lat = I4.lat and I3.lon = I4.lon
    group by I3.pid
  )
  group by I1.pid
) as T;

