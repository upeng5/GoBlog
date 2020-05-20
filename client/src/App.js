import React from "react";
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import "./App.css";
import { AuthProvider } from "./contexts/AuthContext";

// Components
import NavBar from "./components/NavBar";
import Post from "./components/Post";

// Views
import Feed from "./views/Feed/Feed";

function App() {
  return (
    <AuthProvider>
      <Router>
        <div className="App">
          <NavBar />
          {/* <Switch>
            <Route exact path="/" component={Feed} />
            <Route exact path="/feed" component={Feed} />
            <Route exact path="/profile" component={Feed} />
            <Route exact path="/login" component={Feed} />
            <Route exact path="/register" component={Feed} />
          </Switch> */}
        </div>
      </Router>
      <Post />
    </AuthProvider>
  );
}

export default App;
