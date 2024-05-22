import React, { useEffect, useState } from "react";
import "./groupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const GroupsNotes = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("http://localhost:5001/api/groups");
        setGroups(response.data);
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
                link={"/notes?id=" + obj.id}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default GroupsNotes;
