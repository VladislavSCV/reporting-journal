import React, { useEffect, useState } from "react";
import "./groups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import add from "./../../assets/Groups/Add.svg";
import axios from "axios";
const Groups = () => {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("http://localhost:8000/api/group");
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
        <div className="groups__list">
          {groups.map((obj) => {
            return <GroupCard name={obj.name} key={obj.id} id={obj.id} />;
          })}

          <div className="groups__add" data-modal="modalGroupAdd">
            <img src={add} alt="" className="groups__add-img" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Groups;
