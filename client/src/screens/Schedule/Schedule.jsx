import React, { useEffect, useState } from "react";
import "./schedule.scss";
import { monday } from "../../helpers/Schedule/monday";
import { tuesday } from "../../helpers/Schedule/tuesday";
import { wednesday } from "../../helpers/Schedule/wednesday";
import { thursday } from "../../helpers/Schedule/thursday";
import { friday } from "../../helpers/Schedule/friday";
import LessonCard from "../../components/LessonCard/LessonCard";
import left from "../../assets/Schedule/left.svg";
import right from "../../assets/Schedule/right.svg";
import add from "./../../assets/GroupCard/Add.svg";
import axios from "axios";
const Schedule = () => {
  const [schedule, setSchedule] = useState([]);
  let groupId = Number(window.location.search.substring(1).split("=")[1]);
  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/schedule");
        setSchedule(response.data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchGroups();
  }, []);

  const daysOfWeek = [
    { name: "Понедельник", key: "Monday" },
    { name: "Вторник", key: "Tuesday" },
    { name: "Среда", key: "Wednesday" },
    { name: "Четверг", key: "Thursday" },
    { name: "Пятница", key: "Friday" },
  ];
  return (
    <div className="schedule">
      <div className="schedule__container">
        {daysOfWeek.map((day) => {
          const daySchedule = schedule.filter(
            (lesson) =>
              lesson.dayOfWeek === day.key && lesson.groupId === groupId
          );
          return (
            <div className="schedule__day" key={day.key}>
              <h1 className="schedule__day-title">{day.name}</h1>
              <div className="schedule__day-container">
                {daySchedule.length > 0 ? (
                  daySchedule.map((obj) => (
                    <LessonCard
                      key={obj.id}
                      id={obj.id}
                      lesson={obj.subject}
                      teacher={obj.teacher}
                    />
                  ))
                ) : (
                  <p></p>
                )}
                <div
                  className="schedule__add"
                  data-modal="ModalScheduleAdd"
                  data-day={day.key}
                  data-id={groupId}
                >
                  <img src={add} alt="" className="schedule__add-img" />
                </div>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Schedule;
