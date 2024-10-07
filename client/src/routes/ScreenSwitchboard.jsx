import React from "react";
import { Route, Routes, Navigate } from "react-router-dom";

import Groups from "../screens/Groups/Groups";
import StudentsList from "../screens/StudentsList/StudentsList";
import Schedule from "../screens/Schedule/Schedule";
import CuratorGroups from "../screens/CuratorGroups/CuratorGroups";
import Main from "../screens/Main/Main";
import StudentAttendance from "../screens/StudentAttendance/StudentAttendance";
import Notes from "../screens/Notes/Notes";
import GroupsNote from "../screens/GroupsNotes/GroupsNotes";
import GroupsSchedule from "../screens/GroupsSchedule/GroupsSchedule";
import GroupsStudentsList from "../screens/GroupsStudentsList/GroupsStudentsList";
import CuratorGroupsNotes from "../screens/CuratorGroupsNotes/CuratorGroupsNotes";
import CuratorGroupsSchedule from "../screens/CuratorGroupsSchedule/CuratorGroupsSchedule";
import CuratorGroupsStudentsList from "../screens/CuratorGroupsStudentsList/CuratorGroupsStudentsList";
import AdminPanel from "../screens/AdminPanel/AdminPanel";
import { useDispatch, useSelector } from "react-redux";

const ScreenSwitchboard = () => {
  const isAuth = useSelector((state) => state.user.isAuth);
  return (
      <>
        {!isAuth ? (
            <Routes>
              <Route path="/main" element={<Main />} />
              <Route path="*" element={<Navigate to="/main" />} />
            </Routes>
        ) : (
            <Routes>
              <Route path="/main" element={<Main />} />
              <Route path="groups" element={<Groups />} />
              <Route path="/curatorgroups/*" element={<CuratorGroups />} />
              <Route path="/studentsList" element={<StudentsList />} />
              <Route path="/schedule" element={<Schedule />} />
              <Route path="/main" element={<Main />} />
              <Route path="/studentAttendance" element={<StudentAttendance />} />
              <Route path="/notes" element={<Notes />} />
              <Route path="/GroupsNotes" element={<GroupsNote />} />
              <Route path="/GroupsSchedule" element={<GroupsSchedule />} />
              <Route path="/GroupsStudentsList" element={<GroupsStudentsList />} />
              <Route path="/CuratorGroupsNotes" element={<CuratorGroupsNotes />} />
              <Route
                  path="/CuratorGroupsSchedule"
                  element={<CuratorGroupsSchedule />}
              />
              <Route
                  path="/CuratorGroupsStudentsList"
                  element={<CuratorGroupsStudentsList />}
              />
              <Route
                  path="/AdminPanel"
                  element={<AdminPanel />}
              />
            </Routes>
        )}
      </>
  );
};

export default ScreenSwitchboard;
