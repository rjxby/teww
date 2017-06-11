import React from 'react';
import PropTypes from 'prop-types';
import TextInput from '../common/TextInput';

const LoginForm = ({onSave, onChange, saving, errors}) => {
    return (
        <form>
            <h1>Login</h1>
            <TextInput
                name="username"
                label="Username / Email"
                onChange={onChange}
                error={errors.username}/>

            <div className="form-group">
                <label htmlFor="password">Password</label>
                <div className="field">
                    <input
                        type="password"
                        name="password"
                        className="form-control"
                        onChange={onChange}/>
                </div>
            </div>

            <input
                type="submit"
                disabled={saving}
                value={saving ? 'Login...' : 'Login'}
                className="btn btn-primary"
                onClick={onSave}/>
        </form>
    );
};

LoginForm.propTypes = {
    onSave: PropTypes.func.isRequired,
    onChange: PropTypes.func.isRequired,
    saving: PropTypes.bool,  
    errors: PropTypes.object     
};

export default LoginForm;
