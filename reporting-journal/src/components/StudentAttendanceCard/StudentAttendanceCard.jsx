import React from "react";
import "./StudentAttendanceCard.scss";
const StudentAttendanceCard = (obj) => {
  return (
    <div className="studentCard">
      <div className="studentCard__color-block"></div>
      <div className="studentCard__container">
        <div className="studentCard__student">
          <p className="studentCard__student-name">
            {obj.surname} {obj.name} {obj.patronymic}
          </p>
          <p className="studentCard__student-role">Роль: {obj.role}</p>
        </div>
        <div className="studentCard__attendanceBlock">
          <label
            htmlFor="studentAttendance"
            className="studentCard__attendanceBlock-label"
          >
            Посещение:{" "}
          </label>
          <select
            name=""
            id="studentAttendance"
            className="studentCard__attendanceBlock-attendance"
          >
            <option value="1"></option>
            <option value="2">Прогрул</option>
            <option value="3">Болеет</option>
          </select>
        </div>
      </div>
    </div>
  );
};

export default StudentAttendanceCard;
