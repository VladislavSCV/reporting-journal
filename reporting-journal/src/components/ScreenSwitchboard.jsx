import React from "react";
import { Route, Routes } from "react-router-dom";

import Groups from "./../components/Groups/Groups";
import StudentsList from "./../components/StudentsList/StudentsList";
import Schedule from "./../components/Schedule/Schedule";

const ScreenSwitchboard = () => {
  return (
    <Routes>
      <Route path="/groups" element={<Groups />} />

      <Route path="/studentsList" element={<StudentsList />} />

      <Route path="/schedule" element={<Schedule />} />

      <Route path="/" element={<Groups />} />
    </Routes>
  );
};

export default ScreenSwitchboard;
