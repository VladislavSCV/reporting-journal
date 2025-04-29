import React, { useEffect, useState } from "react";
import "./schedule.scss";
import LessonCard from "../../components/LessonCard/LessonCard";
import { useParams } from "react-router-dom";
import axios from "axios";

const Schedule = () => {
  const [schedule, setSchedule] = useState([]);
  const params = useParams(); // Получаем параметры из URL
  let groupId = ""

  useEffect(() => {
    // Определяем, что использовать: id из URL или из localStorage
      groupId = localStorage.getItem("group_id");
    console.log(groupId);
    if (!groupId) {
      console.error("Group ID is not found in URL or localStorage");
      return;
    }

    // Функция для получения расписания
    const fetchGroups = async () => {
      try {
        const response = await axios.get(`/api/schedule/${groupId}`, {
          headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
        });
        setSchedule(response.data.schedule || []);
        console.log("Received schedule:", response.data.schedule);
      } catch (error) {
        console.error("Error fetching schedule:", error);
      }
    };

    fetchGroups();
  }, [params.id]); // Выполняем эффект при изменении id в URL

  // Дни недели
  const daysOfWeek = [
    { name: "Понедельник", key: 1 },
    { name: "Вторник", key: 2 },
    { name: "Среда", key: 3 },
    { name: "Четверг", key: 4 },
    { name: "Пятница", key: 5 },
  ];

  return (
      <div className="schedule">
        <div className="schedule__container">
          {daysOfWeek.map((day) => {
            const daySchedule = schedule.filter(
                (lesson) => lesson.dayOfWeek === day.key
            );

            return (
                <div className="schedule__day" key={day.key}>
                  <h1 className="schedule__day-title">{day.name}</h1>
                  <div className="schedule__day-container">
                    {daySchedule.length > 0 ? (
                        daySchedule.map((obj) => (
                          <LessonCard
                          key={obj.scheduleID}
                          id={obj.scheduleID}
                          lesson={obj.subjectName}
                          teacher={obj.teacherName}
                        />
                        
                        ))
                    ) : (
                        <p>Нет занятий на этот день</p>
                    )}
                  </div>
                </div>
            );
          })}
        </div>
      </div>
  );
};

export default Schedule;
