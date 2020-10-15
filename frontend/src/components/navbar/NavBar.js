import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import RefreshIcon from '@material-ui/icons/Refresh';
import IconButton from '@material-ui/core/IconButton';
import { blue } from '@material-ui/core/colors';
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';

const useStyles = makeStyles((theme) => ({
    discover: {
        flexGrow: 0.5,
    },
    title: {
        flexGrow: 1,
    },
}));

const NavBar = () => {
    const classes = useStyles();

    return (

        <AppBar position="static" color="inherit">
            <Toolbar>
                <IconButton edge="start" style={{ color: blue[500] }}>
                    <ArrowBackIosIcon />
                </IconButton>

                <Typography variant="body2" className={classes.discover} >
                    Discover
                </Typography>


                <Typography variant="body2" className={classes.title} >
                    My Feed
                </Typography>


                <IconButton edge="end" style={{ color: blue[500] }}>
                    <RefreshIcon />
                </IconButton>

            </Toolbar>
        </AppBar>

    );
}

export default NavBar;