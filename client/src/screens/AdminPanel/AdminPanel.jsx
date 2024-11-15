import React, { useEffect, useState } from "react";
import "./adminPanel.scss";
import UserCard from "../../components/UserCard/UserCard";
import del from "../../assets/AdminPanel/delete.svg";

const AdminPanel = () => {
  const [users, setUsers] = useState([]);
  const [first_name, setFirstName] = useState("");
  const [middle_name, setMiddleName] = useState("");
  const [last_name, setLastName] = useState("");
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const [roleList, setRoleList] = useState([]);
  const [value, setValue] = useState("");

  // Fetch users
  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await fetch("https://reporting-journal-2.onrender.com/api/user");
        if (!response.ok) throw new Error(`Ошибка запроса: ${response.status}`);
        const data = await response.json();
        setUsers(data.users);
      } catch (error) {
        console.error("Ошибка загрузки пользователей:", error);
      }
    };

    const fetchRoles = async () => {
      try {
        const response = await fetch("https://reporting-journal-2.onrender.com/api/role");
        if (!response.ok) throw new Error(`Ошибка запроса: ${response.status}`);
        const data = await response.json();
        setRoleList(data.roles);
      } catch (error) {
        console.error("Ошибка загрузки ролей:", error);
      }
    };

    fetchUsers();
    fetchRoles();
  }, []);


  // Add user
  const addUser = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch("https://reporting-journal-2.onrender.com/api/auth/registration", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ first_name, middle_name, last_name, login, password, role: 1}),
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      const data = await response.json();
      setUsers([...users, data]);
    } catch (error) {
      console.error(error);
    }
  };

  // Add role
  const addRole = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch("https://reporting-journal-2.onrender.com/api/auth/role", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ value }),
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      const data = await response.json();
      setRoleList([...roleList, data]);
    } catch (error) {
      console.error(error);
    }
  };

  // Delete role
  const deleteRole = async (key) => {
    try {
      const response = await fetch(`https://reporting-journal-2.onrender.com/api/auth/role/${key}`, {
        method: "DELETE",
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      setRoleList(roleList.filter((role) => role.id !== key));
    } catch (error) {
      console.error(error);
    }
  };

  return (
      <div className="adminPanel">
        <div className="adminPanel__container">
          <h1 className="adminPanel__title">Панель управления</h1>
          <div className="adminPanel__usersControl">
            <form className="adminPanel__usersControl-form" onSubmit={addUser}>
              <div className="">
              <label htmlFor="" className="adminPanel__usersControl-label">
                Фамилия:
              </label>
              <input
                  onChange={(e) => setFirstName(e.target.value)}
                  type="text"
                  className="adminPanel__usersControl-input"
              />
          </div>
          <div className="">
            <label htmlFor="" className="adminPanel__usersControl-label">
              Имя:
            </label>
            <input
                onChange={(e) => setMiddleName(e.target.value)}
                type="text"
                className="adminPanel__usersControl-input"
            />
          </div>
          <div className="">
            <label htmlFor="" className="adminPanel__usersControl-label">
              Отчество:
            </label>
            <input
                onChange={(e) => setLastName(e.target.value)}
                type="text"
                className="adminPanel__usersControl-input"
            />
            </div>

              <div>
                <label className="adminPanel__usersControl-label">Логин пользователя:</label>
                <input
                    onChange={(e) => setLogin(e.target.value)}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div>
                <label className="adminPanel__usersControl-label">Пароль пользователя:</label>
                <input
                    onChange={(e) => setPassword(e.target.value)}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div>
                <label className="adminPanel__usersControl-label">Роль пользователя:</label>
                <select
                    className="adminPanel__usersControl-input"
                    onChange={(e) => setRole(e.target.value)}
                >
                  {roleList.map((role) => (
                      <option value={role.value} key={role.id}>
                        {role.value}
                      </option>
                  ))}
                </select>
              </div>
              <button type="submit" className="adminPanel__usersControl-button">
                Добавить
              </button>
            </form>
            <div className="adminPanel__usersControl-users">
              {users.map((user) => (
                  <UserCard
                      key={user.id}
                      name={user.login}
                      id={user.id}
                      role={user.role}
                      password={user.password}
                  />
              ))}
            </div>
          </div>
          <div className="adminPanel__roleControl">
            <form className="adminPanel__roleControl-form" onSubmit={addRole}>
              <label className="adminPanel__roleControl-label">Новая роль:</label>
              <input
                  type="text"
                  onChange={(e) => setValue(e.target.value)}
                  className="adminPanel__roleControl-input"
              />
              <button type="submit" className="adminPanel__roleControl-button">
                Добавить
              </button>
            </form>
            <div className="adminPanel__roleControl-roles">
              {roleList.map((role) => (
                  <p key={role.id} className="adminPanel__roleControl-roles-title">
                    {role.value}
                    <img
                        src={del}
                        alt="delete"
                        onClick={() => deleteRole(role.id)}
                        className="adminPanel__roleControl-roles-title-delete"
                    />
                  </p>
              ))}
            </div>
          </div>
        </div>
      </div>
  );
};

export default AdminPanel;
