import React from "react";
import { objectStudentsList } from "../../helpers/objectStudentsList";
import StudentCard from "../../components/StudentCard/StudentCard";
import add from "./../../assets/StudentsList/Add.svg";

import "./studentList.scss";

const StudentsList = () => {
  return (
    <div className="studentsList">
      <div className="studentsList__container" id="studentList">
        {objectStudentsList.map((obj, index) => {
          return (
            <StudentCard
              surname={obj.surname}
              name={obj.name}
              patronymic={obj.patronymic}
              role={obj.role}
              key={index}
            />
          );
        })}
        <div className="studentsList__add" data-modal="ModalStudentAdd">
          <img src={add} alt="" className="studentsList__add-img" />
        </div>
      </div>
    </div>
  );
};

export default StudentsList;
