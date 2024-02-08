import React from "react";
import "./notesCard.scss";
import del from "../../assets/NotesCard/delete.svg";
import settings from "../../assets/NotesCard/settings.svg";
const NotesCard = (obj) => {
  return (
    <div className="notesCard">
      <div className="notesCard__container">
        <div className="notesCard__info">
          <h1 className="notesCard__noteName">{obj.name}</h1>
          <p className="notesCard__noteDescription">{obj.description}</p>
        </div>
        <div className="notesCard__control">
          <img
            src={del}
            alt=""
            className="notesCard__btn"
            data-modal="modalNotesDelete"
          />
          <img
            src={settings}
            alt=""
            className="notesCard__btn"
            data-modal="modalNotesSettings"
          />
        </div>
      </div>
    </div>
  );
};

export default NotesCard;
