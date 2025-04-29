-- Вставка ролей
INSERT INTO roles (value) VALUES 
('student'), 
('teacher'), 
('admin');

-- Вставка групп
INSERT INTO groups (name) VALUES
('Group A'), ('Group B'), ('Group C'), ('Group D'), ('Group E'),
('Group F'), ('Group G'), ('Group H'), ('Group I'), ('Group J'),
('Group K'), ('Group L'), ('Group M'), ('Group N'), ('Group O'),
('Group P'), ('Group Q'), ('Group R'), ('Group S'), ('Group T'),
('Group U'), ('Group V'), ('Group W'), ('Group X'), ('Group Y'),
('Group Z'), ('Group AA'), ('Group AB'), ('Group AC'), ('Group AD'),
('Group AE'), ('Group AF'), ('Group AG'), ('Group AH'), ('Group AI'),
('Group AJ'), ('Group AK'), ('Group AL'), ('Group AM'), ('Group AN'),
('Group AO'), ('Group AP'), ('Group AQ'), ('Group AR'), ('Group AS'),
('Group AT'), ('Group AU'), ('Group AV'), ('Group AW'), ('Group AX');

-- Вставка предметов
INSERT INTO subjects (name) VALUES
('Math'), ('Physics'), ('Chemistry'), ('Biology'), ('History'),
('Geography'), ('English'), ('Programming'), ('Databases'), ('Algorithms'),
('Data Structures'), ('Networking'), ('Operating Systems'), ('Cybersecurity'), ('AI'),
('Machine Learning'), ('Robotics'), ('Statistics'), ('Philosophy'), ('Economics'),
('Business'), ('Accounting'), ('Art'), ('Design'), ('Music'),
('Literature'), ('Drama'), ('Sociology'), ('Psychology'), ('Law'),
('Ethics'), ('Astronomy'), ('Genetics'), ('Medicine'), ('Ecology'),
('Geology'), ('Meteorology'), ('Anthropology'), ('Linguistics'), ('Cryptography'),
('Mobile Dev'), ('Web Dev'), ('Backend Dev'), ('Frontend Dev'), ('Game Dev'),
('UI/UX'), ('3D Modeling'), ('Cloud Computing'), ('DevOps'), ('Blockchain');

-- Вставка пользователей (по 50)
-- Роли: 1 — student, 2 — teacher, 3 — admin
INSERT INTO users (first_name, middle_name, last_name, role_id, group_id, login, password, salt) VALUES
('John', 'Michael', 'Smith', 1, 1, 'johnsmith1', 'pass1', 'salt1'),
('Jane', 'Anna', 'Doe', 2, NULL, 'janedoe2', 'pass2', 'salt2'),
('Mark', 'Lee', 'Brown', 3, 2, 'markbrown3', 'pass3', 'salt3'),
('Emily', 'Grace', 'Davis', 1, 3, 'emilyd4', 'pass4', 'salt4'),
('Robert', 'James', 'Miller', 1, 4, 'robertm5', 'pass5', 'salt5'),
-- добавь ещё 45 строк аналогично для 50 пользователей...

-- Вставка teacher_groups (связь учителей с группами)
INSERT INTO teacher_groups (teacher_id, group_id) VALUES
(2, 1),
(2, 2),
(2, 3),
(2, 4),
(2, 5),
(2, 6),
(2, 7),
(2, 8),
(2, 9),
(2, 10),
-- ещё 40 связей...

-- Вставка schedules
INSERT INTO schedules (group_id, day_of_week, subject_id, teacher_id, location) VALUES
(1, 1, 1, 2, 'Room 101'),
(2, 2, 2, 2, 'Room 102'),
(3, 3, 3, 2, 'Room 103'),
(4, 4, 4, 2, 'Room 104'),
(5, 5, 5, 2, 'Room 105'),
-- ещё 45 строк...

-- Вставка notes
INSERT INTO notes (title, body, group_id, user_id) VALUES
('Reminder 1', 'This is a note for group 1.', 1, 1),
('Reminder 2', 'This is a note for group 2.', 2, 1),
('Reminder 3', 'This is a note for group 3.', 3, 1),
-- ещё 47 заметок...

-- Вставка attendance
INSERT INTO attendance (student_id, date, status) VALUES
(1, '2025-04-01', 'present'),
(1, '2025-04-02', 'absent'),
(1, '2025-04-03', 'present'),
(2, '2025-04-01', 'present'),
(2, '2025-04-02', 'present'),
-- ещё 45 записей...

