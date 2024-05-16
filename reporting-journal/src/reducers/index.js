import { applyMiddleware, combineReducers } from "redux";
import { composeWithDevTools } from "@redux-devtools/extension";
import { configureStore } from "@reduxjs/toolkit";

import { thunk } from "redux-thunk";
import userReducer from "./userReducer";

const rootReducer = combineReducers({
  user: userReducer,
});

export const store = configureStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(thunk))
);
