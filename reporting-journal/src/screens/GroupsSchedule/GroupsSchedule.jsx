import React from "react";
import "./groupsSchedule.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import { objectGroupCard } from "../../helpers/objectGroupCard";
import add from "./../../assets/Groups/Add.svg";
const GroupsSchedule = () => {
  return (
    <div className="groups">
      <div className="groups__container">
        <h1 className="groups__title">Выберете группу расписания</h1>
        <div className="groups__list">
          {objectGroupCard.map((obj, index) => {
            return (
              <GroupCard group={obj.group} key={index} link={obj.schedule} />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default GroupsSchedule;
