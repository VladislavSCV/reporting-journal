import React from "react";
import { Route, Routes } from "react-router-dom";

import Groups from "../screens/Groups/Groups";
import StudentsList from "../screens/StudentsList/StudentsList";
import Schedule from "../screens/Schedule/Schedule";
import CuratorGroups from "../screens/CuratorGroups/CuratorGroups";
import Main from "../screens/Main/Main";
import StudentAttendance from "../screens/StudentAttendance/StudentAttendance";
import Notes from "../screens/Notes/Notes";
const ScreenSwitchboard = () => {
  return (
    <Routes>
      <Route path="/" element={<Main />} />
      <Route path="groups" element={<Groups />} />
      <Route path="/curatorgroups/*" element={<CuratorGroups />} />
      <Route path="/studentsList" element={<StudentsList />} />
      <Route path="/schedule" element={<Schedule />} />
      <Route path="/main" element={<Main />} />
      <Route path="/studentAttendance" element={<StudentAttendance />} />
      <Route path="/notes" element={<Notes />} />
    </Routes>
  );
};

export default ScreenSwitchboard;
