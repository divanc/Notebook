| [< Previous abstract](11%20Todo%20React%202.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](13%20Refactoring.md) |
| ----------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | -------------------------------------- |


# React Todo List Example: Filtering Todos

First of all let's create `FilterLink` component, which are links user would click:

```js
// Children are todos scalable by the filter.
export const FilterLink = ({ filter, children }) => {
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        store.dispatch({
          // Hence, we don't need to specify the reducer, just global kid
          type: "SET_VISIBILITY_FILTER",
          filter
        });
      }}
    >
      {children}
    </a>
  );
};
```

Now we can insert those in a rendering `TodoApp`:

```js
{
  /* Show links for visibility */
}
<p>
  Show: {"   "}
  <FilterLink filter="SHOW_ALL">All</FilterLink> {"  "}
  <FilterLink filter="SHOW_ACTIVE">Active</FilterLink> {"  "}
  <FilterLink filter="SHOW_COMPLETED">Completed</FilterLink> {"  "}
</p>;
```

Now let's build that filter:

```js
// Just like a `todo` this is a higher reducer
const getVisibleTodos = (todos, filter) => {
  switch (filter) {
    case "SHOW_ALL":
      return todos;
    case "SHOW_ACTIVE":
      return todos.filter(todo => !todo.isCompleted);
    case "SHOW_COMPLETED":
      return todos.filter(todo => todo.isCompleted);
  }
};
```

In `TodoApp` we want to replace rendering todos:

```js
render() {
  // Instead of **just** todos we would render visible todos
  const visibleTodos = getVisibleTodos(
    this.props.todos,
    this.props.visibilityFilter
  );
...
...
//Instead of
//<ul>
//  {this.props.todos.map(todo => (
//    <li ...
<ul>
  {visibleTodos.map(todo => ( ...
```

And set that prop in the `index`. We used to call just for `state.todos`:

```js
ReactDOM.render(
  <TodoApp todos={store.getState().todos} />,
  document.getElementById("app")
);
```

Yet we can do like this:

```js
ReactDOM.render(
  //That's absolutely ingenious
  <TodoApp {...store.getState()} />,
  document.getElementById("app")
);
```

`store.getStates()` consists of our reducers: `todos` and `visibilityFilter`.

React pasrses it so it can be read as:

```js
ReactDOM.render(
  //That's absolutely  not genious
  <TodoApp
    todos={store.getState().todos}
    visibilityFilter={store.getState().visibilityFilter}
  />,
  document.getElementById("app")
);
```

Let's highlight the current filter:

```js
render() {
  //Let's make it cleaner
  const { todos, visibilityFilter } = this.props;

  // Instead of **just** todos we would render visible todos
  const visibleTodos = getVisibleTodos(todos, visibilityFilter);
...
```

And pass `visibilityFilter` To `FilterLink`s lower as props:

```js
...
{/* Show links for visibility */}
<p>
  Show: {"   "}
  <FilterLink filter="SHOW_ALL" currentFilter={visibilityFilter}>
    All
  </FilterLink>{" "}
  {"  "}
  <FilterLink filter="SHOW_ACTIVE" currentFilter={visibilityFilter}>
    Active
  </FilterLink>{" "}
  {"  "}
  <FilterLink filter="SHOW_COMPLETED" currentFilter={visibilityFilter}>
    Completed
  </FilterLink>{" "}
  {"  "}
</p>
```

```js
export const FilterLink = ({ filter, children, currentFilter }) => {
  if (currentFilter === filter) return <span>{children}</span>;
  return ( ...
```
