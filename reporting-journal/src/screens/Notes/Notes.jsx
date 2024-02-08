import React from "react";
import "./notes.scss";
import { objectNotes } from "../../helpers/objectNotes";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
const Notes = () => {
  return (
    <div className="notes">
      <div className="notes__container">
        {objectNotes.map((obj, index) => {
          return (
            <NotesCard
              name={obj.name}
              description={obj.description}
              key={index}
            />
          );
        })}

        <div className="notes__add">
          <div className="notes__add-container" data-modal="modalNotesAdd">
            <img src={add} alt="" className="notes__add-img" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Notes;
