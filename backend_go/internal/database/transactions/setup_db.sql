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
                       first_name VARCHAR(100) NOT NULL,
                       middle_name VARCHAR(100) NOT NULL,
                       last_name VARCHAR(100),
                       role_id INT NOT NULL REFERENCES roles(id), -- ON DELETE CASCADE Если роль удалена, удаляются все пользователи с этой ролью
                       group_id INT REFERENCES groups(id) ON DELETE SET NULL,  -- Для преподавателей group_id может быть NULL
                       login VARCHAR(100) NOT NULL UNIQUE,  -- Уникальность логина
                       password TEXT NOT NULL,  -- Пароль теперь может быть длиннее (например, после хеширования)
                       salt TEXT NOT NULL  -- Соль для хеширования пароля
                       -- token TEXT NOT NULL -- Токен для авторизации
);

-- Таблица для связи преподавателей и групп
CREATE TABLE teacher_groups (
                                id SERIAL PRIMARY KEY,  -- Уникальный идентификатор записи
                                teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,  -- ID преподавателя
                                group_id INT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,  -- ID группы
                                UNIQUE (teacher_id, group_id)  -- Уникальная пара преподаватель-группа
);


-- Заметки
create table notes (
                       id serial primary key,
                       title varchar(100) not null,
                       body text not null,
                       group_id INT REFERENCES groups(id) ON DELETE SET NULL,
                       user_id int not null references users(id) on delete cascade
);

-- Таблица предметов
CREATE TABLE subjects (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL UNIQUE  -- Уникальность на название предмета
);

-- Расписание
CREATE TABLE schedules (
                           id SERIAL PRIMARY KEY,
                           group_id INT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
                           day_of_week SMALLINT NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),  -- Ограничение на дни недели (1-7)
                           subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,  -- Предмет связан с ID из таблицы subjects
                           teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,  -- Преподаватель связан с таблицей пользователей
--                         -- TODO лучше перевести в тип int
                           location VARCHAR(100)  -- Место проведения
);

CREATE TABLE attendance (
                            id SERIAL PRIMARY KEY,
                            student_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            date DATE NOT NULL,
                            status VARCHAR(20) NOT NULL, -- Статус посещения
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
