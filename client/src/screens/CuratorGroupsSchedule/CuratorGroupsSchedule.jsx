import "./curatorGroupsSchedule.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import { objectCuratorGroupCard } from "../../helpers/objectCuratorGroupCard";
import add from "./../../assets/Groups/Add.svg";
const CuratorGroupsSchedule = () => {
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        <h1 className="groups__title">Расписание группы:</h1>
        <div className="groups__list">
          {objectCuratorGroupCard.map((obj, index) => {
            return (
              <GroupCardLinks
                link={obj.schedule}
                group={obj.group}
                key={index}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default CuratorGroupsSchedule;
