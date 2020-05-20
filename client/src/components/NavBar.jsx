import React from "react";
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
    color: "#fff",
    textDecoration: "none",
    fontWeight: "bold",
    letterSpacing: 2,
  },
  title: {
    flexGrow: 1,
  },
}));

function NavBar() {
  const classes = useStyles();
  // Just for testing use context api for production
  const isAuth = false;
  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            className={classes.menuButton}
            color="inherit"
            aria-label="menu"
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            GoBlog
          </Typography>
          {isAuth ? (
            <div>
              <Link to="/feed" className={classes.menuButton}>
                <Button color="inherit">Home</Button>
              </Link>
              <Link to="/profile" className={classes.menuButton}>
                <Button color="inherit">Profile</Button>
              </Link>
            </div>
          ) : (
            <div>
              <Link to="/login" className={classes.menuButton}>
                <Button color="inherit">Login</Button>
              </Link>
              <Link to="/register" className={classes.menuButton}>
                <Button color="inherit">Register</Button>
              </Link>
            </div>
          )}
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default NavBar;
