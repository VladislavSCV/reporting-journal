import React, { useEffect, useState } from "react";
import "./adminPanel.scss";
import UserCard from "../../components/UserCard/UserCard";
import del from "../../assets/AdminPanel/delete.svg";

const AdminPanel = () => {
  const [users, setUsers] = useState([]);
  const [firstName, setFirstName] = useState("");
  const [middleName, setMiddleName] = useState("");
  const [lastName, setLastName] = useState("");
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const [roleList, setRoleList] = useState([]);
  const [value, setValue] = useState("");

  // Fetch users and roles
  useEffect(() => {
    const fetchData = async () => {
      try {
        const userResponse = (await fetch("https://reporting-journal-2.onrender.com/api/user", {
        }))
        if (!userResponse.ok) throw new Error(`Ошибка запроса: ${userResponse.status}`);
        const userData = await userResponse.json();
        setUsers(userData.users);

        const roleResponse = await fetch("https://reporting-journal-2.onrender.com/api/role");
        if (!roleResponse.ok) throw new Error(`Ошибка запроса: ${roleResponse.status}`);
        const roleData = await roleResponse.json();
        setRoleList(roleData.roles);
      } catch (error) {
        console.error("Ошибка загрузки данных:", error);
      }
    };
    fetchData();
  }, []);

  // Add user
  const addUser = async (e) => {
    e.preventDefault();
    if (!firstName || !middleName || !login || !password || !role) {
      alert("Пожалуйста, заполните все поля.");
      return;
    }

    try {
      const response = await fetch("https://reporting-journal-2.onrender.com/api/auth/registration", {
        method: "POST",
        mode: "no-cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ first_name: firstName, middle_name: middleName, last_name: lastName, login, password, role_id: 1 }),
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      const data = await response.json();
      setUsers([...users, data]);
      clearForm();
    } catch (error) {
      console.error("Ошибка при добавлении пользователя:", error);
    }
  };

  // Add role
  const addRole = async (e) => {
    e.preventDefault();
    if (!value) {
      alert("Пожалуйста, введите название роли.");
      return;
    }

    try {
      const response = await fetch("https://reporting-journal-2.onrender.com/api/auth/role", {
        method: "POST",
        mode: "no-cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ value }),
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      const data = await response.json();
      setRoleList([...roleList, data]);
      setValue("");  // Clear the input after adding
    } catch (error) {
      console.error("Ошибка при добавлении роли:", error);
    }
  };

  // Delete role
  const deleteRole = async (roleId) => {
    try {
      const response = await fetch(`https://reporting-journal-2.onrender.com/api/auth/role/${roleId}`, {
        method: "DELETE",
        mode: "no-cors"
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      setRoleList(roleList.filter((role) => role.id !== roleId));
    } catch (error) {
      console.error("Ошибка при удалении роли:", error);
    }
  };

  // Clear user form after adding
  const clearForm = () => {
    setFirstName("");
    setMiddleName("");
    setLastName("");
    setLogin("");
    setPassword("");
    setRole("");
  };

  return (
      <div className="adminPanel">
        <div className="adminPanel__container">
          <h1 className="adminPanel__title">Панель управления</h1>
          <div className="adminPanel__usersControl">
            <form className="adminPanel__usersControl-form" onSubmit={addUser}>
              <div className="">
                <label className="adminPanel__usersControl-label">Фамилия:</label>
                <input
                    onChange={(e) => setFirstName(e.target.value)}
                    value={firstName}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div className="">
                <label className="adminPanel__usersControl-label">Имя:</label>
                <input
                    onChange={(e) => setMiddleName(e.target.value)}
                    value={middleName}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div className="">
                <label className="adminPanel__usersControl-label">Отчество:</label>
                <input
                    onChange={(e) => setLastName(e.target.value)}
                    value={lastName}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div>
                <label className="adminPanel__usersControl-label">Логин пользователя:</label>
                <input
                    onChange={(e) => setLogin(e.target.value)}
                    value={login}
                    type="text"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div>
                <label className="adminPanel__usersControl-label">Пароль пользователя:</label>
                <input
                    onChange={(e) => setPassword(e.target.value)}
                    value={password}
                    type="password"
                    className="adminPanel__usersControl-input"
                />
              </div>
              <div>
                <label className="adminPanel__usersControl-label">Роль пользователя:</label>
                <select
                    className="adminPanel__usersControl-input"
                    onChange={(e) => setRole(e.target.value)}
                    value={role}
                >
                  <option value="">Выберите роль</option>
                  {roleList.map((role) => (
                      <option value={role.id} key={role.id}>
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
                  value={value}
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
