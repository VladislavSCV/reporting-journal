import React from "react";
import settingsUser from "./../../assets/StudentCard/settings.svg";
import deleteUser from "./../../assets/StudentCard/delete.svg";
import "./studentCard.scss";
import axios from "axios";
const StudentCard = (obj) => {
  const deleteStudent = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/students/${key}`);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <div className="studentCard" key={obj.id}>
      <div className="studentCard__color-block"></div>
      <div className="studentCard__container">
        <div className="studentCard__student">
          <p className="studentCard__student-name">{obj.name}</p>
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
            // data-modal="ModalStudentDelete"
            onClick={() => deleteStudent(obj.id)}
            className="studentCard__buttons-delete"
          />
        </div>
      </div>
    </div>
  );
};

export default StudentCard;
