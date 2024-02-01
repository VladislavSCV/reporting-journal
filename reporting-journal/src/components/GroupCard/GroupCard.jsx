import "./groupCard.scss";
import add from "./../../img/GroupCard/Add.svg";
import settings from "./../../img/GroupCard/settings.svg";


const GroupCard = (obj) => {
  return (
    
      <div className="groupCard">
        <div className="groupCard__header">
          <h1 className="groupCard__group">{obj.group}</h1>
        </div>
        <div className="groupCard__container">
          <div className="groupCard__buttons">
            <img src={settings} alt="" className="groupCard__add-btn" data-modal="modalGroupSettings"/>
            <img src={add} alt="" className="groupCard__add-btn" data-modal="modalGroupLinks"/>
          </div>
        </div>
      </div>
  );
};

export default GroupCard;
