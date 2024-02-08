import { objectNotes } from "../helpers/objectNotes";

function addStudent() {
  let noteName = document.getElementById("noteName").value;
  let noteDescription = document.getElementById("noteDescription").value;

  const newObject = { name: `${noteName}`, description: `${noteDescription}` };

  objectNotes.push(newObject);
}

export default addStudent;
