# Write your MySQL query statement below

select S1.id, ifnull(S2.student, S1.student) as student
from Seat as S1
left outer join Seat as S2
on (
    case
    when S1.id % 2 = 1 then S1.id + 1 = S2.id 
    else S1.id - 1 = S2.id
    end
);

