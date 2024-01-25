import UserNav from "../UserNav/UserNav";
import './curatorGroups.scss'
import GroupCard from '../GroupCard/GroupCard'

import {objectCuratorGroupCard} from "./../../helpers/objectCuratorGroupCard"

const CuratorGroups = () => {
  return (
    <div className="curatorGroups">
      <UserNav />
      <div className="curatorGroups__container">
      {objectCuratorGroupCard.map((obj, index)=>{
        return <GroupCard group={obj.group} key={index}/>})}
      </div>
    </div>
  );
};

export default CuratorGroups;