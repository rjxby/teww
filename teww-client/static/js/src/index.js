/* eslint-disable import/default */
import 'babel-polyfill';
import React from 'react';
import {render} from 'react-dom';
import configureStore from './store/configureStore';
import {Provider} from 'react-redux';
import {Router, hashHistory} from 'react-router';
import routes from './routes';
import {checkAuth} from './actions/authActions';
import '../../css/style.css';
import '../../../node_modules/bootstrap/dist/css/bootstrap.min.css';
import '../../../node_modules/bootstrap-material-design/dist/css/bootstrap-material-design.css';
import '../../../node_modules/bootstrap-material-design/dist/css/ripples.min.css';

const store = configureStore();
store.dispatch(checkAuth());

render(
  <Provider store={store}>
    <Router history={hashHistory} routes={routes} />
  </Provider>,
  document.getElementById('app')
);
