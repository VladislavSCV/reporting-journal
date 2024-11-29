import React, { useState, useEffect } from "react";
import "./curatorGroupsStudentsList.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";

const CuratorGroupsStudentsList = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await fetch("/api/teacher/groups", {
          method: "GET",
          headers: { "Authorization": `Bearer ${localStorage.getItem('token')}` },
        });

        if (!response.ok) {
          throw new Error(`Ошибка запроса: ${response.status}`);
        }

        const database64 = await response.json().groups; // Получение JSON-ответа
        console.log("Raw database64:", database64);

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

        console.log("Parsed data:", data);
        setGroups(data?.groups || []);
      } catch (error) {
        console.error("Ошибка при получении групп:", error);
        setGroups([]); // Устанавливаем пустой массив при ошибке
      }
    };

    fetchGroups();
  }, []); // Пустой массив зависимостей

  return (
      <div className="curatorGroups">
        <div className="curatorGroups__container">
          <h1 className="groups__title">Список студентов группы:</h1>
          <div className="groups__list">
            {groups.length > 0 ? (
                groups.map((obj) => (
                    <GroupCardLinks
                        group={obj.name}
                        key={obj.id}
                        id={obj.id}
                        link={`/studentsList?id=${obj.id}`}
                    />
                ))
            ) : (
                <p>Нет данных о группах.</p>
            )}
          </div>
        </div>
      </div>
  );
};

export default CuratorGroupsStudentsList;
