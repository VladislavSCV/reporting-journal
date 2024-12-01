import React from "react";
import settingsUser from "./../../assets/StudentCard/settings.svg";
import deleteUserSVG from "./../../assets/StudentCard/delete.svg";
import "./studentCard.scss";

const StudentCard = ({ id, first_name, middle_name, last_name, role, status, onDelete }) => {
    const userRole = localStorage.getItem("userRole");
    const adminPrivileges = userRole === "Admin";

    return (
        <div className="studentCard">
            <div className="studentCard__color-block"></div>
            <div className="studentCard__container">
                <div className="studentCard__student">
                    <p className="studentCard__student-name">{first_name}</p>
                    {middle_name && (
                        <p className="studentCard__student-name studentCard__student-name--middle">
                            {middle_name}
                        </p>
                    )}
                    {last_name && (
                        <p className="studentCard__student-name studentCard__student-name--last">
                            {last_name}
                        </p>
                    )}
                    {status && (
                        <p className="studentCard__student-status">
                            Статус: {status}
                        </p>
                    )}
                    <p className="studentCard__student-role">Роль: {role}</p>
                </div>
                <div className="studentCard__buttons">
                    {adminPrivileges && (
                        <>
                            <img
                                src={settingsUser}
                                alt="Настройки"
                                className="studentCard__buttons-settings"
                                data-modal="ModalStudentSettings"
                                data-id={id}
                            />
                            <img
                                src={deleteUserSVG}
                                alt="Удалить"
                                onClick={() => onDelete(id)}
                                className="studentCard__buttons-delete"
                            />
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default StudentCard;
