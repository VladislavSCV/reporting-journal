import './groupCard.scss'
import add  from './../../img/GroupCard/Add.svg'
const GroupCard = (obj) => {
  return (
    <div className='groupCard'>
      <div className="groupCard__header">
        <h1 className='groupCard__group'>{obj.group}</h1>
      </div>
      <div className="groupCard__container">
        
      <img src={add} alt="" className='groupCard__add-btn'/>

      </div>
    </div>

  ) 
}

export default GroupCard