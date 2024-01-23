import React, { useState } from "react";
import Navigation from "./components/Navigation/Navigation";
import Groups from "./components/Groups/Groups";
import StudentsList from "./components/StudentsList/StudentsList";
import Schedule from "./components/Schedule/Schedule";
import ScreenSwitchboard from "./components/ScreenSwitchboard";
import ModalManager from "./components/ModalManager";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
function App() {
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
            <Navigation />
            <ScreenSwitchboard />
            <ModalManager closeFn={closeModal} modal={modalOpen} />
          </div>
        </BrowserRouter>
    </>
  );
}

export default App;
