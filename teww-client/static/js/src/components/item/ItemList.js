import React from 'react';
import PropTypes from 'prop-types';
import ItemListRow from './ItemListRow';

const ItemList = ({items}) => {
    return (
        <table className="table">
            <thead>
                <tr>
                    <th>&nbsp;</th>
                    <th>Length</th>
                    <th>Description</th>
                    <th>Tags</th>
                </tr>
            </thead>
            <tbody>
                {
                    items.map(item => 
                    <ItemListRow key={item.id} item={item} />
                    )
                }
            </tbody>
        </table>
    );
};

ItemList.propTypes = {
    items: PropTypes.array.isRequired
};

export default ItemList;