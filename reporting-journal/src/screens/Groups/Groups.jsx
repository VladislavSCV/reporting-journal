import React from "react";
import "./groups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import { objectGroupCard } from "../../helpers/objectGroupCard";
import add from "./../../assets/Groups/Add.svg";
const CuratorGroup = () => {
  return (
    <div className="groups">
      <div className="groups__container">
        {objectGroupCard.map((obj, index) => {
          return <GroupCard group={obj.group} key={index} />;
        })}

        <div className="groups__add" data-modal="modalGroupAdd">
          <div className="groups__add-container">
            <img src={add} alt="" className="groups__add-img"/>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CuratorGroup;
