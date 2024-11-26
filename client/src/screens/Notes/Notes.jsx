import React, { useEffect, useState } from "react";
import "./notes.scss";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
import axios from "axios";
const Notes = () => {
  const [notes, setNotes] = useState([]);
  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const response = await axios.get(`/api/note/${localStorage.getItem("user_id")}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          }
        });
        console.log(response)
        setNotes(response.data.notes);
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
          // data-id={groupId}
        >
          <img src={add} alt="" className="notes__add-img" />
        </div>
        {notes.map((note) => {
          // if (note.groupId === groupId) {
            return (
              <NotesCard
                key={note.id}
                id={note.id}
                title={note.title}
                body={note.body}
              />
            );
          // }
        })}
      </div>
    </div>
  );
};

export default Notes;
