import React from "react";
import "./groupCardLinks.scss";
import add from "./../../assets/GroupCard/Add.svg";
import settings from "./../../assets/GroupCard/settings.svg";
import { Link } from "react-router-dom";
const GroupCard = (obj) => {
  return (
    <Link to={obj.link}>
      <div className="groupCard" key={obj.id}>
        <div className="groupCard__header">
          <h1 className="groupCard__group">{obj.group}</h1>
        </div>
      </div>
    </Link>
  );
};

export default GroupCard;
