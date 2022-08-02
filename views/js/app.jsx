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
          <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block mb-5">Sign In</a>
        </div>
      </div>
    )
  }
}
class Blog extends React.Component {
  render() {
    return (
      <div class="card" style="width: 18rem;">
        {/* <img class="card-img-top" src="..." alt="Card image cap"> */}
          <div class="card-body">
            <h5 class="card-title">Card title</h5>
            <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
            <a href="#" class="btn btn-primary">Go somewhere</a>
          </div>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('app'))
