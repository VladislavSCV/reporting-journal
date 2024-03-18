import React from "react";
import { useState } from "react";
import { Outlet, Link } from "react-router-dom";
import "./navigation.scss";
import row from "./../../assets/Navigation/left.svg";
import row2 from "./../../assets/Navigation/right.svg";
import group from "./../../assets/Navigation/groups.svg";
import curatorGroup from "./../../assets/Navigation/curator_group.svg";
import calendar from "./../../assets/Navigation/Schedule.svg";
import menu from "./../../assets/Navigation/menu.svg";
import menu2 from "./../../assets/Navigation/menu2.svg";
import tables from "./../../assets/Navigation/tables.svg";
import schedule2 from "./../../assets/Navigation/schedule1.svg";
import notes from "./../../assets/Navigation/notes.svg";
import studentsLits from "./../../assets/Navigation/studentsList.svg";
const Navigation = () => {
  const [isActive, setIsActive] = useState(false);

  const inv = () => {
    setIsActive(!isActive);
  };

  const [isActiveBurger, setIsActiveBurger] = useState(false);

  const burger = () => {
    setIsActiveBurger(!isActiveBurger);
  };

  return (
    <div className={isActiveBurger ? " navigation_on" : ""}>
      <img src={menu} alt="" className="navigation__burger" onClick={burger} />
      <aside className={`navigation${isActive ? " invisible" : ""}`}>
        <div className="navigation__container">
          <nav className="navigation__lists">
            {/* <Link to="/main">
              <div className="navigation__logoBlock">
                <img src={logo} alt="" className="navigation__logo" />
                <img src={logo_text} alt="" className="navigation__logo-text" />
              </div>
            </Link> */}
            <div className="navigation__user">
              <div className="navigation__user-avatar"></div>
              <div className="navigation__user-container">
                <p className="navigation__user-name">Гилоян Роман</p>
                <p className="navigation__user-role">Преподаватель</p>
              </div>
            </div>
            <img
              src={isActive ? row : row2}
              onClick={inv}
              className="navigation__collapse"
            />

            <ul className="navigation__list">
              <li>
                <Link to="/groups" className="navigation__list-element">
                  <img src={group} alt="" />
                  <p className="navigation__list-element-text">Группы</p>
                </Link>
              </li>
              <li>
                <Link to="/schedule" className="navigation__list-element">
                  <img src={calendar} alt="" />
                  <p className="navigation__list-element-text">Расписание</p>
                </Link>
              </li>
            </ul>
            <p className="navigation__list-title">Преподавателю</p>
            <ul className="navigation__list">
              <li>
                <Link
                  to="/groupsStudentsList"
                  className="navigation__list-element"
                >
                  <img src={studentsLits} alt="" />
                  <p className="navigation__list-element-text">
                    Список студентов
                  </p>
                </Link>
              </li>
              <li>
                <Link to="/GroupsSchedule" className="navigation__list-element">
                  <img src={schedule2} alt="" />
                  <p className="navigation__list-element-text">Расписание</p>
                </Link>
              </li>
              <li>
                <Link to="/GroupsNotes" className="navigation__list-element">
                  <img src={notes} alt="" />
                  <p className="navigation__list-element-text">Вложения</p>
                </Link>
              </li>
            </ul>
            <p className="navigation__list-title">Куратору</p>
            <ul className="navigation__list">
              <li>
                <Link
                  to="/curatorgroups/*"
                  className="navigation__list-element"
                >
                  <img src={curatorGroup} alt="" />
                  <p className="navigation__list-element-text">Ваши группы</p>
                </Link>
              </li>

              <li>
                <Link
                  to="/CuratorGroupsNotes"
                  className="navigation__list-element"
                >
                  <img src={studentsLits} alt="" />
                  <p className="navigation__list-element-text">
                    Список студентов
                  </p>
                </Link>
              </li>
              <li>
                <Link
                  to="/CuratorGroupsSchedule"
                  className="navigation__list-element"
                >
                  <img src={schedule2} alt="" />
                  <p className="navigation__list-element-text">Расписание</p>
                </Link>
              </li>
              <li>
                <Link
                  to="/CuratorGroupsStudentsList"
                  className="navigation__list-element"
                >
                  <img src={notes} alt="" />
                  <p className="navigation__list-element-text">Вложения</p>
                </Link>
              </li>
              <li>
                <Link to="#" className="navigation__list-element">
                  <img src={tables} alt="" />
                  <p className="navigation__list-element-text">
                    Экспорт таблиц
                  </p>
                </Link>
              </li>
            </ul>
          </nav>
          {/* <div className="navigation__user">
            <div className="navigation__user-avatar"></div>
            <div className="navigation__user-container">
              <p className="navigation__user-name">Гилоян Роман</p>
              <p className="navigation__user-role">Преподаватель</p>
            </div>
          </div> */}
        </div>
      </aside>
    </div>
  );
};

export default Navigation;
