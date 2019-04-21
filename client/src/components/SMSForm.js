import React, { Component } from "react";

class SMSForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      messages: [{ message: "" }],
      phone: "",
      telErr: "",
      emptyMsgErr: ""
    };

    this.submitButton = React.createRef();
  }

  handleTelChange = event => {
    this.setState({ phone: event.target.value, telErr: "" });
  };

  handleAddMessage = () => {
    if (this.state.messages.length === 3) {
      return;
    }
    this.setState(prevState => ({
      messages: [...prevState.messages, { message: "" }]
    }));
  };

  createExtraMessages() {
    let addedMessages = this.state.messages.slice(
      1,
      this.state.messages.length
    );

    return addedMessages.map((el, i) => (
      <div key={i + 1}>
        <div className="pure-u-20-24 mrg-top">
          <textarea
            className="pure-u-1"
            maxLength="160"
            placeholder="SMS Message"
            value={el.message || ""}
            onChange={e => this.handleMessageChange(e, i + 1)}
          />
        </div>
        <div className="pure-u-4-24 mrg-top">
          <button
            type="button"
            className="pure-button button-left red"
            onClick={e => this.handleRemoveMessage(i + 1)}
          >
            <i className="fa fa-trash-o" />
          </button>
        </div>
      </div>
    ));
  }

  handleMessageChange = (e, i) => {
    const { value } = e.target;
    let messages = [...this.state.messages];
    messages[i] = { ...messages[i], message: value };
    this.setState({ messages: messages, emptyMsgErr: "" });
  };

  handleRemoveMessage = i => {
    let messages = [...this.state.messages];
    messages.splice(i, 1);
    this.setState({ messages });
  };

  getMessages() {
    let messages = [];
    this.state.messages.forEach(function(element) {
      if (element.message.trim().length > 0) {
        messages.push(element);
      }
    });
    return messages;
  }

  hasAtleastOneMessage(messages) {
    return messages.length > 0;
  }

  hasValidPhone() {
    if (!this.state.phone) {
      return false;
    }

    const mobileNumber = /^\d{11}$/;
    if (!this.state.phone.match(mobileNumber)) {
      return false;
    }
    return true;
  }

  handleSubmit = event => {
    event.preventDefault();

    if (!this.hasValidPhone()) {
      this.setState({ telErr: "Please enter a valid phone number" });
      return;
    }

    let messages = this.getMessages();

    if (!this.hasAtleastOneMessage(messages)) {
      this.setState({
        emptyMsgErr: "Please enter the text message to be sent"
      });
      return;
    }

    this.submitButton.current.disabled = true;
  };

  render() {
    if (this.props.deliveryStatus && this.submitButton.current != null) {
      this.submitButton.current.disabled = false;
    }

    return (
      <div className="pure-g">
        <div className="pure-u-1 content">
          <form className="pure-form" onSubmit={this.handleSubmit}>
            <fieldset>
              <legend>
                <strong>Send up to 3 SMS messages</strong>
                <button
                  type="submit"
                  className="pure-button right black"
                  ref={this.submitButton}
                >
                  <i className="fa fa-location-arrow" />
                  &nbsp; Send
                </button>
              </legend>
              <div className="pure-u-1">
                <span className="right animated fadeOut">
                  {this.props.deliveryStatus}
                </span>
              </div>
              <div className="pure-u-1">
                <label htmlFor="phone">Enter mobile number</label>
              </div>
              <div className="pure-u-20-24">
                <input
                  className="pure-u-1"
                  id="phone"
                  type="tel"
                  value={this.state.phone}
                  placeholder="61xxxxxxxx"
                  pattern="[0-9]{11}"
                  onChange={this.handleTelChange}
                />
              </div>
              <div className="pure-u-20-24">
                <span className="small-fonts err-text">
                  {this.state.telErr}
                </span>
              </div>

              <div className="pure-u-1 mrg-top">
                <label htmlFor="message">
                  Enter message
                  <span className="small-fonts"> (160 characters max) </span>
                  <span className="small-fonts err-text">
                    {this.state.emptyMsgErr}
                  </span>
                </label>
              </div>
              <div className="pure-u-20-24">
                <textarea
                  className="pure-u-1"
                  id="message"
                  maxLength="160"
                  placeholder="SMS Message"
                  value={this.state.messages[0].message || ""}
                  onChange={e => this.handleMessageChange(e, 0)}
                />
              </div>
              <div className="pure-u-20-24 mrg-top">
                <button
                  type="button"
                  onClick={this.handleAddMessage}
                  className="pure-button button-left green"
                >
                  Add more messages
                </button>
              </div>
              {this.createExtraMessages()}
            </fieldset>
          </form>
        </div>
      </div>
    );
  }
}

export default SMSForm;
