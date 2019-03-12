# Refacturing The Refactured Todo!

## FilterLink & Footer

We have to pass a lot of components in the tree :c

Check it out yourself:

```jsx
export const FilterLink = ({ filter, children, currentFilter, onClick }) => {
  if (currentFilter === filter) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick(filter);
      }}
    >
      {children}
    </a>
  );
};
```

And father-`Footer` needs to know all that and his papa needs to know everything was mentioned, kinda a lot.

## Let's add more containers!

Cut all props from the `Footer`:

```jsx
export const Footer = () => (
  <p>
    Show: {"   "}
    <FilterLink filter="SHOW_ALL">All</FilterLink> {"  "}
    <FilterLink filter="SHOW_ACTIVE">Active</FilterLink> {"  "}
    <FilterLink filter="SHOW_COMPLETED">Completed</FilterLink> {"  "}
  </p>
);
```

And our `FilterLink` to be honest not as presentational as it could be:

It was:

```jsx
export const FilterLink = ({ filter, children, currentFilter, onClick }) => {
  if (currentFilter === filter) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick(filter);
      }}
    >
      {children}
    </a>
  );
};
```

Let's cut it's functions and make it just `Link`:

```jsx
export const Link = ({ active, children, onClick }) => {
  if (active) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick();
      }}
    >
      {children}
    </a>
  );
};
```

Well, it looks presentational as hell

However, container for link would be more complex:

```jsx
export class FilterLink extends Component {
  componentDidMount() {
    this.unsubscribe = store.subscribe(() => this.forceUpdate());
  }

  componentWillUnmount() {
    this.unsubscribe();
  }
  render() {
    const props = this.props;
    const state = store.getState();

    return (
      <Link
        active={props.filter === state.visibilityFilter}
        onClick={() => {
          store.dispatch({
            type: "SET_VISIBILITY_FILTER",
            filter: props.filter
          });
        }}
      >
        {props.children}
      </Link>
    );
  }
}
```

## VisibleTodoList

That is a container for `TodoList`:

```jsx
export class VisibleTodoList extends Component {
  componentDidMount() {
    this.unsubscribe = store.subscribe(() => this.forceUpdate());
  }

  componentWillUnmount() {
    this.unsubscribe();
  }
  render() {
    const props = this.props;
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
```

So now before footer it would be enough to do:

```jsx
<VisibleTodoList />
```

## AddTodo

And backtrack `AddTodo`:

```jsx
let nextTodoId = 0;

export const AddTodo = () => {
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
```

# CHECK THIS OUT:

Now our main `TodoApp` is going to look like this:

```jsx
import React from "react";

import { AddTodo } from "./AddTodo";
import { VisibleTodoList } from "./VisibleTodoList";
import { Footer } from "./Footer";

export const TodoApp = () => (
  <div>
    <AddTodo />
    <VisibleTodoList />
    <Footer />
  </div>
);
```

DO YOU SEE THAT?

Also `index`:

```jsx
import React from "react";
import ReactDOM from "react-dom";

import { TodoApp } from "./components/TodoApp";

ReactDOM.render(<TodoApp />, document.getElementById("app"));
```
