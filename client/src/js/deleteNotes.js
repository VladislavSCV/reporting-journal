function deleteNotes(notes, id) {
  let del = notes.filter((e) => e.id !== id);
  notes.splice(del, 1);
}

export default deleteNotes;
