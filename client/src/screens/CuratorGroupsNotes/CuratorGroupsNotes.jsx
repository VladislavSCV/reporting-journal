import React, { useState, useEffect } from "react";
import "./curatorGroupsNotes.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const CuratorGroupsNotes = () => {
  const [groups, setGroups] = useState([]);
  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("/api/teacher/groups", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          }
        });
        setGroups(response.data.groups);
      } catch (error) {
        console.error(error);
      }
    };

    fetchGroups();
  }, []);
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
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

export default CuratorGroupsNotes;
