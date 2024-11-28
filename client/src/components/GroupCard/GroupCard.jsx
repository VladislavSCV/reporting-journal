import React from "react";
import "./groupCard.scss";
import settings from "./../../assets/GroupCard/settings.svg";
import del from "../../assets/GroupCard/delete.svg";
import axios from "axios";
import {Link} from "react-router-dom";

const GroupCard = (obj) => {
  const deleteGroup = async (key) => {
    try {
      await axios.delete(`http://localhost:8000/api/groups/${key}`, {
        headers: { "Authorization": `Bearer ${localStorage.getItem('token')}` }
      });
    } catch (error) {
      console.error(error);
    }
  };
  return (
      <Link to={obj.link}>
    <div className="groupCard" key={obj.id}>
      <div className="groupCard__header">
        <h1 className="groupCard__group">{obj.name}</h1>
      </div>
      <div className="groupCard__container">
        <div className="groupCard__buttons">
          <img
            src={settings}
            alt=""
            className="groupCard__buttons-links"
            data-modal="modalGroupSettings"
            data-id={obj.id}
          />
          <img
            src={del}
            alt=""
            className="groupCard__buttons-settings"
            onClick={() => deleteGroup(obj.id)}
            // data-modal="modalGroupDelete"
          />
        </div>
      </div>
    </div>
      </Link>



  );
};

export default GroupCard;
