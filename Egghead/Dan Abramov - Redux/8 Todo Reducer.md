| [< Previous abstract](7%20Avoiding%20Object%20Mutations.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](9%20Todo%20Composition.md) |
| ----------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | -------------------------------------------- |


# Writing a todo reducer

```js
const todos = (state = [], action) => {};

const testAddTodos = () => {
  const stateBefore = [];
  const action = {
    type: "ADD_TODO",
    id: 0,
    text: "Learn Redux"
  };

  const stateAfter = [
    {
      id: 0,
      text: "Learn Redux",
      isCompleted: false
    }
  ];
};

testAddTodos();
```

As before we face a problem of our reducer not knowing the action:

```js
const todos = (state = [], action) => {
  switch (action.type) {
    case "ADD_TODO":
      return [
        ...state,
        {
          id: action.id,
          text: action.text,
          isCompleted: false
        }
      ];
    default:
      return state;
  }
};
```

That's the basic reducer

## Toggling todo

```js
testToggleTodos = () => {
  const stateBefore = [
    {
      id: 0,
      text: "Learn Redux",
      isCompleted: false
    },
    {
      id: 1,
      text: "Go shopping",
      isCompleted: false
    }
  ];
  const action = {
    type: "TOGGLE_TODO",
    id: 1
  };

  const stateAfter = [
    {
      id: 0,
      text: "Learn Redux",
      isCompleted: false
    },
    {
      id: 1,
      text: "Go shopping",
      isCompleted: true
    }
  ];
};

testToggleTodo();
```

Lets remember that reducer must be a pure function

Now let's handle that action:

```js
const todos = (state = [], action) => {
  switch (action.type) {
    case "ADD_TODO":
      return [
        ...state,
        {
          id: action.id,
          text: action.text,
          isCompleted: false
        }
      ];
    //That's our new action
    case "TOGGLE_TODO":
      return state.map( => {
        if (todo.id !== action.id) return todo;
        return {
          ...todo,
          completed: !todo.completed
        };
      })
    default:
      return state;
  }
};
```
