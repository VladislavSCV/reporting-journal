import React from "react";
import { Outlet, Link } from "react-router-dom";
import "./navigation.scss";
import logo from "./../../img/nav/Logo.svg";
import group from "./../../img/nav/group.svg";
import calendar from "./../../img/nav/calnedar.svg";
import check from "./../../img/nav/check.svg";
const Navigation = () => {
  return (
    <aside className="navigation">
      <div className="navigation__container">
        <img src={logo} alt="" className="navigation__logo" />
        <nav className="navigation__lists">
          <ul className="navigation__list">
            <p className="navigation__list-title">Меню</p>
            <li>
              <Link to="/" className="navigation__list-element">
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
          <ul className="navigation__list">
            <p className="navigation__list-title">Кураторство</p>
            <li >
            <Link to="/curatorgroups/*" className="navigation__list-element">
              <img src={group} alt="" /> Ваши группы
              </Link>
            </li>
            <li >
            <Link to="/schedule" className="navigation__list-element">
              <img src={check} alt="" /> Посещаемость
              </Link>
            </li>
          </ul>
        </nav>
      </div>
    </aside>
  );
};

export default Navigation;
