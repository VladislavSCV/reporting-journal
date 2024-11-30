import React, { useEffect, useState } from "react";
import "./adminPanel.scss";
import UserCard from "../../components/UserCard/UserCard";
import del from "../../assets/AdminPanel/delete.svg";
import {registerUser, createRole} from "../../actions/api";

const AdminPanel = () => {
  const [users, setUsers] = useState([]);
  const [firstName, setFirstName] = useState("");
  const [middleName, setMiddleName] = useState("");
  const [lastName, setLastName] = useState("");
  const [group, setGroup] = useState("");
  const [groups, setGroups] = useState([]);
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const [roleList, setRoles] = useState([]);
  const [value, setValue] = useState("");
  const [loading, setLoading] = useState(true);

  const userRole = localStorage.getItem("userRole");
  let adminPrivileges = userRole === "Admin";

  useEffect(() => {
    const fetchData = async () => {
      try {
        const userResponse = await fetch("/api/admin/AdminPanel", {
          method: "GET",
          headers: { "Authorization": `Bearer ${localStorage.getItem('token')}` }
        });
        if (!userResponse.ok) throw new Error(`Ошибка запроса: ${userResponse.status}`);
        const database64 = await userResponse.json();
        const data = JSON.parse(atob(database64));

        console.log('Received data:', data); // Логируем полученные данные
        setUsers(data.users || []);
        setGroups(data.groups || []);
        setRoles(data.roles || []);
      } catch (error) {
        console.error("Ошибка загрузки данных:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div>Загрузка...</div>;
  }

  // Add user
  const addUser = async (e) => {
    e.preventDefault();
    if (!firstName || !middleName || !lastName || !login || !password || !role || group) {
      alert("Пожалуйста, заполните все поля.");
      return;
    }
    console.log("I AM HERE!")
    registerUser(firstName, middleName, lastName, role, group, login, password);
  }

  // Add role
  const addRole = async (e) => {
    e.preventDefault();
    if (!value) {
      alert("Пожалуйста, введите название роли.");
      return;
    }

    try {
      const response = await fetch("/api/role", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ value }),
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      const data = await response.json();
      setRoles([...roleList, data]);
      setValue("");  // Clear the input after adding
    } catch (error) {
      console.error("Ошибка при добавлении роли:", error);
    }
  };

  // Delete role
  const deleteRole = async (roleId) => {
    try {
      const response = await fetch(`/api/role/${roleId}`, {
        method: "DELETE",
      });
      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      setRoles(roleList.filter((role) => role.id !== roleId));
    } catch (error) {
      console.error("Ошибка при удалении роли:", error);
    }
  };

  const handleChange = (e) => {
    // Преобразуем значение в число и сохраняем в состоянии
    setGroup(parseInt(e.target.value, 10)); // Преобразуем строку в число
  };

  // Clear user form after adding
  const clearForm = () => {
    setFirstName("");
    setMiddleName("");
    setLastName("");
    setGroup("");
    setLogin("");
    setPassword("");
    // setRole("");
  };

  function register() {
    clearForm()
    registerUser(firstName, middleName, lastName, login, password, role, group)
  }

  return (
      <div className="adminPanel">
        {adminPrivileges ? (
            <div className="adminPanel__container">
              <h1 className="adminPanel__title">Панель управления</h1>
              <div className="adminPanel__usersControl">
                <form action="">
                  <div className="">
                    <label className="adminPanel__usersControl-label">*Фамилия:</label>
                    <input
                        onChange={(e) => setFirstName(e.target.value)}
                        value={firstName}
                        type="text"
                        className="adminPanel__usersControl-input"
                    />
                  </div>
                  <div className="">
                    <label className="adminPanel__usersControl-label">*Имя:</label>
                    <input
                        onChange={(e) => setMiddleName(e.target.value)}
                        value={middleName}
                        type="text"
                        className="adminPanel__usersControl-input"
                    />
                  </div>
                  <div className="">
                    <label className="adminPanel__usersControl-label">*Отчество:</label>
                    <input
                        onChange={(e) => setLastName(e.target.value)}
                        value={lastName}
                        type="text"
                        className="adminPanel__usersControl-input"
                    />
                  </div>
                  <div>
                    <label className="adminPanel__usersControl-label">Группа:</label>
                    <select
                        className="adminPanel__usersControl-input"
                        onChange={handleChange}
                        value={group || ''}  // Если group null, то значение будет ''
                    >
                      <option value="">Выберите группу</option>
                      {groups.length > 0 ? (
                          groups.map((group) => (
                              <option value={group.id} key={group.id}>
                                {group.name}
                              </option>
                          ))
                      ) : (
                          <option disabled>Нет доступных групп</option>
                      )}
                    </select>

                  </div>
                  <div>
                    <label className="adminPanel__usersControl-label">*Логин пользователя:</label>
                    <input
                        onChange={(e) => setLogin(e.target.value)}
                        value={login}
                        type="text"
                        className="adminPanel__usersControl-input"
                    />
                  </div>
                  <div>
                    <label className="adminPanel__usersControl-label">*Пароль пользователя:</label>
                    <input
                        onChange={(e) => setPassword(e.target.value)}
                        value={password}
                        type="password"
                        className="adminPanel__usersControl-input"
                    />
                  </div>
                  <div>
                    <label className="adminPanel__usersControl-label">*Роль:</label>
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
                  <button
                      className="adminPanel__usersControl-button"
                      onClick={() => register()}
                  >
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
                <form className="adminPanel__roleControl-form">
                  <label className="adminPanel__roleControl-label">Новая роль:</label>
                  <input
                      type="text"
                      onChange={(e) => setValue(e.target.value)}
                      value={value}
                      className="adminPanel__roleControl-input"
                  />
                  <button type="submit" className="adminPanel__roleControl-button"
                          onClick={() => createRole(value)}>
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
        ) : (
            <div className="adminPanel__message">
              <p>Доступ к этой панели есть только у администратора.</p>
            </div>
        )}
      </div>
  );
};

export default AdminPanel;
