import React from 'react'

export default class CmcOut extends React.Component {
  constructor() {
    super()
    this.state = {
      loading: 0
    }
    this.startLoading()
  }
  startLoading() {
    setTimeout(() => {
      this.setState(state => ({
        loading: state.loading == 2 ? 0 : state.loading + 1
      }), () => {
        this.startLoading()
      })
    }, 400)
  }
  render() {
    const command = this.props.command
    return <div className="cmdOutBlock">
      <div className="header">
        {command ? command.Command : ''}
      </div>
      {command ? 
        <div className="lines">
          {command.Out.length == 0 ?
            <div className="noLogs">No logs</div>
          : command.Out.map((line, key) => 
            <div className="line" key={key}>{line}</div>
          )}
          {command.State == 'running' ?
            <div className="loading">
              {([0,1,2]).map(el => 
                <div key={el} className={'dot ' + (this.state.loading == el ? 'show' : 'hide')}></div>
              )}
            </div>
          :''}
        </div>
      : ''}
    </div>;
  }
}
