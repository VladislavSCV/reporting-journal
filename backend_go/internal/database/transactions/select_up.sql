----------------------------------------
-- Получение групп, которые ведет преподаватель
SELECT g.name AS group_name
FROM groups g
         JOIN teacher_groups tg ON g.id = tg.group_id
WHERE tg.teacher_id = 1;

-- Получение преподавателей группы
SELECT u.first_name, u.last_name
FROM users u
         JOIN teacher_groups tg ON u.id = tg.teacher_id
WHERE tg.group_id = 1;
----------------------------------------