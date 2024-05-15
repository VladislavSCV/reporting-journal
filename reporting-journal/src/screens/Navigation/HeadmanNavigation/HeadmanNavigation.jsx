import React from "react";
import { useState } from "react";
import { Outlet, Link } from "react-router-dom";
import "./headmanNavigation.scss";
import row from "./../../../assets/Navigation/left.svg";
import row2 from "./../../../assets/Navigation/right.svg";
import group from "./../../../assets/Navigation/groups.svg";
import curatorGroup from "./../../../assets/Navigation/curator_group.svg";
import calendar from "./../../../assets/Navigation/Schedule.svg";
import menu from "./../../../assets/Navigation/menu.svg";
import menu2 from "./../../../assets/Navigation/menu2.svg";
import tables from "./../../../assets/Navigation/tables.svg";
import schedule2 from "./../../../assets/Navigation/schedule1.svg";
import notes from "./../../../assets/Navigation/notes.svg";
import studentsLits from "./../../../assets/Navigation/studentsList.svg";
const HeadmanNavigation = () => {
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
              src={isActive ? row2 : row}
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

export default HeadmanNavigation;
