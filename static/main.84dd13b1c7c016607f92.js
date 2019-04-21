!function(e){function t(t){for(var n,o,l=t[0],u=t[1],c=t[2],m=0,p=[];m<l.length;m++)o=l[m],r[o]&&p.push(r[o][0]),r[o]=0;for(n in u)Object.prototype.hasOwnProperty.call(u,n)&&(e[n]=u[n]);for(i&&i(t);p.length;)p.shift()();return s.push.apply(s,c||[]),a()}function a(){for(var e,t=0;t<s.length;t++){for(var a=s[t],n=!0,l=1;l<a.length;l++){var u=a[l];0!==r[u]&&(n=!1)}n&&(s.splice(t--,1),e=o(o.s=a[0]))}return e}var n={},r={0:0},s=[];function o(t){if(n[t])return n[t].exports;var a=n[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,o),a.l=!0,a.exports}o.m=e,o.c=n,o.d=function(e,t,a){o.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},o.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},o.t=function(e,t){if(1&t&&(e=o(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(o.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)o.d(a,n,function(t){return e[t]}.bind(null,n));return a},o.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return o.d(t,"a",t),t},o.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},o.p="";var l=window.webpackJsonp=window.webpackJsonp||[],u=l.push.bind(l);l.push=t,l=l.slice();for(var c=0;c<l.length;c++)t(l[c]);var i=u;s.push([48,1]),a()}({48:function(e,t,a){a.r(t);var n=a(0),r=a.n(n),s=a(17),o=a.n(s),l=a(5),u=a(19),c=a.n(u);function i(){return(i=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var a=arguments[t];for(var n in a)Object.prototype.hasOwnProperty.call(a,n)&&(e[n]=a[n])}return e}).apply(this,arguments)}function m(e){if(void 0===e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return e}function p(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}var h=function(e){var t,a;function n(t){var a;return p(m(a=e.call(this,t)||this),"handleTelChange",function(e){a.setState({phone:e.target.value,telErr:""})}),p(m(a),"handleAddMessage",function(){3!==a.state.messages.length&&a.setState(function(e){return{messages:[].concat(e.messages,[{message:""}])}})}),p(m(a),"handleMessageChange",function(e,t){var n=e.target.value,r=[].concat(a.state.messages);r[t]=i({},r[t],{message:n}),a.setState({messages:r,emptyMsgErr:""})}),p(m(a),"handleRemoveMessage",function(e){var t=[].concat(a.state.messages);t.splice(e,1),a.setState({messages:t})}),p(m(a),"handleSubmit",function(e){if(e.preventDefault(),a.hasValidPhone()){var t=a.getMessages();a.hasAtleastOneMessage(t)?(a.submitButton.current.disabled=!0,a.props.sendSMS({phone:a.state.phone,messages:t})):a.setState({emptyMsgErr:"Please enter the text message to be sent"})}else a.setState({telErr:"Please enter a valid phone number"})}),a.state={messages:[{message:""}],phone:"",telErr:"",emptyMsgErr:""},a.submitButton=r.a.createRef(),a}a=e,(t=n).prototype=Object.create(a.prototype),t.prototype.constructor=t,t.__proto__=a;var s=n.prototype;return s.createExtraMessages=function(){var e=this;return this.state.messages.slice(1,this.state.messages.length).map(function(t,a){return r.a.createElement("div",{key:a+1},r.a.createElement("div",{className:"pure-u-20-24 mrg-top"},r.a.createElement("textarea",{className:"pure-u-1",maxLength:"160",placeholder:"SMS Message",value:t.message||"",onChange:function(t){return e.handleMessageChange(t,a+1)}})),r.a.createElement("div",{className:"pure-u-4-24 mrg-top"},r.a.createElement("button",{type:"button",className:"pure-button button-left red",onClick:function(t){return e.handleRemoveMessage(a+1)}},r.a.createElement("i",{className:"fa fa-trash-o"}))))})},s.getMessages=function(){var e=[];return this.state.messages.forEach(function(t){t.message.trim().length>0&&e.push(t)}),e},s.hasAtleastOneMessage=function(e){return e.length>0},s.hasValidPhone=function(){if(!this.state.phone)return!1;return!!this.state.phone.match(/^\d{11}$/)},s.render=function(){var e=this;return this.props.deliveryStatus&&null!=this.submitButton.current&&(this.submitButton.current.disabled=!1),r.a.createElement("div",{className:"pure-g"},r.a.createElement("div",{className:"pure-u-1 content"},r.a.createElement("form",{className:"pure-form",onSubmit:this.handleSubmit},r.a.createElement("fieldset",null,r.a.createElement("legend",null,r.a.createElement("strong",null,"Send up to 3 SMS messages"),r.a.createElement("button",{type:"submit",className:"pure-button right black",ref:this.submitButton},r.a.createElement("i",{className:"fa fa-location-arrow"}),"  Send")),r.a.createElement("div",{className:"pure-u-1"},r.a.createElement("span",{className:"right animated fadeOut"},this.props.deliveryStatus)),r.a.createElement("div",{className:"pure-u-1"},r.a.createElement("label",{htmlFor:"phone"},"Enter mobile number")),r.a.createElement("div",{className:"pure-u-20-24"},r.a.createElement("input",{className:"pure-u-1",id:"phone",type:"tel",value:this.state.phone,placeholder:"61xxxxxxxx",pattern:"[0-9]{11}",onChange:this.handleTelChange})),r.a.createElement("div",{className:"pure-u-20-24"},r.a.createElement("span",{className:"small-fonts err-text"},this.state.telErr)),r.a.createElement("div",{className:"pure-u-1 mrg-top"},r.a.createElement("label",{htmlFor:"message"},"Enter message",r.a.createElement("span",{className:"small-fonts"}," (160 characters max) "),r.a.createElement("span",{className:"small-fonts err-text"},this.state.emptyMsgErr))),r.a.createElement("div",{className:"pure-u-20-24"},r.a.createElement("textarea",{className:"pure-u-1",id:"message",maxLength:"160",placeholder:"SMS Message",value:this.state.messages[0].message||"",onChange:function(t){return e.handleMessageChange(t,0)}})),r.a.createElement("div",{className:"pure-u-20-24 mrg-top"},r.a.createElement("button",{type:"button",onClick:this.handleAddMessage,className:"pure-button button-left green"},"Add more messages")),this.createExtraMessages()))))},n}(n.Component),f=Object(l.b)(function(e){return{deliveryStatus:e.deliveryStatus}},{sendSMS:function(e){return function(t){c.a.post("/messaging/v1/sms/",e).then(function(e){var a=e.data.status;t({type:"SEND_SMS",payload:a})}).catch(function(e){t({type:"SEND_SMS",payload:e})})}}})(h);var d=function(e){var t,a;function n(){return e.apply(this,arguments)||this}return a=e,(t=n).prototype=Object.create(a.prototype),t.prototype.constructor=t,t.__proto__=a,n.prototype.render=function(){return r.a.createElement(f,null)},n}(n.Component),g=a(2),v=a(20),b=Object(g.c)({deliveryStatus:function(e,t){switch(void 0===e&&(e=[]),t.type){case"SEND_SMS":return t.payload;default:return e}}}),E=Object(g.d)(b,Object(g.a)(v.a));o.a.render(r.a.createElement(l.a,{store:E},r.a.createElement(d,null)),document.querySelector("#root"))}});