import React, { useEffect, useState } from "react";
import "./groups.scss";
import GroupCard from "../../components/GroupCard/GroupCard"; // Компонент для отображения групп
import { getGroups } from "../../actions/api"; // Функция для получения групп

const Groups = () => {
    const [groups, setGroups] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        getGroups().then((data) => {
            setGroups(data.groups);
        }).catch((error) => {
            setError("Ошибка загрузки групп");
        });
    }, []);

    return (
        <div className="groups">
            <div className="groups__container">
                <div className="groups__list">
                    {error && <p>{error}</p>}
                    {groups.map((obj) => (
                        <GroupCard
                            key={obj.id}
                            name={obj.name}
                            id={obj.id}
                            link={`/studentsList?id=${obj.id}`}
                            // onClick={() => localStorage.setItem("group_id", obj.id)} // Сохранение group_id при клике
                        />
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Groups;
