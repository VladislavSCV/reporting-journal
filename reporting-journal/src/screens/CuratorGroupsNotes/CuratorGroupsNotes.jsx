import "./curatorGroupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import { objectCuratorGroupCard } from "../../helpers/objectCuratorGroupCard";
import add from "./../../assets/Groups/Add.svg";
const CuratorGroupsNotes = () => {
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        <h1 className="groups__title">Вложения группы:</h1>
        <div className="groups__list">
          {objectCuratorGroupCard.map((obj, index) => {
            return (
              <GroupCardLinks link={obj.notes} group={obj.group} key={index} />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default CuratorGroupsNotes;
