import React from "react";
import "./groupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import { objectGroupCard } from "../../helpers/objectGroupCard";
import add from "./../../assets/Groups/Add.svg";
const GroupsNotes = () => {
  return (
    <div className="groups">
      <div className="groups__container">
        <h1 className="groups__title">Вложения группы:</h1>
        <div className="groups__list">
          {objectGroupCard.map((obj, index) => {
            return (
              <GroupCardLinks group={obj.group} key={index} link={obj.notes} />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default GroupsNotes;
