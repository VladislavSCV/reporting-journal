import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./groupList.scss";

const GroupList = () => {
    const navigate = useNavigate();
    const [groups, setGroups] = useState([]);

    useEffect(() => {
        const fetchGroups = async () => {
            try {
                const response = await fetch("/api/teacher/groups", {
                    method: "GET",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                });

                if (!response.ok) {
                    throw new Error(`Ошибка запроса: ${response.status}`);
                }

                const responseData = await response.json();

                if (responseData.groups) {
                    try {
                        const decodedData = JSON.parse(atob(responseData.groups));
                        setGroups(decodedData || []);
                    } catch (error) {
                        console.error("Ошибка декодирования Base64:", error);
                        setGroups([]);
                    }
                } else {
                    setGroups([]);
                }
            } catch (error) {
                console.error("Ошибка получения данных:", error);
            }
        };

        fetchGroups();
    }, []);

    const handleGroupClick = (groupId) => {
        navigate(`/studentAttendance/${groupId}`);
    };

    return (
        <div className="container">
            <h1 className="title">Список групп</h1>
            <div className="cardContainer">
                {groups.map((group) => (
                    <div
                        key={group.id}
                        className="groupCard"
                        onClick={() => handleGroupClick(group.id)}
                    >
                        <h3>{group.name}</h3>
                        <button>Просмотр</button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default GroupList;
