import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "../../screens/Groups/groups.scss";
import GroupCard from "../GroupCard/GroupCard.jsx";

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
        console.log(groups)

        fetchGroups();
    }, []);

    const handleGroupClick = (groupId) => {
        navigate(`/studentAttendance/${groupId}`);
    };

    return (
        <div className="groups">
            <h1 className="groups__title">Список групп</h1>
            <div className="groups__list">
                {groups.map((group) => (
                    // <div
                    //     key={group.id}
                    //     className="groupCard"
                    //     onClick={() => handleGroupClick(group.id)}
                    // >
                    //     <h3>{group.name}</h3>
                    //     <button>Просмотр</button>
                    // </div>
                    <GroupCard
                        key={group.id}
                        name={group.name}
                        id={group.id}
                        link={`/studentAttendance/${group.id}`}
                    />
                ))}
            </div>
        </div>
    );
};

export default GroupList;
