import React from "react";
import settingsUser from "./../../assets/StudentCard/settings.svg";
import deleteUserSVG from "./../../assets/StudentCard/delete.svg";
import "./studentCard.scss";
import {deleteUser} from "../../actions/api";

const StudentCard = (obj) => {
  return (
    <div className="studentCard" key={obj.id}>
      <div className="studentCard__color-block"></div>
      <div className="studentCard__container">
        <div className="studentCard__student">
          <p className="studentCard__student-name">{obj.first_name}</p>
          {/*<p className="studentCard__student-name">{obj.middle_name}</p>*/}
          {/*<p className="studentCard__student-name">{obj.last_name}</p>*/}
          <p className="studentCard__student-role">Роль: {obj.role}</p>
        </div>
        <div className="studentCard__buttons">
        <img
            src={settingsUser}
            alt=""
            className="studentCard__buttons-settings"
            data-modal="ModalStudentSettings"
            data-id={obj.id}
          />

          <img
            src={deleteUserSVG}
            alt=""
            // data-modal="ModalStudentDelete"
            onClick={() => deleteUser(obj.id)}
            className="studentCard__buttons-delete"
          />
        </div>
      </div>
    </div>
  );
};

export default StudentCard;
