import React from "react";
import settingsUser from "./../../img/StudentCard/settings.svg";
import deleteUser from "./../../img/StudentCard/delete.svg";
import "./studentCard.scss";
const StudentCard = (obj) => {
  return (
    <div className="studentCard">
      <div className="studentCard__color-block"></div>
      <div className="studentCard__container">
        <p className="studentCard__student">
          {obj.name} {obj.surname}
        </p>
        <p className="studentCard__student-role">Роль: {obj.role}</p>
        <div className="studentCard__buttons">
     
            
            <img
              src={settingsUser}
              alt=""
              className="studentCard__settings-btn"
              data-modal="ModalStudentSettings"/>
   
          <img src={deleteUser} alt="" className="studentCard__delete-btn" />
        </div>
      </div>
    </div>
  );
};

export default StudentCard;
