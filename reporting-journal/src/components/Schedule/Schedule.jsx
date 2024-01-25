import React from "react";
import "./schedule.scss";
import add from './../../img/GroupCard/Add.svg'
import UserNav from "../UserNav/UserNav";
const Schedule = () => {
  return (
    <div className="schedule">
      <UserNav />
      <div className="schedule__container">
        <div className="schedule__day">
          <h1 className="schedule__day-title">Понедельник</h1>
          <div className="schedule__add">
            <div className="schedule__add-container" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Вторник</h1>
          <div className="schedule__add">
            <div className="schedule__add-container" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Среда</h1>
          <div className="schedule__add">
            <div className="schedule__add-container" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Четверг</h1>
          <div className="schedule__add">
            <div className="schedule__add-container" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
        <div className="schedule__day">
          <h1 className="schedule__day-title">Пятница</h1>
          <div className="schedule__add">
            <div className="schedule__add-container" data-modal="ModalScheduleAdd">
              <img src={add} alt="" className="schedule__add-img" />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Schedule;
