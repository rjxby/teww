import React from 'react';
import {Route, IndexRoute} from 'react-router';
import App from './components/App';
import NotFound from './components/NotFound';
import ManageLoginPage from './components/auth/ManageLoginPage';
import HomePage from './components/home/HomePage';
import AboutPage from './components/about/AboutPage';
import ItemsPage from './components/item/ItemsPage';
import ManageItemPage from './components/item/ManageItemPage';
import requireAuth from './components/RequireAuth';

const Routes = (
  <Route path="/" component={App}>
    <IndexRoute component={HomePage} />
    <Route path="login" component={ManageLoginPage} />
    <Route path="signup" component={NotFound} />
    <Route path="items" component={requireAuth(ItemsPage)} />
    <Route path="item" component={requireAuth(ManageItemPage)} />
    <Route path="item/:id" component={requireAuth(ManageItemPage)} />
    <Route path="about" component={AboutPage} />
    <Route path="*" component={NotFound} />
  </Route>
);

export default Routes;
