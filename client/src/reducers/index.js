import { combineReducers } from "redux";
import smsReducer from "./smsReducer";

export default combineReducers({
  deliveryStatus: smsReducer
});
