import React from "react";
import "./StudentAttendanceCard.scss";

const StudentAttendanceCard = ({ surname, name, patronymic, role, selectedValue, onChange }) => {
    const handleSelectChange = (e) => {
        onChange(e.target.value); // Передаем новое значение посещаемости для студента
    };

    return (
        <div className="studentCard">
            <div className="studentCard__color-block"></div>
            <div className="studentCard__container">
                <div className="studentCard__student">
                    <p className="studentCard__student-name">
                        {surname} {name} {patronymic}
                    </p>
                    <p className="studentCard__student-role">Роль: {role}</p>
                </div>
                <div className="studentCard__attendanceBlock">
                    <label
                        htmlFor={`attendance-${surname}`} // Уникальный id для каждого студента
                        className="studentCard__attendanceBlock-label"
                    >
                        Посещение:{" "}
                    </label>
                    <select
                        id={`attendance-${surname}`} // Уникальный id для каждого студента
                        className="studentCard__attendanceBlock-attendance"
                        value={selectedValue} // Привязываем состояние к ID студента
                        onChange={handleSelectChange} // Обработчик изменений
                    >
                        <option value="">Выберите статус</option>
                        <option value="Присутствует">Присутствует</option>
                        <option value="Прогул">Прогул</option>
                        <option value="Болеет">Болеет</option>
                    </select>
                </div>
            </div>
        </div>
    );
};

export default StudentAttendanceCard;
