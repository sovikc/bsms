import axios from "axios";

export const sendSMS = sms => dispatch => {
  axios
    .post("/messaging/v1/sms/", sms)
    .then(function(response) {
      let deliveryStatus = response.data.status;
      console.log(response.data.status);
      dispatch({ type: "SEND_SMS", payload: deliveryStatus });
    })
    .catch(function(error) {
      dispatch({ type: "SEND_SMS", payload: error });
    });
};
