-- Индексы для таблицы users
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_group_id ON users(group_id);
CREATE INDEX idx_users_login ON users(login); -- Хотя login уже уникален, явный индекс ускорит поиск
CREATE INDEX idx_users_token ON users(token); -- Для быстрого поиска по токенам (если используется)

-- Индексы для таблицы teacher_groups
CREATE INDEX idx_teacher_groups_teacher_id ON teacher_groups(teacher_id);
CREATE INDEX idx_teacher_groups_group_id ON teacher_groups(group_id);

-- Индексы для таблицы groups
CREATE INDEX idx_groups_name ON groups(name);

-- Индексы для таблицы notes
CREATE INDEX idx_notes_group_id ON notes(group_id);
CREATE INDEX idx_notes_user_id ON notes(user_id);

-- Индексы для таблицы schedules
CREATE INDEX idx_schedules_group_id ON schedules(group_id);
CREATE INDEX idx_schedules_subject_id ON schedules(subject_id);
CREATE INDEX idx_schedules_teacher_id ON schedules(teacher_id);
CREATE INDEX idx_schedules_day_of_week ON schedules(day_of_week);
CREATE INDEX idx_schedules_time_range ON schedules(start_time, end_time);

-- Индексы для таблицы subjects
CREATE INDEX idx_subjects_name ON subjects(name);
