## Redux Store Methods

```js
import { createStore } from "redux";
```

Store binds all three Redux dogmas in it already! When you create one, you need to specify the reducer:

```js
const store = createStore(reducer);
```

Store has 3 importants methods...

### `getState()`

```js
console.log(store.getState());
// Would return 0, as it is initial in our reducer
```

### `dispatch()`

```js
store.dispatch({ type: "INCREMENT" });
console.log(store.getState());
// Now 1!!
```

### `subscribe()`

```js
store.subscribe(() => {
  document.body.innerText = store.getState();
});

document.addEventListener("click", () => {
  store.dispatch({ type: "INCREMENT" });
});
```

But it doesn't render the initial state. That's because we render only on action. Let's fix it that way: make a separate function of `render` and then call it once beyond action:

```js
const render = () => {
  document.body.innerText = store.getState();
};

store.subscribe(render);
render();
```

Now we have our counter app!
