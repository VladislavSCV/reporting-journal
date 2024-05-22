import "./groupCard.scss";
import add from "./../../assets/GroupCard/Add.svg";
import settings from "./../../assets/GroupCard/settings.svg";
import { Link } from "react-router-dom";

const GroupCard = (obj) => {
  return (
    <div className="groupCard" key={obj.id}>
      <div className="groupCard__header">
        <h1 className="groupCard__group">{obj.name}</h1>
      </div>
      <div className="groupCard__container">
        <div className="groupCard__buttons">
          <img
            src={settings}
            alt=""
            className="groupCard__buttons-links"
            data-modal="modalGroupSettings"
            data-id={obj.id}
            onClick={() => console.log(obj.id)}
          />
          {/* <img
              src={add}
              alt=""
              className="groupCard__buttons-settings"
              data-modal="modalGroupLinks"
            /> */}
        </div>
      </div>
    </div>
  );
};

export default GroupCard;
