import React from "react";
import { Route, Routes } from "react-router-dom";

import Groups from "./screens/Groups/Groups";
import StudentsList from "./screens/StudentsList/StudentsList";
import Schedule from "./screens/Schedule/Schedule";
import CuratorGroups from "./screens/CuratorGroups/CuratorGroups";
import MainPage from "./screens/MainPage/MainPage";
import StudentAttendance from './screens/StudentAttendance/StudentAttendance'
const ScreenSwitchboard = () => {
  return (
    <Routes>
      <Route path="/" element={<MainPage />} />
      <Route path="groups" element={<Groups />} />
      <Route path="/curatorgroups/*" element={<CuratorGroups />} />
      <Route path="/studentsList" element={<StudentsList />} />
      <Route path="/schedule" element={<Schedule />} />
      <Route path="/mainPage" element={<MainPage />} />
      <Route path="/studentAttendance" element={<StudentAttendance/>} />
    </Routes>
  );
};

export default ScreenSwitchboard;
