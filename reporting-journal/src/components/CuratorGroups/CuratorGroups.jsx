import "./curatorGroups.scss";
import GroupCard from "../GroupCard/GroupCard";
import { Route, Routes, Link, Outlet} from "react-router-dom";
import { objectCuratorGroupCard } from "./../../helpers/objectCuratorGroupCard";
import StudentsList from "./../StudentsList/StudentsList";
const CuratorGroups = () => {
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        {objectCuratorGroupCard.map((obj, index) => {
          return <GroupCard link={obj.link} group={obj.group} key={index} />;
        })}
      </div>
    </div>
  );
};

export default CuratorGroups;
