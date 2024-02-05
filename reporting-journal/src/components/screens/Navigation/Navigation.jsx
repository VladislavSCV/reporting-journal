import React from "react";
import { Outlet, Link } from "react-router-dom";
import "./navigation.scss";
import logo from "./../../../img/nav/Logo.svg";
import group from "./../../../img/nav/group.svg";
import calendar from "./../../../img/nav/calnedar.svg";

const Navigation = () => {
  return (
    <aside className="navigation">
      <div className="navigation__container">
        <nav className="navigation__lists">
          <Link to="/mainPage">
            <img src={logo} alt="" className="navigation__logo" />
          </Link>
            <p className="navigation__list-title">Меню</p>
          <ul className="navigation__list">
            <li>
              <Link to="/groups" className="navigation__list-element">
                <img src={group} alt="" /> Группы
              </Link>
            </li>
            <li>
              <Link to="/schedule" className="navigation__list-element">
                <img src={calendar} alt="" />
                Расписание
              </Link>
            </li>
          </ul>
            <p className="navigation__list-title">Кураторство</p>
          <ul className="navigation__list">
            <li>
              <Link to="/curatorgroups/*" className="navigation__list-element">
                <img src={group} alt="" /> Ваши группы
              </Link>
            </li>
          </ul>
        </nav>
        <div className="navigation__user">
          <div className="navigation__user-avatar"></div>
          <div className="navigation__user-container">
            <div className="navigation__user-name">Гилоян Роман</div>
            <div className="navigation__user-role">Преподаватель</div>
          </div>
        </div>
      </div>
    </aside>
  );
};

export default Navigation;
