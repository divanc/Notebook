# Refactoring our todo

    Keep it presentational

## Todo

First of all — `todo` component

Mapping an array we had this:

```js
<li
  key={todo.id}
  // Adding onClick dispatching
  onClick={() => {
    store.dispatch({
      type: "TOGGLE_TODO",
      id: todo.id
    });
  }}
  style={{
    textDecoration: todo.isCompleted ? "line-through" : "none"
  }}
>
  {todo.text}
</li>
```

Making it a sole component we would do certain notes:

1. It is no longer a mapping so we don't need `key`:

```js
<li
// v Now useless v
//  key={todo.id}
 ...
```

2. We had `onClick` with dispatching, however we want to keep such components _presentational_, they shouldn't have any logic within

```js
...
//  v Now unclever v
//onClick={() => {
//  store.dispatch({
//    type: "TOGGLE_TODO",
//    id: todo.id
//  });
//}}

// Instead we pass a prop:
onClick={onClick}
```

3. Instead of passing a `todo` Object we want it to render real data so it can be universal:

```js
// v Now barbarian v
//  style={{
//    textDecoration: todo.isCompleted ? "line-through" : "none"
//  }}
//>
//  {todo.text}
//</li>
...
// Instead we pass each as a prop
  style={{
    textDecoration: isCompleted ? "line-through" : "none"
  }}
>
  {text}
...
```

At last we have:

```js
export const Todo = ({ onClick, isCompleted, text }) => (
  <li
    onClick={onClick}
    style={{
      textDecoration: isCompleted ? "line-through" : "none"
    }}
  >
    {text}
  </li>
);
```

## TodoList

Then `TodoList` deserves to become a component too!

We still want to make it presentational, onClick passes over

```js
export const TodoList = ({ todos, onTodoClick }) => (
  <ul>
    {todos.map(todo => (
      <Todo
        key={todo.id}
        {...todo} /* I can't stop praising it */
        onClick={() => onTodoClick(todo.id)}
      />
    ))}
  </ul>
);
```

## Container Components

We want wrap presentational components in a _container_ component to specify the behaviour.

It is basically `TodoApp`, there let's use our components:

```js
...
  Add Todo
</button>

<TodoList
  todos={visibleTodos}
  onTodoClick={id => {
    store.dispatch({
      type: "TOGGLE_TODO",
      id
    });
  }}
/>
```

## AddTodo: Input + Button

We had this:

```jsx
{/* They do it using `ref`, whatever */}
<input value={this.state.value} onChange={this.onChange} />

<button
  onClick={() => {
    // `store.dispatch` just adds to a store
    store.dispatch({
      type: "ADD_TODO",
      text: this.state.value,
      id: nextTodoId++
    });
    // Erases current input value
    this.setState({ value: "" });
  }}
>
  Add Todo
</button>
```

Now our presentational component:

```jsx
export const AddTodo = ({ onAddClick }) => {
  let input;
  return (
    <div>
      <input
        ref={node => {
          input = node;
        }}
      />

      <button
        onClick={() => {
          onAddClick(input.value);
          input.value = "";
        }}
      >
        Add Todo
      </button>
    </div>
  );
};
```

And in container component:

```jsx
<div>
  <AddTodo
    onAddClick={text =>
      store.dispatch({
        type: "ADD_TODO",
        id: nextTodoId++,
        text
      })
    }
  />
```

## Footer: Links

We had:

```jsx
{
  /* Show links for visibility */
}
<p>
  Show: {"   "}
  <FilterLink filter="SHOW_ALL" currentFilter={visibilityFilter}>
    All
  </FilterLink> {"  "}
  <FilterLink filter="SHOW_ACTIVE" currentFilter={visibilityFilter}>
    Active
  </FilterLink>{" "}
  {"  "}
  <FilterLink filter="SHOW_COMPLETED" currentFilter={visibilityFilter}>
    Completed
  </FilterLink> {"  "}
</p>;
```

We want to make `Footer` and `FilterLink` to become presentational, so in `FooterLink` we uplift `onClick`:

```jsx
...
onClick={e => {
  e.preventDefault();
  onClick(filter);
}}
...
```

So `Footer` now is:

```jsx
export const Footer = ({ visibilityFilter, onFilterClick }) => (
  <p>
    Show: {"   "}
    <FilterLink
      filter="SHOW_ALL"
      currentFilter={visibilityFilter}
      onClick={onFilterClick}
    >
      All
    </FilterLink> {"  "}
    <FilterLink
      filter="SHOW_ACTIVE"
      currentFilter={visibilityFilter}
      onClick={onFilterClick}
    >
      Active
    </FilterLink>{" "}
    {"  "}
    <FilterLink
      filter="SHOW_COMPLETED"
      currentFilter={visibilityFilter}
      onClick={onFilterClick}
    >
      Completed
    </FilterLink> {"  "}
  </p>
);
```

In container:

````jsx
<Footer
  visibilityFilter={visibilityFilter}
  onFilterClick={filter => {
    store.dispatch({
      type: "SET_VISIBILITY_FILTER",
      filter
    });
    ```
````

## Finally

We can set `TodoApp` to become function, not class, it is not needed any longer:

### It was:

```jsx
export class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    // Saving a value of input
    this.state = { value: "" };
    this.onChange = this.onChange.bind(this);
  }

  onChange(event) {
    this.setState({ value: event.target.value });
  }

  render() {
    const { todos, visibilityFilter } = this.props;

    // Instead of **just** todos we would render visible todos
    const visibleTodos = getVisibleTodos(todos, visibilityFilter);
    return (
      <div>
        <AddTodo
          onAddClick={text =>
            store.dispatch({
              type: "ADD_TODO",
              id: nextTodoId++,
              text
            })
          }
        />

        <TodoList
          todos={visibleTodos}
          onTodoClick={id => {
            store.dispatch({
              type: "TOGGLE_TODO",
              id
            });
          }}
        />

        <Footer
          visibilityFilter={visibilityFilter}
          onFilterClick={filter => {
            store.dispatch({
              type: "SET_VISIBILITY_FILTER",
              filter
            });
          }}
        />
      </div>
    );
  }
}
```

### NOW COMES MAGIC

```jsx
export const TodoApp = ({ todos, visibilityFilter }) => (
  <div>
    <AddTodo
      onAddClick={text =>
        store.dispatch({
          type: "ADD_TODO",
          id: nextTodoId++,
          text
        })
      }
    />

    <TodoList
      todos={getVisibleTodos(todos, visibilityFilter)}
      onTodoClick={id => {
        store.dispatch({
          type: "TOGGLE_TODO",
          id
        });
      }}
    />

    <Footer
      visibilityFilter={visibilityFilter}
      onFilterClick={filter => {
        store.dispatch({
          type: "SET_VISIBILITY_FILTER",
          filter
        });
      }}
    />
  </div>
);
```
