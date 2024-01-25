import './groupCard.scss'
import add  from './../../img/GroupCard/Add.svg'
import { Outlet, Link } from "react-router-dom";
const GroupCard = (obj) => {
  return (
    <Link to={obj.link}>
    <div className='groupCard'>
      <div className="groupCard__header">
        <h1 className='groupCard__group'>{obj.group}</h1>
      </div>
      <div className="groupCard__container">
        
      <img src={add} alt="" className='groupCard__add-btn'/>

      </div>
    </div>
    </Link>

  ) 
}

export default GroupCard