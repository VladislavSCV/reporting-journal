import React from "react";
import {Route, Routes, Navigate} from "react-router-dom";

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
import MainPage from "../screens/MainPage/MainPage.jsx";
import {useDispatch, useSelector} from "react-redux";
import GroupList from "../components/GroupList/GroupList.jsx";
import NotesGroup from "../screens/NotesGroup/Notes.jsx";
import CuratorGroupsNotes_notes from "../screens/NotesCuratorGroup/Notes.jsx";

const ScreenSwitchboard = () => {
  const isAuth = localStorage.getItem("token");
  console.log(localStorage.getItem("user_id"))
  console.log(localStorage.getItem("group_id"))
  console.log(isAuth)
  if (!isAuth) {
    console.log("isAuth is false");
  }
  return (
      <>
        {/* {isAuth === null ? (
            <Routes>
              <Route path="/main" element={<Main/>}/>
              <Route path="/" element={<Navigate to="/main"/>}/>
            </Routes>
        ) : ( */}
            <Routes>
              <Route path="/main" element={<Main/>}/>
              <Route path="*" element={<MainPage/>}/>
              <Route path="/mainPage" element={<MainPage/>}/>
              <Route path="/groups" element={<Groups/>}/>
              <Route path="/studentAttendance/:groupId" element={<StudentAttendance />} />
              <Route path="/groupsList" element={<GroupList />} />
              <Route path="/curatorgroups" element={<CuratorGroups/>}/>
              <Route path="/studentsList" element={<StudentsList/>}/>
              <Route path="/studentsList/:id" element={<StudentsList/>}/>
              <Route path="/schedule" element={<Schedule/>}/>
              <Route path="/schedule/:id" element={<Schedule/>}/>
              <Route path="/main" element={<Main/>}/>
              <Route path="/studentAttendance" element={<StudentAttendance/>}/>
              <Route path="/notes" element={<Notes/>}/>
              <Route path="/notes/:id" element={<Notes/>}/>
              <Route path="/group/notes/:id" element={<NotesGroup/>}/>
              <Route path="/CuratorGroup/notes/:id" element={<NotesGroup/>}/>
              <Route path="/GroupsNotes" element={<GroupsNote/>}/>
              <Route path="/GroupsSchedule" element={<GroupsSchedule/>}/>
              <Route path="/GroupsStudentsList" element={<GroupsStudentsList/>}/>
              <Route path="/CuratorGroupsNotes" element={<CuratorGroupsNotes/>}/>
              <Route
                  path="/CuratorGroupsSchedule"
                  element={<CuratorGroupsSchedule/>}
              />
              <Route
                  path="/CuratorGroupsStudentsList"
                  element={<CuratorGroupsStudentsList/>}
              />
              <Route
                  path="/AdminPanel"
                  element={<AdminPanel/>}
              />
            </Routes>
        {/* )} */}
      </>
  );
};

export default ScreenSwitchboard;

