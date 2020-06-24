import React from 'react';
import postCodeData from '../helpers/PostCodeData.js'

// load postcode data
const PostCodeContext = React.createContext(postCodeData);
export default PostCodeContext;