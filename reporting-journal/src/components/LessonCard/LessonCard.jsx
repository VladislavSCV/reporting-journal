import "./lessonCard.scss";
const LessonCard = (obj) => {
  return (
    <div className="schedule__lesson" data-modal="modalScheduleLinks">
      {/* <p className="schedule__lesson-name">{obj.lesson}</p>
      <p className="schedule__lesson-teacher">{obj.teacher}</p> */}
      {obj.lesson}
      <br />
      {obj.teacher}
    </div>
  );
};

export default LessonCard;
