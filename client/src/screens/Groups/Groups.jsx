import React, { useEffect, useState } from "react";
import "./groups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import add from "./../../assets/Groups/Add.svg";
import {getGroups} from "../../actions/api";

const Groups = () => {
  const [groups, setGroups] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    getGroups().then()
        .then((data) => {
          setGroups(data.groups);
    })
  }, []);

  return (
      <div className="groups">
        <div className="groups__container">
          <div className="groups__list">
            {error && <p>{error}</p>}
            {groups.map((obj) => (
                <GroupCard name={obj.name} key={obj.id} id={obj.id} />
            ))}
            <div className="groups__add" data-modal="modalGroupAdd">
              <img src={add} alt="" className="groups__add-img" />
            </div>
          </div>
        </div>
      </div>
  );
};

export default Groups;