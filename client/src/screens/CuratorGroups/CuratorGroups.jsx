import React, { useEffect, useState } from "react";
import "./curatorGroups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import axios from "axios";

const CuratorGroups = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await fetch("/api/teacher/groups", {
          method: "GET",
          headers: { "Authorization": `Bearer ${localStorage.getItem('token')}` }
        });
        if (!response.ok) {
          throw new Error(`Ошибка запроса: ${response.status}`);
        }

        const database64 = await response.json(); // Получение JSON-ответа
        console.log("Raw database64:", database64);

        // if (!database64) {
        //   console.warn("С сервера получен null. Устанавливаем пустой массив.");
        //   setGroups([]);
        //   return;
        // }

        console.log(typeof database64)

        // Проверяем, требуется ли декодирование Base64
        let data;
        // if (typeof database64 === "string") {
          try {
            data = JSON.parse(atob(database64.groups)); // Декодируем Base64, если это строка
          } catch (error) {
            console.error("Ошибка декодирования Base64. Возможно, данные уже JSON.");
            setGroups([]);
            return;
          }
        // } else {
        //   data = database64; // Если это объект, декодирование не требуется
        // }
        console.log(data)
        setGroups(data || []);
      } catch (error) {
        console.error(error);
        setGroups([]);
      }
    };

    fetchGroups();
  }, []);



  return (
      <div className="curatorGroups">
        <div className="curatorGroups__container">
          {groups.length === 0 ? (
              <div className="adminPanel__message-container" style={{backgroundColor: "#f0f0ff"}}>
                <div className="adminPanel__message">
                  <p>Доступ к этой панели есть только у куратора.</p>
                </div>
              </div>
          ) : (
              <div className="groups__list">
                {groups.map((obj) => (
                    <GroupCard
                        key={obj.id}
                        id={obj.id}
                        name={obj.name}
                        curator={obj.curator}
                        students={obj.students}
                        link={`/studentsList?id=${obj.id}`}
                    />
                ))}
              </div>
          )}
        </div>
      </div>
  );
};

export default CuratorGroups;