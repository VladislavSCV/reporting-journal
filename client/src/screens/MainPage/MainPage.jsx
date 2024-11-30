import React, { useEffect, useState } from "react";
import {Link, useNavigate} from "react-router-dom";
import "./mainPage.scss";

const Main = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [userRole, setUserRole] = useState('');

    const handleLogout = () => {
        localStorage.clear()
        navigate("/main")
    }

    useEffect(() => {
        const storedUsername = localStorage.getItem('username');
        const storedUserRole = localStorage.getItem('userRole');

        if (storedUsername && storedUserRole) {
            setUsername(storedUsername);
            setUserRole(storedUserRole);
        } else {
            const fetchUser = async () => {
                try {
                    const response = await fetch('/api/auth/', {
                        method: 'GET',
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': `Bearer ${localStorage.getItem('token')}`
                        }
                    });
                    const data = await response.json();
                    const fullName = `${data.user.first_name} ${data.user.last_name} ${data.user.middle_name}`;
                    setUsername(fullName);
                    setUserRole(data.user.role);

                    // Сохранение данных в localStorage
                    localStorage.setItem("username", fullName);
                    localStorage.setItem("userRole", data.user.role);
                } catch (error) {
                    console.error('Error fetching user:', error);
                }
            }
            fetchUser();
        }
    }, []);

    return (
        <div className="main-page">
            <header className="main-header">
                <h1 className="main-title">Reporting Journal</h1>
                <p className="main-subtitle">Выберите раздел для работы:</p>
                <div className="user-info">
                    <div className="user-details">
                        <h2 className="user-name">{username}</h2>
                        <p className="user-role">{userRole}</p>
                    </div>
                    <button className="logout-button" onClick={handleLogout}>Выйти</button>
                </div>
            </header>

            <div className="main-links">
                <Link to="/groups" className="main-link">Группы</Link>
                <Link to="/curatorgroups" className="main-link">Группы куратора</Link>
                <Link to="/studentsList" className="main-link">Список студентов</Link>
                <Link to="/schedule" className="main-link">Расписание</Link>
                <Link to="/groupsList" className="main-link">Посещаемость студентов</Link>
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
