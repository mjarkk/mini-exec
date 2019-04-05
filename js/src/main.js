import React from 'react'
import ReactDOM from 'react-dom'
import './style.styl'

import CmcOut from './cmdOut'
import net from './networking'

class Home extends React.Component {
  constructor() {
    super()

    this.state = {
      commands: [
        {
          Command: 'cd',
          Out: [],
          State: 'completed'
        },
        {
          Command: 'test1',
          Out: ['line1', 'line2'],
          State: 'running'
        },
        {
          Command: 'test2',
          Out: [],
          State: 'waiting'
        }
      ],
      viewingCommand: 0
    }
  }
  render() {
    return (
      <div className="site">
        <div className="Commands">
          <div className="controlls">
            <button className="update" >Update</button>
          </div>
          {this.state.commands.map((el, key) => 
            <div 
              key={key} 
              className={'command ' + el.State}
              onClick={() => {
                this.setState({
                  viewingCommand: key
                })
              }}
            >
              {el.Command}
            </div>
          )}
        </div>
        <div className="Output">
          <CmcOut command={this.state.commands[this.state.viewingCommand]}/>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<Home/>, document.getElementById("app"));
