import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as authActions from '../../actions/authActions';
import LoginForm from './LoginForm';

class ManageLoginPage extends Component {
  constructor(props, context){
    super(props, context);

    this.state = {
        auth: {},
        errors: {},
        saving: false
    };

    this.updateAuthState = this.updateAuthState.bind(this);
    this.onAuthentication = this.onAuthentication.bind(this);
  }

  updateAuthState(event) {
        const field = event.target.name;
        let auth = this.state.auth;
        auth[field] = event.target.value;
        return this.setState({auth: auth});
    }

  authFormIsValid(){
      let formIsValid = true;
      let errors = {};

      if(this.state.auth.username.length < 3){
          errors.username = 'Username must be at least 3 charachters';
          formIsValid = false;
      }

      this.setState({errors: errors});
      return formIsValid;
  }

  onAuthentication(event){
      event.preventDefault();

      if(!this.authFormIsValid()){
          return;
      }

      this.setState({saving: true});
      this.props.actions.onAuthentication(this.state.auth)
        .then(() => this.redirect())
        .catch(error => {
            alert(error);
            this.setState({saving: false});
        });
  }

  redirect(){
      this.setState({saving: false});
      this.context.router.push('/');
  }

  render() {
    return (
      <div className="row">
        <div className="col-md-4 col-md-offset-4">
          <LoginForm 
            onChange={this.updateAuthState}
            onSave={this.onAuthentication}
            errors={this.state.errors}
            saving={this.state.saving}
          />
        </div>
      </div>
    );
  }
}

ManageLoginPage.propTypes = {
    auth: PropTypes.object,
    actions: PropTypes.object.isRequired
};

ManageLoginPage.contextTypes = {
    router: PropTypes.object
};

function mapDispatchToProps(dispatch){
    return {
        actions: bindActionCreators(authActions, dispatch)
    };
}

export default connect(null, mapDispatchToProps)(ManageLoginPage);
