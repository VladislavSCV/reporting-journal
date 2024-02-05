import "./lessonCard.scss";
const LessonCard = (obj) => {
  return (
    <div className="schedule__lesson" data-modal="modalScheduleLinks">
      <p className="schedule__lessonName">{obj.lesson}</p>
      <p className="schedule__lessonTeacher">{obj.teacher}</p>
    </div>
  );
};

export default LessonCard;
