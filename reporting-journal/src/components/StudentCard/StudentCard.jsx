import React from "react";
import settingsUser from "./../../assets/StudentCard/settings.svg";
import deleteUser from "./../../assets/StudentCard/delete.svg";
import "./studentCard.scss";
const StudentCard = (obj) => {
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
        <div className="studentCard__buttons">
          <img
            src={settingsUser}
            alt=""
            className="studentCard__buttons-settings"
            data-modal="ModalStudentSettings"
          />

          <img
            src={deleteUser}
            alt=""
            data-modal="ModalStudentDelete"
            className="studentCard__buttons-delete"
          />
        </div>
      </div>
    </div>
  );
};

export default StudentCard;
