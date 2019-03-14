| [< Previous abstract](4%20Store%20From%20Scratch.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](6%20Avoiding%20Array%20Mutations.md) |
| ---------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ |


# React Counter Example

Earlier we did:

```js
store.subscribe(() => {
  document.body.innerText = store.getState();
});
```

We don't want to do this in big apps. At least let it be:

```js
const render = () => {
  ReactDOM.render(
    <Counter value={store.getState()} />,
    document.getElementById("root")
  );
};
```

```js
const Counter = ({ value }) => {
  <h1>{value}</h1>;
};
```

Now let's add variativity:

```js
const Counter = ({ value, onIncrement, onDecrement }) => {
  <div>
    <h1>{value}</h1>

    <button onClick={onIncrement}>+</button>
    <button onClick={onDecrement}>-</button>
  </div>;
};
```

```js
const render = () => {
  ReactDOM.render(
    <Counter
      value={store.getState()}
      onIncrement={() => store.dispatch({ type: "INCREMENT" })}
      onDecrement={() => store.dispatch({ type: "DECREMENT" })}
    />,
    document.getElementById("root")
  );
};
```

Let's recap:

`Counter` component is a _dumb_ component. It does not contain business logic. It just builds what is to render.
`store.subscribe(render);` keeps state fresh.
