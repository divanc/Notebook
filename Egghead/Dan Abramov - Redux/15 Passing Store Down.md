| [< Previous abstract](14%20Refactoring%202.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](16%20connect.md) |
| ---------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | ---------------------------------- |


# Passing The Store Down Explicitly

Previously we passed the store to components, so they can subscribe to, dispatch on and get state from it.

However, working with multiple files is complex and just passing isn't appropriate as export.

## via Props

Instead of just creating store somewhere, we would pass it as a prop to the core component:

```jsx
ReactDOM.render(
  <TodoApp store={createStore(todoApp)} />,
  document.getElementById("app")
);
```

Hence, every component needs a reference to the store

```jsx
export const TodoApp = ({ store }) => (
  <div>
    <AddTodo store={store} />
    <VisibleTodoList store={store} />
    <Footer store={store} />
  </div>
);
```

No man, this is pure madness

## via Context

We can create context component `Provider`:

```jsx
export class Provider extends Component {
  render() {
    return this.props.children;
  }
}
```

We gonna wrap around all render:

```jsx
ReactDOM.render(
  <Provider store={createStore(todoApp)}>
    <TodoApp />
  </Provider>,
  document.getElementById("app")
);
```

Now to pass it all down we have some magic:

```jsx
export class Provider extends Component {
  // Some new React magic
  getChildrenContext() {
    return { store: this.props.store };
  }

  render() {
    return this.props.children;
  }
}
```

Now we have to add this to every component:

```jsx
Provider.childContextTypes = {
  store: React.PropTypes.object
};
```

and call store via `this.context`

```jsx
import { Component } from "react";
import { TodoList } from "./TodoList";
import { getVisibleTodos } from "../Redux";
import PropTypes from "prop-types";

export class VisibleTodoList extends Component {
  componentDidMount() {
    const { store } = this.context;
    this.unsubscribe = store.subscribe(() => this.forceUpdate());
  }

  componentWillUnmount() {
    this.unsubscribe();
  }
  render() {
    const props = this.props;
    const { store } = this.context;
    const state = store.getState();

    return (
      <TodoList
        todos={getVisibleTodos(state.todos, state.visibilityFilter)}
        onTodoClick={id =>
          store.dispatch({
            type: "TOGGLE_TODO",
            id
          })
        }
      />
    );
  }
}

VisibleTodoList.contextTypes = {
  store: PropTypes.object
};
```

And for functional components `context` is passed as a second prop after `props`:

```jsx
let nextTodoId = 0;

export const AddTodo = (props, { store }) => {
  let input;
  const ref = useRef(null);

  return (
    <div>
      <input ref={r => (ref.current = r)} />

      <button
        onClick={() => {
          store.dispatch({
            type: "ADD_TODO",
            id: nextTodoId++,
            text: ref.current.value
          });
          ref.current.value = "";
        }}
      >
        Add Todo
      </button>
    </div>
  );
};

AddTodo.contextTypes = {
  store: PropTypes.object
};
```

# You guess it, we have React-Redux Provider built-in

```jsx
import { Provider } from "react-redux";
```

And you've nailed the component
