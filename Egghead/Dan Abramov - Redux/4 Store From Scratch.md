| [< Previous abstract](3%20Store%20Methods.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](5%20React%20Counter.md) |
| --------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | ----------------------------------------- |


# Implemeting Store From Scratch

Let's build our own Redux from scratch, you know, just for fun.

Instead of `const store = createStore(reducer);` we will write our own. We know arg is reducer...

```js
const createStore = reducer => {
  //We also know it keeps our state
  let state;

  // We know it has 3 methods
  const getState = () => state;

  const dispatch = action => {};

  const subscribe = listener => {};

  // And this is what we call a Redux store
  return { getState, dispatch, subscribe };
};
```

Because we can call subcribe several times, we need to keep out listeners:

```js
const createStore = reducer => {
  let state;

  let listeners = [];

  const getState = () => state;

  const dispatch = action => {
    //This is how we add a state
    state = reducer(state, action);
    //We should also notify each listener about the change
    listeners.forEach(listener => listener());
  };

  const subscribe = listener => {
    //Every time it is called we want to push our new listener into the array
    listeners.push(listener);
  };

  return { getState, dispatch, subscribe };
};
```

We would also like to add unsubcscribe method right in the `subscribe`:

```js
const subscribe = listener => {
  listeners.push(listener);

  return () => {
    listeners = listeners.filter(l => l !== listener);
  };
};
```

Finally we would like to return initial value, so we just `dispatch` before _store_ is returned:

```js
dispatch({});
return { getState, dispatch, subscribe };
```
