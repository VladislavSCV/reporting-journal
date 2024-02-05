import { objectGroupCard } from "./../helpers/objectGroupCard"


function addStudent() {
  let groupName = document.getElementById("groupName").value;


  const newObject = { group: `${groupName}`,
  }

  objectGroupCard.push(newObject)

  
}

export default addStudent;
