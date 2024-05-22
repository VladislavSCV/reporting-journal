import "./groupCard.scss";
import settings from "./../../assets/GroupCard/settings.svg";
import del from "../../assets/GroupCard/delete.svg";
import axios from "axios";

const GroupCard = (obj) => {
  const deleteGroup = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/groups/${key}`);
    } catch (error) {
      console.error(error);
    }
  };
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
          />
          <img
            src={del}
            alt=""
            className="groupCard__buttons-settings"
            onClick={() => deleteGroup(obj.id)}
            // data-modal="modalGroupDelete"
          />
        </div>
      </div>
    </div>
  );
};

export default GroupCard;
