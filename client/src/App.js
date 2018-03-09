import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  state = {messsage: ""}
  async componentDidMount() {
    let greeting = await fetch("/api/hello").then(r => r.json());
    let version = await fetch("/api/version").then(r => r.json());
    this.setState({message: greeting.message, version: version});
  }
  renderVersion(version) {
    let style = {
      padding: "1em"
    };
    return (
      <p>
        <span style={style}><em>Commit</em> {version.commit}</span>
        <span style={style}><em>Branch</em> {version.branch}</span>
        <span style={style}><em>State</em> {version.state}</span>
        <span style={style}><em>Timestamp</em> {version.timestamp}</span>
      </p>
    );
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">{this.state.message==="" ?  "Welcome to React" : this.state.message}</h1>
        </header>
        <p className="App-intro">
          {!!this.state.version ? this.renderVersion(this.state.version) : "Loading..." }
        </p>
      </div>
    );
  }
}

export default App;
