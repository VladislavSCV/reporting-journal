import React, { useEffect, useState } from "react";
import "./curatorGroups.scss";
import GroupCard from "../../components/GroupCard/GroupCard";
import axios from "axios";
// import {VerifyTokenAndGetId} from "../../actions/api";
import { store } from "../../reducers";
const CuratorGroups = () => {
  const [groups, setGroups] = useState([]);
  const [id, setId] = useState();

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await axios.get("/api/group");
        setGroups(response.data.groups);
      } catch (error) {
        console.error(error);
      }
    };

    const fetchId = async () => {
      try {
        const response = await fetch('/api/auth/verify', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
          },
          body: JSON.stringify({ token: localStorage.getItem('token') })
        });

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log(data);

        if (data.id) {
          setId(data.id);
        } else {
          console.error('No ID in response:', data);
        }
      } catch (error) {
        console.error('Error fetching ID:', error);
      }
    };

    console.log(id)
    console.log(localStorage.getItem("token"));

    fetchId()
    fetchGroups();
  }, []);

  return (
      <div className="curatorGroups">
        <div className="curatorGroups__container">
          <div className="groups__list">
            {groups.map((obj) => {
              return (
                <GroupCard
                  key={obj.id}
                  id={obj.id}
                  name={obj.name}
                  curator={obj.curator}
                  students={obj.students}
                />
              );
            })}
          </div>
        </div>
      </div>
  );
};

export default CuratorGroups;