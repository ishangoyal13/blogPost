class App extends React.Component {
  render() {
    return (
      <Home />
      //<BasicExample />
    )
  }
}

class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
          <h2>Blog App</h2>
          <p>Sign in to get access </p>
          <a onClick={this.authenticate}  className="btn btn-primary btn-lg btn-login btn-block mb-5">Sign In</a>
        </div>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('app'))
