import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

class App extends React.Component {
  
    handleClick() {

      console.log('creating ws connection');

      var ws = new WebSocket('ws://localhost:6161');

      ws.onopen = () => {
        console.log('ws connected');
        // connection opened
        // ws.send('something');  // send a message
      };

    }

    render() {
      return (
        <button className='xts-button' onClick={this.handleClick}>
          Connect
        </button>
      );
    }
  }
  
  // ========================================
  
  ReactDOM.render(
    <App />,
    document.getElementById('root')
  );
