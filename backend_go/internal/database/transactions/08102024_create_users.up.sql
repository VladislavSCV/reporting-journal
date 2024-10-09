create table users (
                       id serial primary key,
                       name varchar(100) not null,
                       role_id int not null references roles(id),
                       group_id int references groups(id),  -- для студентов (для преподавателей может быть null)
                       login varchar(100) not null unique,
                       password varchar(100) not null
);

create table roles (
                       id serial primary key,
                       value varchar(100) not null unique
);

create table notes (
                       id serial primary key,
                       title varchar(100) not null,
                       body text not null,
                       group_id int not null references groups(id),
                       user_id int not null references users(id)
);

create table groups (
                        id serial primary key,
                        name varchar(100) not null,
                        description text
);

create table schedules (
                           id serial primary key,
                           group_id int not null references groups(id),
                           day_of_week smallint not null, -- от 1 (Понедельник) до 7 (Воскресенье)
                           start_time time not null,      -- время начала занятия
                           end_time time not null,        -- время окончания занятия
                           subject varchar(100) not null,
                           teacher_id int not null references users(id), -- ссылка на пользователя-преподавателя
                           location varchar(100),         -- место проведения
                           recurrence varchar(50)         -- повторяемость (если требуется)
);