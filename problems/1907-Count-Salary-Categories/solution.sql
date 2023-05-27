# Write your MySQL query statement below

select "Low Salary" as category, sum(income < 20000) as accounts_count
from Accounts
union
select "Average Salary" as category, sum(income >= 20000 and income <= 50000) as accounts_count
from Accounts
union
select "High Salary" as category, sum(income > 50000) as accounts_count
from Accounts

