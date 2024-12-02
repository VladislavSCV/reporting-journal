// ErrorNotification.jsx
import React from "react";
import './ErrorNotification.scss';  // Добавим стили для ошибки

const ErrorNotification = ({ message }) => {
    return (
        <div className="error-notification">
            <p>{message}</p>
        </div>
    );
};

export default ErrorNotification;
