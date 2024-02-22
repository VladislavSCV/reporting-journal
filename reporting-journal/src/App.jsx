import React, { useState } from "react";
import Navigation from "./screens/Navigation/Navigation";
import ScreenSwitchboard from "./routes/ScreenSwitchboard";
import ModalManager from "./routes/ModalManager";
import Footer from "./screens/Footer/Footer";
import { BrowserRouter } from "react-router-dom";
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
        <Footer />
      </BrowserRouter>
    </>
  );
}

export default App;
