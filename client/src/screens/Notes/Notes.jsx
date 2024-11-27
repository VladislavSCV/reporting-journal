import React, { useEffect, useState } from "react";
import "./notes.scss";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
import { useParams } from "react-router-dom";
import axios from "axios";

const Notes = () => {
  const [notes, setNotes] = useState([]);
  const params = useParams(); // Получаем параметры из URL

  useEffect(() => {
    // Определяем id пользователя или группы
    const userId = params.id || localStorage.getItem("user_id");

    if (!userId) {
      console.error("User ID is not found in URL or localStorage");
      return;
    }

    // Если параметр id есть, обновляем localStorage
    if (params.id) {
      localStorage.setItem("user_id", params.id);
    }

    // Функция для получения заметок
    const fetchNotes = async () => {
      try {
        const response = await axios.get(`/api/note/${userId}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        console.log(response);
        setNotes(response.data.notes || []);
      } catch (error) {
        console.error("Error fetching notes:", error);
      }
    };

    fetchNotes();
  }, [params.id]); // Следим за изменением параметра id в URL

  return (
      <div className="notes">
        <div className="notes__container">
          <div
              className="notes__add"
              data-modal="modalNotesAdd"
              // data-id={userId} // Можно передавать userId в модальное окно
          >
            <img src={add} alt="" className="notes__add-img" />
          </div>
          {notes.map((note) => (
              <NotesCard
                  key={note.id}
                  id={note.id}
                  title={note.title}
                  body={note.body}
              />
          ))}
        </div>
      </div>
  );
};

export default Notes;
