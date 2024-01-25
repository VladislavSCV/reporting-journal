import { objectStudentsList } from "./../helpers/objectStudentsList"


function addStudent() {
  let name = document.getElementById("studentname").value;
  let surename = document.getElementById("studentsurename").value;
  let patronymic = document.getElementById("studentpatronymic").value;
  let role = document.getElementById("studentrole").value;

  const newObject = {surname: `${surename}`,
  name: `${name}`,
  patronymic: `${patronymic}`,
  role: `${role}`}

  objectStudentsList.push(newObject)

  
}

export default addStudent;
