import React from 'react';
import PropTypes from 'prop-types';
import {Link, IndexLink} from 'react-router';
import LoadingDots from './LoadingDots';

const Header = ({loading, isAuthenticated, handleLogout}) => {
  return (
    <div className="navbar navbar-default">
      <div className="container-fluid">
        <div className="navbar-header">
          <button
            type="button"
            className="navbar-toggle"
            data-toggle="collapse"
            data-target=".navbar-responsive-collapse">
            <span className="icon-bar"/>
            <span className="icon-bar"/>
            <span className="icon-bar"/>
          </button>
          <IndexLink to="/" className="navbar-brand">Teww</IndexLink>
        </div>
        <div className="navbar-collapse collapse navbar-responsive-collapse">
          <ul className="nav navbar-nav">
            {isAuthenticated && <li>
              <Link to="/items" activeClassName="active">Items</Link>
            </li>}
            <li>
              <Link to="/about" activeClassName="active">About</Link>
              {loading && <LoadingDots interval={100} dots={20}/>}
            </li>
          </ul>
          <ul className="nav navbar-nav navbar-right">
            {isAuthenticated
              ? (
                <li>
                  <Link to="/" onClick={handleLogout}>Log out</Link>
                </li>
              )
              : (
                <li>
                  <Link to="/login">Log in</Link>
                </li>
              )}
          </ul>
        </div>
      </div>
    </div>
  );
};

Header.propTypes = {
  loading: PropTypes.bool.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
  handleLogout: PropTypes.func.isRequired
};

export default Header;
