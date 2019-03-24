import React from 'react'
import ReactDOM from 'react-dom'
import './style.styl'

class Home extends React.Component {
  render() {
    return <div>Hello world</div>;
  }
}

ReactDOM.render(<Home/>, document.getElementById("app"));
