import React, {Component} from 'react';
import {Link} from 'react-router';

class HomePage extends Component {
  render (){
    return (
        <div className="jombotron">
          <h1>Home</h1>
          <p>Text</p>
          <Link to="about" className="btn btn-primary btn-lg">Learn more</Link>
        </div>
    );
  }
}

export default HomePage;
