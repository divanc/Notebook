# Reducer

State mutations should be described as a pure functions.

It has to be pure, hence, has to return a new object.

**Third dogma: Redux takes the previous state tree and action, dispatches action and returns a new tree**

This function, which takes prev and action and makes a new one is calledd the _reducer_.

## Let's write one!

1. Reducer takes state and action as args and returns the next state:

```js
function reducer(state, action) {
  return state;
}
```

2. The problem is `action` hasn't occur yet, so it doesn't understand us. Let's fix it:

```js
function reducer(state, action) {
  if(action.type === "INCREMENT") return state + 1;
  else if action.type === "DECREMENT") return state -1;
}
```

3. But what if it will call for some unknown action? `reducer` would throw an error again! :(

This can be easily fixed:

```js
function reducer(state, action) {
  if (action.type === "INCREMENT") return state + 1;
  else if (action.type === "DECREMENT") return state - 1;
  else return state;
}
```

4. All seemed to be good, yet what if `state` is undefined? Another error? Well, let's handle that:

```js
function reducer(state, action) {
  if (typeof state === "undefined") return 0; // It will set initial state

  if (action.type === "INCREMENT") return state + 1;
  else if (action.type === "DECREMENT") return state - 1;
  else return state;
}
```

5. Let's make it look prettier a bit:

```js
const reducer = (state = 0, action) => {
  switch (action.type) {
    case "INCREMENT":
      return state + 1;
    case "DECREMENT":
      return state - 1;
    default:
      return state;
  }
};
```
