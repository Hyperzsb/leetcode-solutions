# Write your MySQL query statement below

select round((
    select count(*)
    from (
        select player_id, min(event_date) as event_date
        from Activity
        group by player_id
    ) as A1
    inner join Activity as A2
    on A1.player_id = A2.player_id and A1.event_date = date_sub(A2.event_date, interval 1 day)
) / (
    select count(distinct player_id)
    from Activity
), 2) as fraction

