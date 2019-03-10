# Dan Abramov On Redux

<p align="center">
<img src="https://images.contentful.com/34rjphroaymg/4rnTNHjdSoEWmYUg2YgE0y/8fe539e8bb8b6465e107b6154f669f11/redux.svg" width="240">
</p>

[This]() wonderful course.

1. [Redux: The Immutable State Tree](#1)
2. [Reducer](#2)
3. [Redux Store Methods](#3)
4. [Implemeting Store From Scratch](#4)
5. [React Counter Example](#5)
6. [Avoiding Array Mutations](#6)
7. [Avoiding Object Mutations](#7)

Danчик will teach us how to manage state of our React App via Redux;

Redux can turn your mess into a beautiful mess, using loads of code.

Once you are finished with this course be sure to check out part 2: [building-react-applications-with-idiomatic-redux](https://egghead.io/courses/building-react-applications-with-idiomatic-redux)

<a name="1"></a>

## Redux: The Immutable State Tree

**First principle of Redux: no matter how hard your app is, you save all your states in one JS object**.

All mutations are explicit

### Example

With one counter we just hold it's value in state and it's fine, however with several counters we want to give each one an ID in state to make sure which is being changed

If we have a todo app, we want to store all aspects in one object, so if a task is done, it is cheched within an array of this tash within an array of all tasks \*it is called a **state tree\***

### We Describe State Changes with _Actions_

It is a js object, the minimum representation of state:

```js
[object Object] { /// THIS IS STATE TREE
  type: 'INCREMENT' // THIS IS ONE ACTION
}
```

for counter we need only `increment` and `decrement` for adding or substracting.

For several we have to add an `id` to every counter.

**Second dogma**: Redux can only change state via _actions_

### Pure and Impure functions

Pure functions return values depending only on arguments. They are predictable. They do not overwrite arguments.

Impure can call dbs, overwrite and do stuff.

In Redux we want pure funcs

<a name='2'></a>

## Reducer

State mutations should be described as a pure functions.

It has to be pure, hence, has to return a new object.

**Third dogma: Redux takes the previous state tree and action, dispatches action and returns a new tree**

This function, which takes prev and action and makes a new one is calledd the _reducer_.

### Let's write one!

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

<a name='3'></a>

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

<a name='4'></a>

## Implemeting Store From Scratch

Let's build our own Redux from scratch, you know, just for fun.

Instead of `const store = createStore(reducer);` we will write our own. We know arg is reducer...

```js
const createStore = reducer => {
  //We also know it keeps our state
  let state;

  // We know it has 3 methods
  const getState = () => state;

  const dispatch = action => {};

  const subscribe = listener => {};

  // And this is what we call a Redux store
  return { getState, dispatch, subscribe };
};
```

Because we can call subcribe several times, we need to keep out listeners:

```js
const createStore = reducer => {
  let state;

  let listeners = [];

  const getState = () => state;

  const dispatch = action => {
    //This is how we add a state
    state = reducer(state, action);
    //We should also notify each listener about the change
    listeners.forEach(listener => listener());
  };

  const subscribe = listener => {
    //Every time it is called we want to push our new listener into the array
    listeners.push(listener);
  };

  return { getState, dispatch, subscribe };
};
```

We would also like to add unsubcscribe method right in the `subscribe`:

```js
const subscribe = listener => {
  listeners.push(listener);

  return () => {
    listeners = listeners.filter(l => l !== listener);
  };
};
```

Finally we would like to return initial value, so we just `dispatch` before _store_ is returned:

```js
dispatch({});
return { getState, dispatch, subscribe };
```

<a name='5'></a>

## React Counter Example

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

<a name='6'></a>

## Avoiding Array Mutations

### `concat()` is good

If we want to add zero in the past array, we never want to use `append()`: it mutates!

```js
const addCounter = list => {
  return list.concat([0]); // makes a new array
};
```

Another way to write it in ES6:

```js
const addCounter = list => {
  return [...list, 0];
};
```

### `splice()` is bad, sorry

Because it mutates an array, instead we would use:

```js
const removeCounter = (list, index) => {
  return list.slice(0, index).concat(list.slice(index + 1));
};
```

In ES6:

```js
const removeCounter = (list, index) => {
  return [...list.slice(0, index), ...list.slice(index + 1)];
};
```

### Adding

This is mutating our list, not what we really want

```js
const addIncrement = (list, index) => {
  list[index]++;
  return list;
};
```

Instead, let's try it out:

```js
const addIncrement = (list, index) => {
  return list
    .slice(0, index) // We want to take our list just before the position
    .concat([list[index] + 1]) // Insert a value
    .concat(list.slice(index + 1)); // And add the rest
};
```

In ES6:

```js
const addIncrement = (list, index) => {
  return [...list.slice(0, index), list[index] + 1, ...list.slice(index + 1)];
};
```

<a name='7'></a>

## Avoiding Object Mutations

With `.assign()`

```js
const testToggleTodos = () => {
  const todoBefore = {
    id: 0,
    text: "Learn Redux",
    isCompleted: false
  };
  const todoAfter = {
    id: 0,
    text: "Learn Redux",
    isCompleted: true
  };
};

//This func
const toggleTodos = todo => {
  todo.isCompleted = !todo.isCompleted;
  return todo;
};
```

And this is bad, cuz it tries to mutate our Object;

Simpliest solution would be:

```js
const toggleTodos = todo => {
  return {
    id: todo.id,
    text: todo.text,
    isCompleted: !todo.isCompleted
  };
};
```

Yet we can forget to update that from time to time. This is why there are a nice ES6 method called `Object.assign()`

```js
const toggleTodos = todo => {
  return Object.assign({}, todo, {
    isCompleted: !todo.isCompleted
  });
};
```

Perfect!

## Writing a todo reducer

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

### Toggling todo

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

### Reducer Composition With Arrays

We faced a problem, that our reducer traces both whole arrays of todos and separate ones, we can solve it having one extra functions

```js
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
```

Now we can call this reducer in our first reducer:

```js
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
```

### Composition with Objects

What if we want control more advanced options? For ex, user may choose what kind of tasks he wants to see:

```js
const visibilityFilter = (state = "SHOW_ALL", action) => {
  switch (action.type) {
    case "SET_VISIBILIty_FILTER":
      return action.filter;
    default:
      return state;
  }
};
```

Instead of modifying our reducers we can create new, which will contain all of them, as they do different things:

```js
const todoApp = (state = {}, action) => {
  return {
    todos: todos(state.todos, action),
    visibilityFilter: visibilityFilter(state.visibilityFilter, action)
  };
};
```

Now having papa-reducer we would use it to create our store

```js
const store = createStore(todoApp);
```

### Actually, there is a `combineReducers()`

In fact, what we just did already is in Redux:

```js
import { combineReducers } from "redux";
const todoApp = combineReducers({
  todos: todos, //stateTree name and reducer
  visibilityFilter: visibilityFilter
});
```

Also ES6 provides a shorthand literal notation, so we can do this:

```js
import { combineReducers } from "redux";
const todoApp = combineReducers({
  todos,
  visibilityFilter
});
```

As they are the same

### `combineReducer()` From Scratch

```js
const combineReducers = reducers => {
  //We know it returns a reducer
  return (state = {}, action) => {};
};
```

```js
const combineReducers = reducers => {
  //We know it returns a reducer
  return (state = {}, action) => {
    //Object.keys would return keys of our object; In our case: `todos` and `visibilityFilter`
    return Object.keys(reducers).reduce(
      // reduce will affect only one value among, such as `nextState`
      (nextState, key) => {
        //So nextState[key] is `isCompleted` of [todos] for example
        //To change it we would call `todos` reducer === reducers[key]() and pass it arguments
        nextState[key] = reducers[key](state[key], action);

        return nextState;
      },
      {}
    );
  };
};
```

And it's all

Recap:

We call `combineReducers` with args of reducers. It is a little bit complex, mates.

## React Todo List Example

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
