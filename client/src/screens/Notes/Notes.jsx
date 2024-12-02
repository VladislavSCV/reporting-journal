import React, { useEffect, useState } from "react";
import "./notes.scss";
import NotesCard from "../../components/NotesCard/NotesCard";
import add from "../../assets/Notes/Add.svg";
import { useParams } from "react-router-dom";
import axios from "axios";

const Notes = () => {
  const [notes, setNotes] = useState([]);
  const [showForm, setShowForm] = useState(false); // Управление видимостью формы
  const [noteTitle, setNoteTitle] = useState(""); // Заголовок заметки
  const [noteBody, setNoteBody] = useState(""); // Тело заметки
  const params = useParams();
  const userId = params.id || localStorage.getItem("user_id");
  const groupId = localStorage.getItem("group_id");

  useEffect(() => {


    if (!userId) {
      console.error("User ID is not found in URL or localStorage");
      return;
    }

    if (params.id) {
      localStorage.setItem("user_id", params.id);
    }

    const fetchNotes = async () => {
      try {
        const response = await axios.get(`/api/note/${userId}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        setNotes(response.data.notes || []);
      } catch (error) {
        console.error("Error fetching notes:", error);
      }
    };

    fetchNotes();
  }, [params.id]);

  const handleAddNote = async () => {
    console.log(userId);
    console.log(groupId);

    if (!noteTitle.trim() || !noteBody.trim()) {
      alert("Title and body are required!");
      return;
    }

    if (!groupId) {
      console.error("Group ID is missing or invalid.");
      alert("Please select a valid group.");
      return;
    }

    // Временно добавляем заметку сразу, не дожидаясь ответа от API
    const newNote = {
      id: Date.now(), // Создаем уникальный ID для заметки (используем время для простоты)
      title: noteTitle,
      body: noteBody,
    };

    // Добавляем новую заметку в массив
    setNotes((prevNotes) => [...prevNotes, newNote]);

    // Очистка формы
    setNoteTitle("");
    setNoteBody("");
    setShowForm(false);

    // Здесь можно добавить свой код для отправки данных на сервер (если необходимо)
    try {
      const response = await axios.post(
          `/api/note/${userId}`,
          {
            group_id: Number(groupId),
            title: noteTitle,
            body: noteBody,
          },
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
            },
          }
      );
      console.log("API Response:", response.data);
    } catch (error) {
      console.error("Error adding note:", error);
      alert("Failed to add note. Please try again later.");
    }
  };




  return (
      <div className="notes">
        <div className="notes__container">
          <div className="notes__add">
            <img
                src={add}
                alt="Add Note"
                className="notes__add-img"
                onClick={() => setShowForm(!showForm)}
            />
          </div>

          {showForm && (
              <div className="notes__form">
                <input
                    type="text"
                    placeholder="Title"
                    value={noteTitle}
                    onChange={(e) => setNoteTitle(e.target.value)}
                    className="notes__form-input"
                />
                <textarea
                    placeholder="Body"
                    value={noteBody}
                    onChange={(e) => setNoteBody(e.target.value)}
                    className="notes__form-textarea"
                />
                <button onClick={handleAddNote} className="notes__form-button">
                  Save Note
                </button>
                <button
                    onClick={() => setShowForm(false)}
                    className="notes__form-button notes__form-button--cancel"
                >
                  Cancel
                </button>
              </div>
          )}

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
