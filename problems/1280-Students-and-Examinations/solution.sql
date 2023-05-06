-- Write your MySQL query statement below

select Stu.student_id, Stu.student_name, Sub.subject_name, count(Exam.subject_name) as attended_exams
from Students as Stu
inner join Subjects as Sub
left outer join Examinations as Exam
on Stu.student_id = Exam.student_id and Sub.subject_name = Exam.subject_name
group by Stu.student_id, Stu.student_name, Sub.subject_name
order by Stu.student_id ASC, Sub.subject_name ASC;
