import React, { useEffect, useState } from "react";
import StudentAttendanceCard from "../../components/StudentAttendanceCard/StudentAttendanceCard";
import axios from "axios";
import "./studentAttendance.scss";

const StudentAttendance = () => {
    const [users, setUsers] = useState([]);
    const [attendance, setAttendance] = useState({}); // Хранит выбранные значения

    useEffect(() => {
        const groupId = localStorage.getItem("group_id");
        if (!groupId) {
            console.error("Group ID is not found in localStorage");
            return;
        }

        const fetchGroups = async () => {
            try {
                const response = await axios.get(`/api/teacher/studentAttendance/${groupId}`, {
                    headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
                });
                setUsers(response.data.users || []);
            } catch (error) {
                console.error("Error fetching schedule:", error);
            }
        };

        fetchGroups();
    }, []);

    const handleAttendanceChange = (index, value) => {
        setAttendance((prev) => ({
            ...prev,
            [index]: value,
        }));

console.log(index)
        for (let i = 0; i < users.length; i++) {
            if (users[i].id === index) {
                logChange(users[i], value);
                break;
            }
        }
    };

    const logChange = (student, newStatus) => {
        console.log(`Изменение статуса для студента ${student.first_name} ${student.last_name}: ${newStatus}`);

        // Правильная структура запроса
        axios.post("/api/teacher/studentAttendance", {
            studentId: student.id,
            status: newStatus
        }, {
            headers: { Authorization: `Bearer ${localStorage.getItem("token")}` }
        })
            .then(response => {
                console.log("Статус успешно обновлен", response);
            })
            .catch(error => {
                console.error("Ошибка при обновлении статуса", error);
            });
    };


    return (
        <div className="studentAttendance">
            <div className="studentAttendance__container" id="studentList">
                {users.map((user) => (
                    <StudentAttendanceCard
                        key={user.id} // Используем уникальный идентификатор
                        id={user.id} // Передаем ID студента
                        surname={user.last_name}
                        name={user.first_name}
                        patronymic={user.middle_name}
                        role={user.role}
                        selectedValue={attendance[user.id] || ""} // Привязываем состояние к ID
                        onChange={(value) => handleAttendanceChange(user.id, value)} // Передаем обработчик
                    />
                ))}
            </div>
        </div>
    );
};

export default StudentAttendance;
