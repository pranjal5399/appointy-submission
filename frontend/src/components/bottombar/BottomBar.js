import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { blue, green } from '@material-ui/core/colors';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';

import FiberManualRecordIcon from '@material-ui/icons/FiberManualRecord';
import ShareOutlinedIcon from '@material-ui/icons/ShareOutlined';
import BookmarkBorderOutlinedIcon from '@material-ui/icons/BookmarkBorderOutlined';

const useStyles = makeStyles((theme) => ({
    appBar: {
        top: 'auto',
        bottom: 0,
    },
    grow: {
        flexGrow: 1,
    },
}));

const BottomBar = () => {

    const classes = useStyles();
    return (
        <AppBar position="fixed" color="inherit" className={classes.appBar}>
            <Toolbar>
                <div className={classes.grow} />
                <IconButton style={{ color: green[500] }}>
                    <FiberManualRecordIcon />
                </IconButton>
                <div style={{ 'fontSize': '10px' }}>Relevance</div>

                <div className={classes.grow} />
                <IconButton style={{ color: blue[500] }}>
                    <ShareOutlinedIcon />
                </IconButton>
                <div style={{ 'fontSize': '10px' }}>Share</div>

                <div className={classes.grow} />
                <IconButton style={{ color: blue[500] }}>
                    <BookmarkBorderOutlinedIcon />

                </IconButton>
                <div style={{ 'fontSize': '10px' }}>Bookmark</div>

                <div className={classes.grow} />
            </Toolbar>
        </AppBar>
    )
}

export default BottomBar
