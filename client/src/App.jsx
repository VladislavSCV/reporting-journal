import React, { useState, useEffect } from "react";
import AdminNavigation from "./screens/Navigation/AdminNavigation/AdminNavigation";
import TeacherNavigation from "./screens/Navigation/TeacherNavigation/TeacherNavigation";
import HeadmanNavigation from "./screens/Navigation/HeadmanNavigation/HeadmanNavigation";
import UserNavigation from "./screens/Navigation/UserNavigation/UserNavigation";
import ScreenSwitchboard from "./routes/ScreenSwitchboard";
import ModalManager from "./routes/ModalManager";
import Footer from "./screens/Footer/Footer";
import { BrowserRouter } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import axios from "axios";
import { store } from "./reducers";
import { auth } from "./actions/users";
function App() {
  const [modalOpen, setModal] = useState(false);
  const [id, setId] = useState(false);
  const [day, setDay] = useState(false);
  const isAuth = useSelector((state) => state.user.isAuth);
  const dispatch = useDispatch();

  useEffect(() => {
    if (localStorage.getItem("token")) {
      dispatch(auth());
    }
  }, []);

  const openModal = (event) => {
    event.preventDefault();
    const {
      target: {
        dataset: { modal },
        dataset: { id },
        dataset: { day },
      },
    } = event;
    if (modal) {
      setModal(modal);
      setId(id);
      setDay(day);
    }
  };

  const closeModal = () => {
    setModal("");
  };

  function getRole() {
    switch (store.getState().user.currentUser.role) {
      case "Админ":
        return <AdminNavigation />;
      case "Преподаватель":
        return <TeacherNavigation />;
      case "Староста":
        return <HeadmanNavigation />;
      default:
        return <UserNavigation />;
    }
  }

  return (
    <>
      <BrowserRouter>
        <div className="container" onClick={openModal}>
          {isAuth && getRole()}
          {console.log(store.getState())}
          <ScreenSwitchboard />
          <ModalManager
            closeFn={closeModal}
            modal={modalOpen}
            id={id}
            day={day}
          />
        </div>
        <Footer />
      </BrowserRouter>
    </>
  );
}

export default App;
