export function tagsFormattedForDropdown(tags){
    return tags.map(tag => {
        return {
            value: tag.id,
            text: tag.name
        };
    });
}