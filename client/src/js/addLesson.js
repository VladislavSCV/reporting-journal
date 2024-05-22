import { monday } from "../helpers/Schedule/monday";

function addLesson() {
  let lesson = document.getElementById("chooseLesson");
  let lessonOption = lesson.options[lesson.selectedIndex];
  let lessonText = lessonOption.text;

  let teacher = document.getElementById("chooseTeacher");
  let teacherOption = teacher.options[teacher.selectedIndex];
  let teacherText = teacherOption.text;

  const newLesson = { lesson: `${lessonText}`, teacher: `${teacherText}` };

  monday.push(newLesson);
}

export default addLesson;
