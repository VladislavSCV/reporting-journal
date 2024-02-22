import "./groupCard.scss";
import add from "./../../assets/GroupCard/Add.svg";
import settings from "./../../assets/GroupCard/settings.svg";

const GroupCard = (obj) => {
  return (
    <div className="groupCard">
      <div className="groupCard__header">
        <h1 className="groupCard__group">{obj.group}</h1>
      </div>
      <div className="groupCard__container">
        <div className="groupCard__buttons">
          <img
            src={settings}
            alt=""
            className="groupCard__buttons-links"
            data-modal="modalGroupSettings"
          />
          <img
            src={add}
            alt=""
            className="groupCard__buttons-settings"
            data-modal="modalGroupLinks"
          />
        </div>
      </div>
    </div>
  );
};

export default GroupCard;
