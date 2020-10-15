import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    maxWidth: "100vw",
  },
  title: {
    padding: '5px 16px',
    fontWeight: '400',
    fontFamily: 'Roboto',
    fontSize:'17px',
    color:'black'
  },
  details: {
    display: 'flex',
    justifyContent: 'space-between',
    margin: '20px',
  },
  content: {
    overflow: "hidden",
    fontWeight: "300",
    padding: '0 16px',
  }
});

const ImgMediaCard = ({ imageURL, title, content, date }) => {
  const classes = useStyles();

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardMedia
          component="img"
          alt="img"
          height="240"
          image={imageURL}
        />
        <CardContent>

          <div className={classes.title}>
            {title}
          </div>

          <div className={classes.content}>
            {content}
          </div>

          <div className={classes.details}>
            <Typography variant="body2" color="textSecondary" component="h2">{date}</Typography>
          </div>

        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default ImgMediaCard;