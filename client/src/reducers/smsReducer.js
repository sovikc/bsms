export default (state = [], action) => {
  switch (action.type) {
    case "SEND_SMS":
      return action.payload;
    default:
      return state;
  }
};
