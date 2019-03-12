# Action Creator

In our `AddTodo` component we have such dispatch:

```jsx
dispatch({
  type: "ADD_TODO",
  id: nextTodoId++,
  text: ref.current.value
});
```

The problem is `nextTodoId` is local, what if another func want to ADD TODO?

Well, let's make a separate file action

```jsx
<button
  onClick={() => {
    dispatch(addTodo(ref.current.value));
    ref.current.value = "";
  }}
>
```

All `actions/addTodo` would look like:

```jsx
let nextTodoId = 0;

export const addTodo = text => {
  return {
    type: "ADD_TODO",
    id: nextTodoId++,
    text
  };
};
```

Same to others;

# Thank you, Dan
