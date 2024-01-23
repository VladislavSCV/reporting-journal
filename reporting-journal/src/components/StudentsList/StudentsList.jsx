import React from "react";
import UserNav from "../UserNav/UserNav";
import { objectStudentsList } from "../../helpers/objectStudentsList";
import StudentCard from "../StudentCard/StudentCard";
import add from "./../../img/StudentsList/Add.svg";

import "./studentList.scss";

const StudentsList = () => {
  return (
    <div className="studentsList">
      <UserNav />
      <div className="studentsList__container">
        {objectStudentsList.map((obj, index) => {
          return (
            <StudentCard
              name={obj.name}
              surname={obj.surname}
              role={obj.role}
              key={index}
            />
          );
        })}
        <div className="studentsList__add" >
          <div className="studentsList__add-container" data-modal="ModalStudentAdd">
            <img src={add} alt="" className="studentsList__add-img" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default StudentsList;
