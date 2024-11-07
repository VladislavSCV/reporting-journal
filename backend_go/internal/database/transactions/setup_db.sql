-- Роли пользователей
create table roles (
                       id serial primary key,
                       value varchar(100) not null unique  -- Добавлен уникальный индекс на роль
);

-- Группы
create table groups (
                        id serial primary key,
                        name varchar(100) not null unique  -- Добавлен уникальный индекс на название группы
);

-- Пользователи
-- Пример структуры таблицы
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,  -- Если роль удалена, удаляются все пользователи с этой ролью
                       group_id INT REFERENCES groups(id) ON DELETE SET NULL,  -- Для преподавателей group_id может быть NULL
                       login VARCHAR(100) NOT NULL UNIQUE,  -- Уникальность логина
                       password TEXT NOT NULL,  -- Пароль теперь может быть длиннее (например, после хеширования)
                       salt TEXT NOT NULL  -- Соль для хеширования пароля
);

-- Заметки
create table notes (
                       id serial primary key,
                       title varchar(100) not null,
                       body text not null,
                       group_id int not null references groups(id) on delete cascade,
                       user_id int not null references users(id) on delete cascade
);

-- Предметы
create table subjects (
                          id serial primary key,
                          name varchar(100) not null unique  -- Уникальность на название предмета
);

-- Расписание
create table schedules (
                           id serial primary key,
                           group_id int not null references groups(id) on delete cascade,
                           day_of_week smallint not null check (day_of_week >= 1 and day_of_week <= 7),  -- Ограничение на дни недели
                           start_time time not null,
                           end_time time not null,
                           subject varchar(100) not null references subjects(name) on delete cascade,  -- Предмет должен ссылаться на таблицу subjects
                           teacher_id int not null references users(id) on delete cascade,  -- Преподаватель должен существовать в таблице пользователей
                           location varchar(100)  -- Место проведения
);

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
