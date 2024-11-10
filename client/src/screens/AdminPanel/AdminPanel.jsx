import React, { useEffect, useState } from "react";
import "./adminPanel.scss";
import axios from "axios";
import UserCard from "../../components/UserCard/UserCard";
import del from "../../assets/AdminPanel/delete.svg";
const AdminPanel = () => {
  const [users, setUsers] = useState([]);
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const [roleList, setRoleList] = useState([]);
  const [value, setValue] = useState("");

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/user");

        setUsers(response.data);
      } catch (error) {
        console.error(error);
      }
    };

    const fetchRoles = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/auth/role");
        console.log(response.data);
        setRoleList(response.data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchUsers();
    fetchRoles();
  }, []);

  const addUser = async () => {
    try {
      const response = await axios.post(
        "http://localhost:5001/api/auth/registration",
        {
          login,
          password,
          role,
        }
      );
      setUsers([...users, response.data]);
    } catch (error) {
      console.error(error);
    }
  };

  const addRole = async () => {
    try {
      const response = await axios.post("http://localhost:5001/api/auth/role", {
        value,
      });
      setRoleList([...roleList, response.data]);
    } catch (error) {
      console.error(error);
    }
  };

  const deleteRole = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/auth/role/${key}`);
    } catch {
      console.error(error);
    }
  };

  return (
    <div className="adminPanel">
      <div className="adminPanel__container">
        <h1 className="adminPanel__title">Панель управления</h1>
        <div className="adminPanel__usersControl">
          <form action="" className="adminPanel__usersControl-form">
            <div className="">
              <label htmlFor="" className="adminPanel__usersControl-label">
                Логин пользователя:
              </label>

              <input
                onChange={(e) => setLogin(e.target.value)}
                type="text"
                className="adminPanel__usersControl-input"
              />
            </div>
            <div>
              <label htmlFor="" className="adminPanel__usersControl-label">
                Пароль пользователя:
              </label>
              <input
                onChange={(e) => setPassword(e.target.value)}
                type="text"
                className="adminPanel__usersControl-input"
              />
            </div>
            <div>
              <label htmlFor="" className="adminPanel__usersControl-label">
                Роль пользователя:
              </label>
              <select
                type="text"
                className="adminPanel__usersControl-input"
                onChange={(e) => setRole(e.target.value)}
              >
                {roleList.map((role) => {
                  return (
                    <option value={role.value} key={role.id} id={role.id}>
                      {role.value}
                    </option>
                  );
                })}
              </select>
            </div>
            <button
              className="adminPanel__usersControl-button"
              onClick={addUser}
            >
              Добавить
            </button>
          </form>

          <div className="adminPanel__usersControl-users">
            <div className="adminPanel__usersControl-users-titles">
              <p className="adminPanel__usersControl-users-titles-login">
                Логин
              </p>
              <p className="adminPanel__usersControl-users-titles-password">
                Пароль
              </p>
              <p className="adminPanel__usersControl-users-titles-role">
                Права
              </p>
              <p className="adminPanel__usersControl-users-titles-controls">
                Действия
              </p>
            </div>
            {users.map((user) => {
              return (
                <UserCard
                  name={user.login}
                  key={user.id}
                  id={user.id}
                  role={user.role}
                  password={user.password}
                />
              );
            })}
          </div>
        </div>

        <div className="adminPanel__roleControl">
          <form action="" className="adminPanel__roleControl-form">
            <label htmlFor="" className="adminPanel__roleControl-label">
              Новая роль:
            </label>
            <input
              type="text"
              onChange={(e) => setValue(e.target.value)}
              className="adminPanel__roleControl-input"
            />
            <button
              className="adminPanel__roleControl-button"
              onClick={addRole}
            >
              Добавить
            </button>
          </form>

          <div className="adminPanel__roleControl-roles">
            <div className="adminPanel__roleControl-roles-titles">
              <p className="adminPanel__roleControl-roles-titles-title">
                Название роли
              </p>
            </div>
            {roleList.map((role) => {
              return (
                <p
                  className="adminPanel__roleControl-roles-title"
                  key={role.id}
                  id={role.id}
                >
                  {role.value}

                  <img
                    src={del}
                    alt=""
                    onClick={() => deleteRole(role.id)}
                    className="adminPanel__roleControl-roles-title-delete"
                  />
                </p>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
};

export default AdminPanel;
