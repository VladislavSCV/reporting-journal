import React, { useEffect, useState } from "react";
import "./notes.scss";
import { objectNotes } from "../../helpers/objectNotes";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
import axios from "axios";
import deleteNotes from "../../js/deleteNotes";
const Notes = () => {
  const [notes, setNotes] = useState([]);

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

  // const deleteNote = async (key) => {
  //   try {
  //     await axios.delete(`http://localhost:5001/api/notes/${key}`);

  //     const response = await axios.get("http://localhost:5001/api/notes");
  //     setNotes(response.data);
  //   } catch (error) {
  //     console.error(error);
  //   }
  // };

  return (
    <div className="notes">
      <div className="notes__container">
        <div className="notes__add" data-modal="modalNotesAdd">
          <img src={add} alt="" className="notes__add-img" />
        </div>
        {notes.map((note) => {
          return (
            <NotesCard
              key={note.id}
              id={note.id}
              title={note.title}
              body={note.body}
            />
          );
        })}
      </div>
    </div>
  );
};

export default Notes;
