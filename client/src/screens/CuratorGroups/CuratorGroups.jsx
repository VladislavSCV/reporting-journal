import React, { useEffect, useState } from "react";
import "./curatorGroups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import axios from "axios";

const CuratorGroups = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("/api/teacher/groups", {
          headers: { "Authorization": `Bearer ${localStorage.getItem('token')}` }
        });
        setGroups(response.data.groups || []);
      } catch (error) {
        console.error(error);
        setGroups([]);
      }
    };

    fetchGroups();
  }, []);

  return (
      <div className="curatorGroups">
        <div className="curatorGroups__container">
          {groups.length === 0 ? (
              <p>Куратор не курирует группы</p>
          ) : (
              <div className="groups__list">
                {groups.map((obj) => (
                    <GroupCard
                        key={obj.id}
                        id={obj.id}
                        name={obj.name}
                        curator={obj.curator}
                        students={obj.students}
                        link={`/studentsList?id=${obj.id}`}
                    />
                ))}
              </div>
          )}
        </div>
      </div>
  );
};

export default CuratorGroups;