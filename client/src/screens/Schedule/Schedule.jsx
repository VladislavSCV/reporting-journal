import React, { useEffect, useState } from "react";
import "./schedule.scss";
import LessonCard from "../../components/LessonCard/LessonCard";
import add from "./../../assets/GroupCard/Add.svg";
import axios from "axios";

const Schedule = () => {
  const [schedule, setSchedule] = useState([]);

  useEffect(() => {
    const groupId = localStorage.getItem("group_id");
    if (!groupId) {
      console.error("Group ID is not found in localStorage");
      return;
    }

    const fetchGroups = async () => {
      try {
        const response = await axios.get(`/api/schedule/${groupId}`);
        setSchedule(response.data.schedule || []);
      } catch (error) {
        console.error("Error fetching schedule:", error);
      }
    };

    fetchGroups();
  }, []);

  console.log(schedule);

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
                (lesson) => lesson.DayOfWeek === day.key
            );

            return (
                <div className="schedule__day" key={day.key}>
                  <h1 className="schedule__day-title">{day.name}</h1>
                  <div className="schedule__day-container">
                    {daySchedule.length > 0 ? (
                        daySchedule.map((obj) => (
                            <LessonCard
                                key={obj.ScheduleID}
                                id={obj.ScheduleID}
                                lesson={obj.SubjectName}
                                teacher={obj.TeacherName}
                            />
                        ))
                    ) : (
                        <p>Нет занятий на этот день</p>
                    )}
                    {/*<div*/}
                    {/*    className="schedule__add"*/}
                    {/*    data-modal="ModalScheduleAdd"*/}
                    {/*    data-day={day.key}*/}
                    {/*>*/}
                      {/*<img src={add} alt="Add" className="schedule__add-img" />*/}
                    {/*</div>*/}
                  </div>
                </div>
            );
          })}
        </div>
      </div>
  );
};

export default Schedule;
