import React from 'react';
import PropTypes from 'prop-types';
import {Link} from 'react-router';

const ItemListRow = ({item}) => {
    return (
        <tr>
            <td><Link to={'/item/' + item.id}>{item.dateStart} - {item.dateEnd}</Link></td>
            <td>{item.length}</td>
            <td>{item.description}</td>
            <td>{item.tagId}</td>
        </tr>
    );
};

ItemListRow.propTypes = {
    item: PropTypes.object.isRequired
};

export default ItemListRow;