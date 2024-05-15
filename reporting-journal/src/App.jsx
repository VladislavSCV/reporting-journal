import React, { useState } from "react";
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
function App() {
  const dispatch = useDispatch();
  const test = async () => {
    try {
      await axios.get(`http://localhost:5001/api/auth/user`);
    } catch (error) {
      console.error(error);
    }
  };
  test();
  const isAuth = useSelector((state) => state.user.isAuth);
  const [modalOpen, setModal] = useState(false);

  const openModal = (event) => {
    event.preventDefault();
    const {
      target: {
        dataset: { modal },
      },
    } = event;
    if (modal) setModal(modal);
  };

  const closeModal = () => {
    setModal("");
  };

  return (
    <>
      <BrowserRouter>
        <div className="container" onClick={openModal}>
          {isAuth && <UserNavigation />}
          <ScreenSwitchboard />
          <ModalManager closeFn={closeModal} modal={modalOpen} />
        </div>
        <Footer />
      </BrowserRouter>
    </>
  );
}

export default App;
