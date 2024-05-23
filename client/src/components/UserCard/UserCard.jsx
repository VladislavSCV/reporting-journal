import React from "react";
import del from "../../assets/UserCard/delete.svg";
import "./userCard.scss";
import axios from "axios";

const UserCard = (obj) => {
  const deleteUser = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/auth/user/${key}`);
    } catch {
      console.error(error);
    }
  };

  return (
    <div className="userCard" key={obj.id}>
      <div className="userCard__container">
        <div className="userCard__info">
          <h1 className="userCard__name">{obj.name}</h1>
          <p className="userCard__password">{obj.password}</p>
          <p className="userCard__role">{obj.role}</p>
          <div className="userCard__buttons">
            <img
              src={del}
              alt=""
              className="userCard__buttons-delete"
              onClick={() => deleteUser(obj.id)}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default UserCard;
