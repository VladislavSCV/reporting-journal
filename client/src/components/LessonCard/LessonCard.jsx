import "./lessonCard.scss";
import del from "../../assets/LessonCard/delete.svg";
import settings from "../../assets/LessonCard/settings.svg";
import axios from "axios";

const LessonCard = (obj) => {
  const deleteSchedule = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/schedule/${key}`);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <div className="schedule__lesson" key={obj.id}>
      <div className="schedule__lesson-container">
        <div className="schedule__lesson-info">
          <p>{obj.lesson}</p>

          <p>{obj.teacher}</p>
        </div>

        <div className="schedule__lesson-buttons">
          <img
            src={settings}
            alt=""
            className="schedule__lesson-buttons-settings"
            data-modal="modalScheduleSettings"
            data-id={obj.id}
          />
          <img
            src={del}
            alt=""
            className="schedule__lesson-buttons-delete"
            onClick={() => deleteSchedule(obj.id)}
          />
        </div>
      </div>
    </div>
  );
};

export default LessonCard;
