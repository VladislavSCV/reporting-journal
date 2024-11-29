import React, { useState, useEffect } from "react";
import "./curatorGroupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const CuratorGroupsNotes = () => {
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
      if (!response.ok) throw new Error(`Ошибка запроса: ${response.status}`);
      const database64 = await response.json().groups;
      if (!database64) {
        console.warn("С сервера получен null. Устанавливаем пустой массив.");
        setGroups([]);
        return;
      }
      // Проверяем, требуется ли декодирование Base64
      let data;
      if (typeof database64 === "string") {
        try {
          data = JSON.parse(atob(database64)); // Декодируем Base64, если это строка
        } catch (error) {
          console.error("Ошибка декодирования Base64. Возможно, данные уже JSON.");
          setGroups([]);
          return;
        }
      } else {
        data = database64; // Если это объект, декодирование не требуется
      }
      setGroups(data.groups || []);
    } catch (error) {
      console.error(error);
    }
  };

  fetchGroups();
}, []);
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        <h1 className="groups__title">Вложения группы:</h1>
        <div className="groups__list">
          {groups.length > 0 ? (
              groups.map((obj) => (
              <GroupCardLinks
                group={obj.name}
                key={obj.id}
                id={obj.id}
                link={"/notes?id=" + obj.id}
              />
            ))
          ) : (
            <p>Группы не найдены</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default CuratorGroupsNotes;
