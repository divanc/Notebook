| [< Previous abstract](9%20Todo%20Composition.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](11%20Todo%20React%202.md) |
| ------------------------------------------------ | --------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |


# React Todo List Example: Adding a todo

```js
const render = () => {
  ReactDOM.render(<TodoApp />, document.getElementById("root"));
};

store.subscribe(render);
render();
```

```js
let nextTodoId = 0;

class TodoApp extends Component {
  render() {
    return (
      <div>
        <button
          onClick={() => {
            store.dispatch({
              type: "ADD_TODO",
              text: "Test",
              id: nextTodoId++
            });
          }}
        >
          Add Todo
        </button>

        <ul>
          {this.props.todos.map(todo => (
            <li key={todo.id}>{todo.text}</li>
          ))}
        </ul>
      </div>
    );
  }
}
```

Also connect todos in this `TodoApp`:

```js
const render = () => {
  ReactDOM.render(
    <TodoApp todos={store.getState().todo} />,
    document.getElementById("root")
  );
};
```

[Here it is](https://codesandbox.io/s/o5j36jokv5)

Now, onclick we see adding "Test" todos, let's add input

```js
...
    return (
      <div>
        <input
          ref={node => {
            this.input = node;
          }}
        />
        <button
          onClick={() => {
            store.dispatch({
              type: "ADD_TODO",
              text: this.input.value,
              id: nextTodoId++
            });
            this.input.value = "";
          }}
        >
          Add Todo
        </button>
...
```

# Review

All the code looks like this:

```js
import React, { Component } from "react";
import ReactDOM from "react-dom";
import { createStore, combineReducers } from "redux";

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

const visibilityFilter = (state = "SHOW_ALL", action) => {
  switch (action.type) {
    case "SET_VISIBILIty_FILTER":
      return action.filter;
    default:
      return state;
  }
};

const todoApp = combineReducers({
  todos,
  visibilityFilter
});

const store = createStore(todoApp);

let nextTodoId = 0;

class TodoApp extends Component {
  render() {
    return (
      <div>
        <input
          ref={node => {
            this.input = node;
          }}
        />
        <button
          onClick={() => {
            store.dispatch({
              type: "ADD_TODO",
              text: this.input.value,
              id: nextTodoId++
            });
            this.input.value = "";
          }}
        >
          Add Todo
        </button>

        <ul>
          {this.props.todos.map(todo => (
            <li key={todo.id}>{todo.text}</li>
          ))}
        </ul>
      </div>
    );
  }
}

const render = () => {
  ReactDOM.render(
    <TodoApp todos={store.getState().todos} />,
    document.getElementById("app")
  );
};

store.subscribe(render);
render();
```

## Depricated

The way Dan calls `ref` is depricated already :C

Contemporary decision would be hooks. Instead of class `TodoApp` let's make a function:

```js
import React, { Component, useRef, createRef } from "react";
```

But, hey, nay. It is too complex for me now. Let's use normal state:

```js
class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = { value: "" };

    this.onChange = this.onChange.bind(this);
  }

  onChange(event) {
    this.setState({ value: event.target.value });
  }

  render() {
    return (
      <div>
        <input value={this.state.value} onChange={this.onChange} />
        <button
          onClick={() => {
            console.log(this.state.value);
            store.dispatch({
              type: "ADD_TODO",
              text: this.state.value,
              id: nextTodoId++
            });
            this.setState({ value: "" });
          }}
        >
          Add Todo
        </button>

        <ul>
          {this.props.todos.map(todo => (
            <li key={todo.id}>{todo.text}</li>
          ))}
        </ul>
      </div>
    );
  }
}
```

[App is here](https://codesandbox.io/s/o5j36jokv5)
