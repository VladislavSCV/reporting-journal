import { objectCuratorGroupCard } from "./../helpers/objectCuratorGroupCard"


function addStudent() {
  let groupName = document.getElementById("curatorGroupName").value;


  const newObject = { group: `${groupName}`,
  }

  objectCuratorGroupCard.push(newObject)

  
}

export default addStudent;
