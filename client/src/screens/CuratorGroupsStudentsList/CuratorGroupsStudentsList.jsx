import React, { useState, useEffect } from "react";
import "./curatorGroupsStudentsList.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const CuratorGroupsStudentsList = () => {
  const [groups, setGroups] = useState([]);
  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("/api/group", {
          headers: {
            Authorization: localStorage.getItem("token"),
          },
        });
        setGroups(response.data.groups);
      } catch (error) {
        console.error(error);
      }
    };

    fetchGroups();
  }, [groups]);
  return (
    <div className="curatorGroups">
      <div className="curatorGroups__container">
        <h1 className="groups__title">Список студентов группы:</h1>
        <div className="groups__list">
          {groups.map((obj) => {
            return (
              <GroupCardLinks
                group={obj.name}
                key={obj.id}
                id={obj.id}
                link={"/studentsList?id=" + obj.id}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default CuratorGroupsStudentsList;
