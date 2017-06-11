import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as itemActions from '../../actions/itemActions';
import * as tagActions from '../../actions/tagActions';
import ItemList from './ItemList';
import {hashHistory} from 'react-router';

class ItemsPage extends Component {
  constructor(props, context) {
    super(props, context);
    this.redirectToAddItemPage = this
      .redirectToAddItemPage
      .bind(this);
  }

  componentWillMount() {
    this
      .props
      .tagsActions
      .loadTags().then(() => {
        this
          .props
          .itemsActions
          .loadItems();
      });
  }

  itemRow(item, index) {
    return <div key={index}>{item.id}</div>;
  }

  redirectToAddItemPage() {
    hashHistory.push('/item');
  }

  render() {
    const {items} = this.props;

    return (
      <div>
        <h1>Items</h1>
        <input
          type="submit"
          value="Add item"
          className="btn btn-primary"
          onClick={this.redirectToAddItemPage}/>
        <ItemList items={items}/>
      </div>
    );
  }
}

ItemsPage.propTypes = {
  items: PropTypes.array.isRequired,
  itemsActions: PropTypes.object.isRequired,
  tagsActions: PropTypes.object.isRequired
};

function mapStateToProps(state, ownProps) {
  return {items: state.items};
}

function mapDispatchToProps(dispatch) {
  return {
    itemsActions: bindActionCreators(itemActions, dispatch),
    tagsActions: bindActionCreators(tagActions, dispatch)
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(ItemsPage);
