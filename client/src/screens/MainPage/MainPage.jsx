import React from "react";
import { Link } from "react-router-dom";
import "./mainPage.scss";

const Main = () => {
    return (
        <div className="main-page">
            <header className="main-header">
                <h1 className="main-title">Reporting Journal</h1>
                <p className="main-subtitle">Выберите раздел для работы:</p>
            </header>
            <div className="main-links">
                <Link to="/groups" className="main-link">Группы</Link>
                <Link to="/curatorgroups" className="main-link">Группы куратора</Link>
                <Link to="/studentsList" className="main-link">Список студентов</Link>
                <Link to="/schedule" className="main-link">Расписание</Link>
                <Link to="/studentAttendance" className="main-link">Посещаемость студентов</Link>
                <Link to="/notes" className="main-link">Заметки</Link>
                <Link to="/GroupsNotes" className="main-link">Заметки групп</Link>
                <Link to="/GroupsSchedule" className="main-link">Расписание групп</Link>
                <Link to="/GroupsStudentsList" className="main-link">Список студентов групп</Link>
                <Link to="/CuratorGroupsNotes" className="main-link">Заметки групп куратора</Link>
                <Link to="/CuratorGroupsSchedule" className="main-link">Расписание групп куратора</Link>
                <Link to="/CuratorGroupsStudentsList" className="main-link">Список студентов групп куратора</Link>
                <Link to="/AdminPanel" className="main-link admin-link">Панель администратора</Link>
            </div>
        </div>
    );
};

export default Main;
