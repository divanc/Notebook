# Reducer Composition With Arrays

We faced a problem, that our reducer traces both whole arrays of todos and separate ones, we can solve it having one extra functions

```js
const todo = (state, action) => {
  switch (action.type) {
    case "ADD_TODO":
      return {
        id: action.id,
        text: action.text,
        isCompleted: false
      };

    case "TOGGLE_TODO":
      if (state.id !== action.id) return state;
      return {
        ...state,
        completed: !state.completed
      };
  }
};
```

Now we can call this reducer in our first reducer:

```js
const todos = (state = [], action) => {
  switch (action.type) {
    case "ADD_TODO":
      return [...state, todo(undefined, action)];

    case "TOGGLE_TODO":
      return state.map(t => todo(t, action));

    default:
      return state;
  }
};
```

## Composition with Objects

What if we want control more advanced options? For ex, user may choose what kind of tasks he wants to see:

```js
const visibilityFilter = (state = "SHOW_ALL", action) => {
  switch (action.type) {
    case "SET_VISIBILIty_FILTER":
      return action.filter;
    default:
      return state;
  }
};
```

Instead of modifying our reducers we can create new, which will contain all of them, as they do different things:

```js
const todoApp = (state = {}, action) => {
  return {
    todos: todos(state.todos, action),
    visibilityFilter: visibilityFilter(state.visibilityFilter, action)
  };
};
```

Now having papa-reducer we would use it to create our store

```js
const store = createStore(todoApp);
```

### Actually, there is a `combineReducers()`

In fact, what we just did already is in Redux:

```js
import { combineReducers } from "redux";
const todoApp = combineReducers({
  todos: todos, //stateTree name and reducer
  visibilityFilter: visibilityFilter
});
```

Also ES6 provides a shorthand literal notation, so we can do this:

```js
import { combineReducers } from "redux";
const todoApp = combineReducers({
  todos,
  visibilityFilter
});
```

As they are the same

## `combineReducer()` From Scratch

```js
const combineReducers = reducers => {
  //We know it returns a reducer
  return (state = {}, action) => {};
};
```

```js
const combineReducers = reducers => {
  //We know it returns a reducer
  return (state = {}, action) => {
    //Object.keys would return keys of our object; In our case: `todos` and `visibilityFilter`
    return Object.keys(reducers).reduce(
      // reduce will affect only one value among, such as `nextState`
      (nextState, key) => {
        //So nextState[key] is `isCompleted` of [todos] for example
        //To change it we would call `todos` reducer === reducers[key]() and pass it arguments
        nextState[key] = reducers[key](state[key], action);

        return nextState;
      },
      {}
    );
  };
};
```

And it's all

Recap:

We call `combineReducers` with args of reducers. It is a little bit complex, mates.
