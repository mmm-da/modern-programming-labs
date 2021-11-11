import Router from './Router'
import Nav from './components/Nav/Nav'
import Footer from './components/Footer/Footer'
import React from 'react';

const App = () => {
  return (
    <div>
      <Nav/>
      <Router/>
      <Footer/>
    </div>
  );
}

export default App;
