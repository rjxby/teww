import React from 'react';
import PropTypes from 'prop-types';
import TextInput from '../common/TextInput';
import SelectInput from '../common/SelectInput';

const ItemForm = ({item, tags, onSave, onDelete, onCancel, onChange, processing, errors}) => {
    return (
        <form>
            <h1>Manage item</h1>
            <TextInput
                name="dateStart"
                label="Begin date"
                value={item.dateStart}
                onChange={onChange}
                error={errors.dateStart}/>

            <TextInput
                name="dateEnd"
                label="End date"
                value={item.dateEnd}
                onChange={onChange}
                error={errors.dateEnd}/>

            <TextInput
                name="description"
                label="Description"
                value={item.description}
                onChange={onChange}
                error={errors.description}/>

            <SelectInput
                name="tagId"
                label="Tag"
                value={item.tagId}
                defaultOption="Select tag"
                options={tags}
                onChange={onChange}
                error={errors.tagId}/>

            <input
                type="submit"
                disabled={processing}
                value={processing ? 'Saving...' : 'Save'}
                className="btn btn-primary"
                onClick={onSave}/>

            <input
                type="button"
                disabled={processing}
                value="Cancel"
                className="btn btn-danger"
                onClick={onCancel}/>

            <input
                type="button"
                disabled={processing}
                value="Delete"
                className={!item.id ? "hidden" : "btn btn-warning"} 
                onClick={onDelete}/>
        </form>
    );
};

ItemForm.propTypes = {
    item: PropTypes.object.isRequired,
    tags: PropTypes.array,
    onSave: PropTypes.func.isRequired,
    onDelete: PropTypes.func.isRequired,
    onCancel: PropTypes.func.isRequired,
    onChange: PropTypes.func.isRequired,
    processing: PropTypes.bool,  
    errors: PropTypes.object     
};

export default ItemForm;