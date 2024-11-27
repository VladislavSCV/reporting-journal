import React, { useEffect, useState } from "react";
import "./groupsSchedule.scss";
import GroupCardLinks from "../../components/GroupCardLinks/GroupCardLinks";
import axios from "axios";
const GroupsSchedule = () => {
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

  console.log(groups)
  return (
    <div className="groups">
      <div className="groups__container">
        <h1 className="groups__title">Расписание группы:</h1>
        <div className="groups__list">

          {Array.isArray(groups) && groups.map((obj) => (
              <GroupCardLinks
                  group={obj.name}
                  key={obj.id}
                  id={obj.id}
                  link={"/schedule/" + obj.id}
              />
          ))}

        </div>
      </div>
    </div>
  );
};

export default GroupsSchedule;
