import React, { Component } from 'react';
import PropTypes from 'prop-types';

class App extends Component {
  
  static proptypesRecibidos = {
    children:  PropTypes.object
  };
  
  render() {
    debugger;
    const {children} = this.props;
    return (
      <div>{children}</div>
    );
  }
}

export default App;

