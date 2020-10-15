import React from 'react';
import './App.css';
import Wrapper from './components/wrapper/Wrapper';
import NavBar from './components/navbar/NavBar';
import BottomBar from './components/bottombar/BottomBar';
class App extends React.Component {
  // componentDidMount=()=>{
  //   alert('If you are PC please reload website once after changing viewport to mobile device. If already on mobile device you are good to go!');
  // }
  render(){
    return (
      <div className="App">
        <NavBar />
        <Wrapper />
        <BottomBar />
      </div>
    );
  }
  
}

export default App;
