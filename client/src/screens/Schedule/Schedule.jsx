import React from "react";
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
const Schedule = () => {
  return (
    <div className="schedule">
      <div className="schedule__container">
        <div className="schedule__date"></div>

        <div className="schedule__day">
          <h1 className="schedule__day-title">Понедельник</h1>
          <div className="schedule__day-container">
            {monday.map((obj, index) => {
              return (
                <LessonCard
                  lesson={obj.lesson}
                  teacher={obj.teacher}
                  key={index}
                />
              );
            })}
            <div className="schedule__add" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Вторник</h1>
          <div className="schedule__day-container">
            {tuesday.map((obj, index) => {
              return (
                <LessonCard
                  lesson={obj.lesson}
                  teacher={obj.teacher}
                  key={index}
                />
              );
            })}
            <div className="schedule__add" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Среда</h1>
          <div className="schedule__day-container">
            {wednesday.map((obj, index) => {
              return (
                <LessonCard
                  lesson={obj.lesson}
                  teacher={obj.teacher}
                  key={index}
                />
              );
            })}
            <div className="schedule__add" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Четверг</h1>
          <div className="schedule__day-container">
            {thursday.map((obj, index) => {
              return (
                <LessonCard
                  lesson={obj.lesson}
                  teacher={obj.teacher}
                  key={index}
                />
              );
            })}
            <div className="schedule__add" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Пятница</h1>
          <div className="schedule__day-container">
            {friday.map((obj, index) => {
              return (
                <LessonCard
                  lesson={obj.lesson}
                  teacher={obj.teacher}
                  key={index}
                />
              );
            })}
            <div className="schedule__add" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__swipe">
          <div className="schedule__swipe-block">
            <img src={left} alt="" className="schedule__swipe-block-button" />
            <p className="schedule__swipe-block-button-text">Прошлая неделя</p>
          </div>
          <div className="schedule__swipe-block">
            <p className="schedule__swipe-block-button-text">
              Следующая неделя
            </p>
            <img src={right} alt="" className="schedule__swipe-block-button" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Schedule;
