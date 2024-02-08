import React from "react";
import { objectStudentsList } from "../../helpers/objectStudentsList";
import StudentAttendanceCard from "../../components/StudentAttendanceCard/StudentAttendanceCard";

import "./studentAttendance.scss"

const StudentAttendance = () => {
  return (
    <div className="studentsList">
      <div className="studentsList__container" id="studentList">
        {objectStudentsList.map((obj, index) => {
          return (
            <StudentAttendanceCard
              surname={obj.surname}
              name={obj.name}
              patronymic={obj.patronymic}
              role={obj.role}
              key={index}
            />
          );
        })}
       
      </div>
    </div>
  );
};

export default StudentAttendance;
