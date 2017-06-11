import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as itemActions from '../../actions/itemActions';
import ItemForm from './ItemForm';
import {tagsFormattedForDropdown} from '../../selectors/selectors';

class ManageItemPage extends Component {
  constructor(props, context){
    super(props, context);

    this.state = {
        item: Object.assign({}, this.props.item),
        errors: {},
        processing: false
    };

    this.updateItemState = this.updateItemState.bind(this);
    this.saveItem = this.saveItem.bind(this);
    this.deleteItem = this.deleteItem.bind(this);
    this.cancelItem = this.cancelItem.bind(this);
  }

  componentWillReceiveProps(nextProps){
      if(this.props.item.id !== nextProps.item.id){
        this.setState({item: Object.assign({}, nextProps.item)});  
      }
  }

  updateItemState(event) {
      const field = event.target.name;
      let item = this.state.item;
      item[field] = event.target.value;
      return this.setState({item: item});
  }

  itemFormIsValid(){
      let formIsValid = true;
      let errors = {};

    //   if(this.state.item.title.length < 5){ TODO
    //       errors.title = 'Titile must be at least 5 charachters';
    //       formIsValid = false;
    //   }

      this.setState({errors: errors});
      return formIsValid;
  }

  saveItem(event){
      event.preventDefault();

      if(!this.itemFormIsValid()){
          return;
      }

      this.setState({processing: true});
      this.props.actions.saveItem(this.state.item)
        .then(() => this.redirect('Success saved'))
        .catch(error => {
            alert(error);
            this.setState({processing: false});
        });
  }

  deleteItem(event){
      event.preventDefault();

      this.setState({processing: true});
      this.props.actions.removeItem(this.state.item.id)
        .then(() => this.redirect('Success removed'))
        .catch(error => {
            alert(error);
            this.setState({processing: false});
        });
  }

  cancelItem(event){
      event.preventDefault();
      //todo reset validation
      this.setState({item: {}, errors: {}});
      this.context.router.push('/items');
  }

  redirect(message){
      this.setState({processing: false});
      alert(message);
      this.context.router.push('/items');
  }

  render() {
    return(
        <ItemForm 
            tags={this.props.tags}
            onChange={this.updateItemState}
            onSave={this.saveItem}
            onDelete={this.deleteItem}
            onCancel={this.cancelItem}
            item={this.state.item}
            errors={this.state.errors}
            processing={this.state.processing}
        />
    );
  }
}

ManageItemPage.propTypes = {
    item: PropTypes.object.isRequired,
    tags: PropTypes.array.isRequired,
    actions: PropTypes.object.isRequired
};

ManageItemPage.contextTypes = {
    router: PropTypes.object
};

function getItemById(items, id){
    const item = items.filter(item => item.id === id);
    if(item) return item[0];
    return null;
}

function mapStateToProps(state, ownProps){
    const itemId = ownProps.params.id;

    let item = {id: '', dateStart: '', dateEnd: '', length:'', description: '', tags: ''};

    if(itemId && state.items.length > 0){
        item = getItemById(state.items, itemId);
    }

    return {
        item: item,
        tags: tagsFormattedForDropdown(state.tags)
    };
}

function mapDispatchToProps(dispatch){
    return {
        actions: bindActionCreators(itemActions, dispatch)
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(ManageItemPage);