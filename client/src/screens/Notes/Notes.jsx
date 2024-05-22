import React, { useEffect, useState } from "react";
import "./notes.scss";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
import axios from "axios";
const Notes = () => {
  const [notes, setNotes] = useState([]);
  let groupId = Number(window.location.search.substring(1).split("=")[1]);
  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/notes");
        setNotes(response.data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchNotes();
  }, []);

  return (
    <div className="notes">
      <div className="notes__container">
        <div
          className="notes__add"
          data-modal="modalNotesAdd"
          data-id={groupId}
        >
          <img src={add} alt="" className="notes__add-img" />
        </div>
        {notes.map((note) => {
          if (note.groupId === groupId) {
            return (
              <NotesCard
                key={note.id}
                id={note.id}
                title={note.title}
                body={note.body}
              />
            );
          }
        })}
      </div>
    </div>
  );
};

export default Notes;
