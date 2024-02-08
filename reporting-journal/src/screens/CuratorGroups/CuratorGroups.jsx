import "./curatorGroups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import { objectCuratorGroupCard } from "../../helpers/objectCuratorGroupCard";
import add from "./../../assets/Groups/Add.svg";
const CuratorGroups = () => {
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        {objectCuratorGroupCard.map((obj, index) => {
          return <GroupCard link={obj.link} group={obj.group} key={index} />;
        })}
        <div className="groups__add" data-modal="modalCuratorGroupAdd">
          <img src={add} alt="" className="groups__add-img" />
        </div>
      </div>
    </div>
  );
};

export default CuratorGroups;
