import React, { useEffect, useState } from "react";
import StudentCard from "../../components/StudentCard/StudentCard";
import add from "./../../assets/StudentsList/Add.svg";
import axios from "axios";
import "./studentList.scss";
import { useSearchParams } from "react-router-dom";

const StudentsList = () => {
  const [students, setStudents] = useState([]);
  const [searchParams] = useSearchParams(); // Используем хук для работы с параметрами строки запроса
  const groupId = Number(searchParams.get("id")); // Извлекаем параметр id из строки

  useEffect(() => {
    const fetchStudents = async () => {
      try {
        const response = await axios.get("/api/user/students", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        console.log(response.data.students);
        setStudents(response.data.students || []);
      } catch (error) {
        console.error("Ошибка при запросе студентов:", error);
      }
    };

    fetchStudents();
  }, []); // Запрос отправляется один раз при монтировании компонента

  const handleDelete = async (id) => {
    try {
      await axios.delete(`/api/user/${id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      // Убираем удаленного студента из состояния
      setStudents((prevStudents) => prevStudents.filter((student) => student.id !== id));
    } catch (error) {
      console.error("Ошибка при удалении студента:", error);
    }
  };


  return (
      <div className="studentsList">
        <div className="studentsList__container" id="studentList">
          {students
              .filter((student) => (groupId ? student.group_id === groupId : true)) // Фильтрация по groupId, если он есть
              .map((student) => (
                  <StudentCard
                      first_name={`${student.first_name} ${student.middle_name} ${student.last_name}`}
                      role={student.role}
                      key={student.id}
                      id={student.id}
                      onDelete={handleDelete}
                  />

              ))}

          <div
              className="studentsList__add"
              data-modal="ModalStudentAdd"
              data-id={groupId ? groupId.toString() : ""}
          >
            <img src={add} alt="Добавить" className="studentsList__add-img" />
          </div>
        </div>
      </div>
  );
};

export default StudentsList;
