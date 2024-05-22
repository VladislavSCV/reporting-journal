import React, { useEffect, useState } from "react";
import StudentCard from "../../components/StudentCard/StudentCard";
import add from "./../../assets/StudentsList/Add.svg";
import axios from "axios";
import "./studentList.scss";

const StudentsList = () => {
  const [students, setStudents] = useState([]);
  let groupId = Number(window.location.search.substring(1).split("=")[1]);
  useEffect(() => {
    const fetchStudents = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/students");
        setStudents(response.data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchStudents();
  }, []);

  return (
    <div className="studentsList">
      <div className="studentsList__container" id="studentList">
        {students.map((stundet) => {
          if (stundet.groupId === groupId) {
            return (
              <StudentCard
                name={stundet.name}
                role={stundet.role}
                key={stundet.id}
                id={stundet.id}
              />
            );
          }
        })}
        <div
          className="studentsList__add"
          data-modal="ModalStudentAdd"
          data-id={groupId}
        >
          <img src={add} alt="" className="studentsList__add-img" />
        </div>
      </div>
    </div>
  );
};

export default StudentsList;
