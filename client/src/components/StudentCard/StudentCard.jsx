import React from "react";
import settingsUser from "./../../assets/StudentCard/settings.svg";
import deleteUserSVG from "./../../assets/StudentCard/delete.svg";
import "./studentCard.scss";
import {deleteUser} from "../../actions/api";

const StudentCard = (obj) => {
  const { onDelete, id } = obj
  const role = localStorage.getItem("userRole")
  let adminPrivileges = false

  if (role === "Admin") {
    adminPrivileges = true
  }
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


          {adminPrivileges && (
              <img
              src={settingsUser}
              alt=""
              className="studentCard__buttons-settings"
              data-modal="ModalStudentSettings"
              data-id={obj.id}
          /> &&
            <img
            src={deleteUserSVG}
          alt=""
          onClick={() => onDelete(id)}
          className="studentCard__buttons-delete"
        />
          )}


        </div>
      </div>
    </div>
  );
};

export default StudentCard;
