


function addLesson() {
  let lesson = document.getElementById("chooseLesson");
  let lessonOption = lesson.options[lesson.selectedIndex];
  let lesonText = lessonOption.text


  let teacher = document.getElementById("chooseTeacher");
  let teacherOption = teacher.options[teacher.selectedIndex];
  let teacherText = teacherOption.text


  const day = document.getElementById("monday");

  
  const newLesson = `<div class="schedule__lesson">
  <p class="schedule__lessonName">${lesonText}</p>
  <p class="schedule__lessonTeacher">${teacherText}</p>
</div>`

  
  day.insertAdjacentHTML('afterbegin', newLesson)

  
}

export default addLesson;