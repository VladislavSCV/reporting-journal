import React from "react";
import welcome from "./../../assets/Welcome.svg";
import "./mainPage.scss";
const MainPage = () => {
  return (
    <div className="mainPage">
      <div className="mainPage__container">
        <div className="mainPage__aboutBlock">
          <div className="mainPage__about">
            <h1 className="mainPage__title">Что такое Reporting Journal?</h1>
            <p className="mainPage__text">
              <span className="mainPage__text-bold">Reporting Journal</span> -
              это web-приложение, которое упростит работу преподавателя и
              куратора, позволив в одном месте прикреплять для себя темы занятий
              и задания, а также отмечать отсутствующих с дальнейшим экспортом в
              Exel-файл
            </p>
          </div>
          <img src={welcome} alt="" />
        </div>
        <div className="mainPage__loginBlock">
          <div className="mainPage__loginBlock-container">
            <h1 className="mainPage__loginBlock-title">Вход в аккаунт</h1>
            <label htmlFor="login" className="mainPage__loginBlock-label">
              Логин:
            </label>
            <input
              type="text"
              className="mainPage__loginBlock-input"
              id="login"
            />
            <label htmlFor="password" className="mainPage__loginBlock-label">
              Пароль:
            </label>
            <input
              type="password"
              className="mainPage__loginBlock-input"
              id="password"
            />
            <button className="mainPage__loginBlock-button">Войти</button>
            <p className="mainPage__loginBlock-access" data-modal="modalMainPageInfo">Как получить доступ?</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainPage;
