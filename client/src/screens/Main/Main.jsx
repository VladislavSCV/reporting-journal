import React, { useState } from "react";
import welcome from "./../../assets/Welcome.svg";
import "./main.scss";
import { useDispatch, useSelector } from "react-redux";
import { loginUser } from "../../actions/api.js";
import { useNavigate } from "react-router-dom";

const Main = () => {
  const [login, setUserLogin] = useState("");
  const [hash, setPassword] = useState("");
  const [error, setError] = useState(null); // Для отображения ошибки

  const dispatch = useDispatch();
  const navigate = useNavigate();
  const isAuth = useSelector((state) => state.user.isAuth);

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      await loginUser(login, hash); // Ожидание завершения loginUser
      navigate("/mainPage"); // Перенаправление после успешного входа
    } catch (err) {
      console.error("Ошибка входа:", err);
      setError("Неверный логин или пароль"); // Обновляем состояние ошибки
    }
  };

  return (
      <div className="main">
        <div className="main__container">
          <div className="main__aboutBlock">
            <div className="main__about">
              <h1 className="main__title">Что такое Reporting Journal?</h1>
              <p className="main__text">
                <span className="main__text-bold"> Reporting Journal</span> - это
                web-приложение, которое упростит работу преподавателя и куратора,
                позволив в одном месте прикреплять для себя темы занятий и
                задания, а также отмечать отсутствующих с дальнейшим экспортом в
                Exel-файл
              </p>
            </div>
            <img src={welcome} alt="Welcome" />
          </div>

          {!isAuth && (
              <div className="main__loginBlock">
                <div className="main__loginBlock-container">
                  <h1 className="main__loginBlock-title">Вход в аккаунт</h1>
                  <form onSubmit={handleLogin}>
                    <label htmlFor="login" className="main__loginBlock-label">
                      Логин:
                    </label>
                    <input
                        required
                        value={login}
                        onChange={(e) => setUserLogin(e.target.value)}
                        type="text"
                        className="main__loginBlock-input"
                        id="login"
                    />
                    <label htmlFor="password" className="main__loginBlock-label">
                      Пароль:
                    </label>
                    <input
                        required
                        value={hash}
                        onChange={(e) => setPassword(e.target.value)}
                        type="password"
                        className="main__loginBlock-input"
                        id="password"
                    />
                    <button type="submit" className="main__loginBlock-button">
                      Войти
                    </button>
                  </form>
                  {error && <p className="error-message">{error}</p>}
                  <p className="main__loginBlock-access" data-modal="modalMainInfo">
                    Как получить доступ?
                  </p>
                </div>
              </div>
          )}
        </div>
      </div>
  );
};

export default Main;
