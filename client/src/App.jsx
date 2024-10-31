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
import { auth } from "./actions/api.js";

function App() {
  const [modal, setModal] = useState({ isOpen: false, id: null, day: null });
  const isAuth = useSelector((state) => state.user.isAuth);
  const currentUser = useSelector((state) => state.user.currentUser);
  const dispatch = useDispatch();

  useEffect(() => {
    if (localStorage.getItem("token")) {
      dispatch(auth());
    }
  }, [dispatch]);

  const openModal = (event) => {
    event.preventDefault();
    const { modal, id, day } = event.target.dataset;
    if (modal) {
      setModal({ isOpen: true, id, day });
    }
  };

  const closeModal = () => {
    setModal({ isOpen: false, id: null, day: null });
  };

  const getNavigationComponent = () => {
    switch (currentUser?.role) {
      case "Админ":
        return <AdminNavigation />;
      case "Преподаватель":
        return <TeacherNavigation />;
      case "Староста":
        return <HeadmanNavigation />;
      default:
        return <UserNavigation />;
    }
  };

  return (
      <BrowserRouter>
        <div className="container" onClick={openModal}>
          {isAuth && getNavigationComponent()}
          <ScreenSwitchboard />
          <ModalManager closeFn={closeModal} modal={modal.isOpen} id={modal.id} day={modal.day} />
        </div>
        <Footer />
      </BrowserRouter>
  );
}

export default App;
