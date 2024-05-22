import React from "react";
import "./notesCard.scss";
import del from "../../assets/NotesCard/delete.svg";
import settings from "../../assets/NotesCard/settings.svg";
import axios from "axios";
const NotesCard = (obj) => {
  const deleteNote = async (key) => {
    try {
      await axios.delete(`http://localhost:5001/api/notes/${key}`);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="notesCard" key={obj.id}>
      <div className="notesCard__container">
        <div className="notesCard__info">
          <h1 className="notesCard__noteName">{obj.title}</h1>
          <p className="notesCard__noteDescription">{obj.body}</p>
        </div>
        <div className="notesCard__buttons">
          <img
            src={del}
            alt=""
            className="notesCard__buttons-delete"
            // data-modal="modalNotesDelete"
            onClick={() => deleteNote(obj.id)}
          />
          <img
            src={settings}
            alt=""
            className="notesCard__buttons-settings"
            data-modal="modalNotesSettings"
            data-id={obj.id}
          />
        </div>
      </div>
    </div>
  );
};

export default NotesCard;
