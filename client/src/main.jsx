import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.jsx";
import "./scss/main.scss";
import { store } from "./reducers";
import { Provider } from "react-redux";

ReactDOM.createRoot(document.getElementById("root")).render(
  <Provider store={store}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </Provider>
);
