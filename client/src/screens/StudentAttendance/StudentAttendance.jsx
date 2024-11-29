import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import StudentAttendanceCard from "../../components/StudentAttendanceCard/StudentAttendanceCard";
import "./studentAttendance.scss";

const StudentAttendance = () => {
    const { groupId } = useParams();
    const [users, setUsers] = useState([]);
    const [attendance, setAttendance] = useState({});
    const [errorMessage, setErrorMessage] = useState("");

    useEffect(() => {
        if (!groupId) {
            setErrorMessage("Ошибка: groupId не найден.");
            return;
        }

        const fetchStudents = async () => {
            try {
                const response = await fetch(`/api/teacher/studentAttendance/${groupId}`, {
                    method: "GET",
                    headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
                });

                if (!response.ok) {
                    throw new Error(`Ошибка запроса: ${response.status}`);
                }

                const responseData = await response.json();

                if (responseData.users) {
                    const decodedData = JSON.parse(atob(responseData.users));
                    setUsers(decodedData || []);
                } else {
                    setUsers([]);
                    console.error("Поле 'users' отсутствует или некорректное в ответе");
                }
            } catch (error) {
                console.error("Ошибка при загрузке студентов:", error);
                setErrorMessage("Ошибка при загрузке студентов.");
            }
        };

        fetchStudents();
    }, [groupId]);

    console.log(users);

    const handleAttendanceChange = (studentId, value) => {
        setAttendance((prev) => ({
            ...prev,
            [studentId]: value,
        }));

        // Обновление статуса посещаемости
        updateAttendance(studentId, value);
    };

    const updateAttendance = async (studentId, status) => {
        try {
            const response = await fetch("/api/teacher/studentAttendance", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
                body: JSON.stringify({ studentId, status }),
            });

            if (!response.ok) {
                throw new Error(`Ошибка обновления: ${response.status}`);
            }

            console.log("Статус успешно обновлен");
        } catch (error) {
            console.error("Ошибка при обновлении статуса", error);
        }
    };

    return (
        <div className="studentAttendance">
            <div className="studentAttendance__container" id="studentList">
                {errorMessage && <div className="error-message">{errorMessage}</div>}

                {users.map((user) => (
                    <StudentAttendanceCard
                        key={user.id}
                        id={user.id}
                        surname={user.last_name}
                        name={user.first_name}
                        patronymic={user.middle_name}
                        role={user.role}
                        selectedValue={attendance[user.id] || ""}
                        onChange={(value) => handleAttendanceChange(user.id, value)}
                    />
                ))}
            </div>
        </div>
    );
};

export default StudentAttendance;
