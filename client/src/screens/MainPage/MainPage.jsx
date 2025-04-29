import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { LogOut, User, Shield, BookOpen, Calendar } from "lucide-react";
import "./mainPage.scss";

const roleLinks = {
    admin: [
        { path: "/groups", label: "Группы", icon: <Shield /> },
        { path: "/studentsList", label: "Список студентов", icon: <User /> },
        { path: "/schedule", label: "Расписание", icon: <Calendar /> },
        { path: "/groupsList", label: "Посещаемость студентов", icon: <BookOpen /> },
        { path: "/notes", label: "Заметки", icon: <BookOpen /> },
        { path: "/GroupsNotes", label: "Заметки групп", icon: <BookOpen /> },
        { path: "/AdminPanel", label: "Панель администратора", icon: <Shield /> }
    ],
    teacher: [
        { path: "/CuratorGroupsSchedule", label: "Расписание групп куратора", icon: <Calendar /> },
        { path: "/CuratorGroupsStudentsList", label: "Список студентов групп куратора", icon: <User /> }
    ],
    student: [
        { path: "/groups", label: "Группы", icon: <Shield /> },
        { path: "/studentsList", label: "Список студентов", icon: <User /> },
        { path: "/notes", label: "Заметки", icon: <BookOpen /> },
        { path: "/schedule", label: "Расписание", icon: <Calendar /> }
    ]
};

const Main = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [userRole, setUserRole] = useState('');

    const handleLogout = () => {
        localStorage.clear();
        navigate("/main");
    };

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
                    localStorage.setItem("username", fullName);
                    localStorage.setItem("userRole", data.user.role);
                    console.log(fullName);
                } catch (error) {
                    console.error('Ошибка при получении пользователя:', error);
                }
            };
            fetchUser();
        }
    }, []);

    const links = roleLinks[userRole] || [];

    return (
        <div className="main-page">
            <header className="main-header">
                <h1 className="main-title">📘 Reporting Journal</h1>
                <p className="main-subtitle">Привет, <strong>{username}</strong>!</p>
                <div className="user-info">
                    <span className="user-role">{userRole}</span>
                    <button className="logout-btn" onClick={handleLogout}><LogOut size={18} /> Выйти</button>
                </div>
            </header>

            <div className="card-grid">
                {links.map(({ path, label, icon }) => (
                    <Link key={path} to={path} className="nav-card">
                        <div className="icon">{icon}</div>
                        <span className="label">{label}</span>
                    </Link>
                ))}
            </div>
        </div>
    );
};

export default Main;
