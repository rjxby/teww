import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {bindActionCreators} from 'redux';
import Header from './common/Header';
import {connect} from 'react-redux';
import * as authActions from '../actions/authActions';

class App extends Component {

  constructor(props, context) {
    super(props, context);

    this.handleLogout = this
      .handleLogout
      .bind(this);
  }

  handleLogout() {
    this
      .props
      .actions
      .logOut();
  }

  render() {
    const {loading, isAuthenticated, children} = this.props;

    return (
      <div className="container-fluid">
        <Header
          loading={loading}
          isAuthenticated={isAuthenticated}
          handleLogout={this.handleLogout}/> {children}
      </div>
    );
  }
}

App.propTypes = {
  children: PropTypes.object.isRequired,
  loading: PropTypes.bool.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
  actions: PropTypes.object.isRequired
};

function mapStateToProps(state, ownProps) {
  return {
    loading: state.ajaxCallsInProgress > 0,
    isAuthenticated: state.auth.isAuthenticated
  };
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(authActions, dispatch)
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(App);
