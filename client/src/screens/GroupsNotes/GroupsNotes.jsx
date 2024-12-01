import React, { useEffect, useState } from "react";
import "./groupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const GroupsNotes = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("/api/group", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        setGroups(response.data.groups);
      } catch (error) {
        console.error(error);
      }
    };

    fetchGroups();
  }, []);

  return (
    <div className="groups">
      <div className="groups__container">
        <h1 className="groups__title">Вложения группы:</h1>
        <div className="groups__list">
          {groups.map((obj) => {
            return (
              <GroupCardLinks
                group={obj.name}
                key={obj.id}
                id={obj.id}
                link={"/group/notes/" + obj.id}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default GroupsNotes;
