-- Добавление данных в таблицы
INSERT INTO roles (value) VALUES ('student');
INSERT INTO roles (value) VALUES ('teacher');
INSERT INTO roles (value) VALUES ('admin');

INSERT INTO groups (name) VALUES ('22ИС3-2');
INSERT INTO groups (name) VALUES ('22ИС3-1');
-- Дополнительные группы можно добавить по мере необходимости

INSERT INTO subjects (name) VALUES ('Математика');
INSERT INTO subjects (name) VALUES ('Физика');
INSERT INTO subjects (name) VALUES ('Программирование');
-- Добавление предметов по мере необходимости

-- Группы
INSERT INTO groups (name) VALUES
                              ('Group A'),
                              ('Group B'),
                              ('Group C');

-- Пользователи
INSERT INTO users (first_name, middle_name, last_name, role_id, group_id, login, password, salt)
VALUES
    ('Alice', 'Jane', 'Smith', 1, NULL, 'admin1', 'hashed_password1', 'salt1'),
    ('Bob', 'John', 'Doe', 2, NULL, 'teacher1', 'hashed_password2', 'salt2'),
    ('Charlie', 'Chris', 'Brown', 3, 1, 'student1', 'hashed_password3', 'salt3'),
    ('David', 'Adam', 'Johnson', 3, 1, 'student2', 'hashed_password4', 'salt4'),
    ('Eve', 'Marie', 'Clark', 3, 2, 'student3', 'hashed_password5', 'salt5');

-- Привязка преподавателей к группам
INSERT INTO teacher_groups (teacher_id, group_id) VALUES
                                                      (2, 1),
                                                      (2, 2);

-- Предметы
INSERT INTO subjects (name) VALUES
                                ('Mathematics'),
                                ('Physics'),
                                ('History');

-- Расписание
INSERT INTO schedules (group_id, day_of_week, start_time, end_time, subject_id, teacher_id, location)
VALUES
    (1, 1, '09:00', '10:30', 1, 2, 'Room 101'),
    (1, 3, '11:00', '12:30', 2, 2, 'Room 102'),
    (2, 5, '14:00', '15:30', 3, 2, 'Room 201');

-- Заметки
INSERT INTO notes (title, body, group_id, user_id) VALUES
                                                       ('Homework 1', 'Complete exercises 1-10', 1, 3),
                                                       ('Exam info', 'Exam will be on Monday at 10:00', 1, 3),
                                                       ('Project Guidelines', 'Submit the project by next week', 2, 5);
