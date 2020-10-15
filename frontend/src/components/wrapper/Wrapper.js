import React from 'react';
import NewsCard from '../newscard/NewsCard'
import SwipeableViews from 'react-swipeable-views';

const data = require('../../data.json');

const news = data.map(({ id, image, title, content, date }) => <NewsCard key={id} imageURL={image} title={title} content={content} date={date} />)

const Wrapper = () => {
  return (
    <SwipeableViews>
      {news}
    </SwipeableViews>
  );
};


export default Wrapper;