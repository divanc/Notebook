# Redux Crash Course

[this](https://www.youtube.com/watch?v=93p3LxR9xfM) big'o

In React we have states, which are uncomfortable to manage among heavy apps. It is state manager.

State is immutable, it comes from top to down

Let's create basic fetch class:

```js
class Posts extends Component {
  componentWillMount() {
    fetch('https://jsonplaceholder.typicode.com/posts')
    .then(res => res.json())
    .then(data => console.log(data));
  }
  ...
```

And display posts in a typical React-way:

```js
class Posts extends Component {
  constructor(props) {
    super(props);
    this.state = {
      posts: []
    }
  }
  componentWillMount() {
    fetch('https://jsonplaceholder.typicode.com/posts')
    .then(res => res.json())
    .then(data => this.setState({posts: data}));
  }

  render() {
    const postItems =  this.state.posts.map(post => (
      <div key={post.id}>
        <h3>  {post.title} </h3>
        <p>  {post.body}  </p>
      </div>
    ));
    return (
      <div>
        <h1>  Posts  </h1>
        {postItems}
      </div>
    );
  }
}
```

Now let's create a component to add posts. Pretty typical: 

```js
export default class Postform extends Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      body: ''
    };
  }

  onTitleChange = event => this.setState({title: event.target.value});

  onBodyChange = event => this.setState({body: event.target.value});

  render() {
    return (
      <div>
        <form>
        <div>
          <label>Title: </label><br />
          <input type="text" name="title" value={this.state.title} onChange={this.onTitleChange}/>
        </div>
        <div>
          <label> Body: </label><br />
          <input type="text" name="body" value={this.state.body} onChange={this.onBodyChange}/>
        </div>
        <button type='submit'>Submit</button>
        </form>
      </div>
    );
  }
}
```

And add typical `onSubmit`:

```js
  onSubmit = event => {
    event.preventDefault();

    const post = {
      title: this.state.title,
      body: this.state.body
    }

    fetch('https://jsonplaceholder.typicode.com/posts', {
      method: 'POST',
      headers: {
        'content-type': 'application/json'
      },
      body: JSON.stringify(post)
    })
    .then(res => res.json())
    .then(data => console.log(data));
  }
  ```

## Now let's talk Redux!

Install redux:

```console
npm i redux react-redux redux-thunk
```

Let's bring in the *provider*. Provider is a component functions as a glue for React & Redux:

```js
import { Provider } from 'react-redux';
```

and wrap all the `App`'s return into that `Provider`:

```js
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <Postform />
            <hr />
            <Posts />
          </header>
        </div>
      </Provider>
    );
  }
}
```

  On error starting a server, just run `npm install`

Provider provides with a *store*. In order to create it we need `createStore(reducer, [preloadedState], [enhancer])` with reducers in.

Before class write in:

```js
import { createStore, applyMiddleware} from 'redux';

const store = createStore(() => [],{},applyMiddleware());
```

Now it at least can run. Another way is to create a js file in src and do redux magic there:

`store.js`

```js
import { createStore, applyMiddleware} from 'redux';

const store = createStore(() => [],{},applyMiddleware());

export default store;
```

Now lets create folder `reducers` in src with `index.js` and make a typical `store.js` look like:

```js
import { createStore, applyMiddleware} from 'redux';
import  thunk from 'redux-thunk';
import rootReducer from './reducers';

const initialState = {};

const middleware = [thunk];

const store = createStore(
  rootReducer, 
  initialState, 
  applyMiddleware(...middleware)
  );

export default store;
```

In `reducers/index.js` we gonna fill in:

```js
import { combineReducers } from 'redux';
import postReducer from './postReducer';

export default  combineReducers({
  posts: postReducer
});
```

and `postReducer` is our gem.

## postReducer

Create `actions/types.js` in src

Here we can insert all actions we want to use:

```js
export const FETCH_POSTS = 'FETCH_POSTS';
export const NEW_POST = 'NEW_POST';
```

And in `reducers/postReducer` (yes) yet another codewall:

```js
import { FETCH_POSTS, NEW_POST } from '../actions/types';

const initialState = {
  items: [],
  item: {}
}

export default function(state = initialState, action) {
  switch(action.type) {
    default:
      return state;
  }
}
```

Woopsy, not for now. Now — another file: `actions/postActions`:

```js
import { FETCH_POSTS, NEW_POST } from './types';

export function fetchPosts() {
  return function(dispatch) {
  }
}
```

Where dispatch function is kinda like promise. 

Now we can take fetch from `Posts` Component; actually we may remove `ComponentWillMount` and `constructor` cuz state is 2015. Now it will look like:


```js
class Posts extends Component {
  render() {
    const postItems =  this.state.posts.map(post => (
      <div key={post.id}>
        <h3>  {post.title} </h3>
        <p>  {post.body}  </p>
      </div>
    ));
    return (
      <div>
        <h1>  Posts  </h1>
        {postItems}
      </div>
    );
  }
}

export default Posts;
```

However in `postActions` we don't want to just `setState` of a post, we want to dispatch it to the reducers to the store:

```js
import { FETCH_POSTS, NEW_POST } from './actions/types';

export function fetchPosts() {
  return function(dispatch) {
    fetch('https://jsonplaceholder.typicode.com/posts')
    .then(res => res.json())
    .then(posts => dispatch({
      type: FETCH_POSTS,
      payload: posts
    }));
  }
}
```

and in `reducers/postReducer` we just add:

```js
export default function(state = initialState, action) {
  switch(action.type) {
    case FETCH_POSTS:
      return {
        ...state,
        items: action.payload
      }
    default:
      return state;
  }
}
```

Finally, we add redux into `Posts`:

```js
import { connect } from 'react-redux';
import {fetchPosts } from '../actions/postActions';
```

instead of `export default Posts;` we want to use:

```js
class Posts extends Component {
  componentWillMount() {
    this.props.fetchPosts();
  }
  render() {
    const postItems =  this.state.posts.map(post => (
      <div key={post.id}>
        <h3>  {post.title} </h3>
        <p>  {post.body}  </p>
      </div>
    ));
    return (
      <div>
        <h1>  Posts  </h1>
        {postItems}
      </div>
    );
  }
}

export default connect(null, { fetchPosts})(Posts);
```

and now we want to match state to props, as we no longer have `this.state.posts`.

This function can realocate state to props:

```js
const mapStateToProps = state => ({
  posts: state.posts.items
})

export default connect(mapStateToProps, { fetchPosts})(Posts);
```

# Hooray, now we have post submitting feature

Now we add in `Posts`:

```js
import PropTypes from 'prop-types';
//...
//after Posts class;
//...

Posts.propTypes = {
  fetchPosts: PropTypes.func.isRequired,
  posts:PropTypes.array.isRequired
}

```

Dunno what this does

## Redux Dev Tools

If it says `no store found` (and it will):

In `store` import yet one more thing:

```js
import { createStore, applyMiddleware, compose} from 'redux';
```

And wrap compose there. It let us use several enhancers:

```js

const store = createStore(
  rootReducer, 
  initialState, 
  compose (
  applyMiddleware(...middleware),
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__())
  );
```

Now in browser we can see that we have

posts.items which store all 100 posts there. The names are from `reducers/index` and `reducers/postReducer`

## Implement new posts

In `actions/postActions` lets moderate `createPost`, similar to the react version:

```js
export const createPost = (postData) => dispatch => {
  fetch('https://jsonplaceholder.typicode.com/posts', {
    method: 'POST',
    headers: {
      'content-type': 'application/json'
    },
    body: JSON.stringify(postData)
  })
  .then(res => res.json())
  .then(post => dispatch({
    type: NEW_POST,
    payload: post
  }));
}
```

`type: NEW_POST` would call to postReducer, let's handle it

```js
    case NEW_POST:
      return {
        //Since JSON Placeholder isn't actually adding posts, it gonna work in a strange way
        ...state,
        item: action.payload
      }

```

And redux out the `Postform`:

```js
import { connect } from 'react-redux';
import { PropTypes } from 'prop-types';
import {createPost } from '../actions/postActions';

class Postform extends Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      body: ''
    };
  }

  onChange = event => this.setState({[event.target.name]: event.target.value});
 
  onSubmit = event => {
    event.preventDefault();

    const post = {
      title: this.state.title,
      body: this.state.body
    }
    this.props.createPost(post);

  }

  render() {
    return (
      <div>
        <form onSubmit={this.onSubmit}>
        <div>
          <label>Title: </label><br />
          <input type="text" name="title" value={this.state.title} onChange={this.onChange}/>
        </div>
        <div>
          <label> Body: </label><br />
          <input type="text" name="body" value={this.state.body} onChange={this.onChange}/>
        </div>
        <button type='submit'>Submit</button>
        </form>
      </div>
    );
  }
}
Postform.propTypes = {
  createPost: PropTypes.func.isRequired
}

export default connect(null,  { createPost })(Postform);
```

Now we have not only `items`  in our browser, but also an `item` : the state of a new post.


## Summing

Now we want to sum it up.

To have a single source of truth, lets modify in `Posts`:

```js
Posts.propTypes = {
  fetchPosts: PropTypes.func.isRequired,
  posts:PropTypes.array.isRequired,
  newPost: PropTypes.object
}

const mapStateToProps = state => ({
  posts: state.posts.items,
  newPost: state.posts.item
})

```

And to dynamically add it to the page there is a React method called `componentWillRecieveProps()`:

```js
  componentWillReceiveProps(nextProps) {
    if (nextProps.newPost) {
      this.props.posts.unshift(nextProps.newPost);
    }
  }

```

#TADA