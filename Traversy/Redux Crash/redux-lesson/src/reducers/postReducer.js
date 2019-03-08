import { FETCH_POSTS, NEW_POST } from '../actions/types';

const initialState = {
  items: [],
  item: {}
}

export default function(state = initialState, action) {
  switch(action.type) {
    case FETCH_POSTS:
      return {
        ...state,
        items: action.payload
      }
    
    case NEW_POST:
      return {
        //Since JSON Placeholder isn't actually adding posts, it gonna work in a strange way
        ...state,
        item: action.payload
      }
    default:
      return state;
  }
}