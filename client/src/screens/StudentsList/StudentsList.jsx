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
        const response = await axios.get("/api/user/students", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        console.log(response.data.students);
        setStudents(response.data.students || []);
      } catch (error) {
        console.error("Ошибка при запросе студентов:", error);
      }
    };
    fetchStudents();
  }, []);

console.log(students)

  return (
    <div className="studentsList">
      <div className="studentsList__container" id="studentList">
        {students
            // .filter(student => groupId ? student.group_id === groupId : true)
            .map(student => (
                <StudentCard
                    first_name={student.first_name + " " + student.middle_name + " " + student.last_name}
                    role={student.role}
                    key={student.id}
                    id={student.id}
                />
            ))}

        {/*<div*/}
        {/*  className="studentsList__add"*/}
        {/*  data-modal="ModalStudentAdd"*/}
        {/*  data-id={typeof groupId === "number" ? groupId.toString() : ""}*/}

        {/*  // data-id={groupId}*/}
        {/*>*/}
        {/*  /!*<img src={add} alt="" className="studentsList__add-img" />*!/*/}
        {/*</div>*/}
      </div>
    </div>
  );
};

export default StudentsList;
