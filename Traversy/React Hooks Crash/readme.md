# Traversy: Introducing React Hooks

[This](https://www.youtube.com/watch?v=mxK8b99iJTg) video.

**Hooks**: How to use states without classes.

```js
import React, { useState } from 'react';
```

Firstly, lets set a frame for a todo app:

```js
function App() {
  const [todos, setTodos] = useState([
    {
      text: 'Learn Hooks',
      isCompleted: false
    }, {
      text: 'Build Todo with design, not like Yuras',
      isCompleted: false
    }, {
      text: 'Go to Yandex',
      isCompleted: false
    },
  ]);
  return (
    <div className="app">
      <div className='todo-list'>
        {todos.map((todo, index) => (
          <Todo key={index} index={index} todo={todo} />
        ))}
      </div>
    </div>
  );
}
```

And Todo component:

```js
function Todo({ todo, index}) {
  return(
    <div className="todo">
      { todo.text }
    </div>
  );
}
```

Then, adding Todo Form component, which renders user input:

```js
function TodoForm({addTodo}) {
  const [value, setValue] = useState('');

  const handleSubmit = event => {
    event.preventDefault();
    if(!value) return;
    addTodo(value);
    setValue('');
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type='text' className='input' value={value} onChange={e => setValue(e.target.value)} />

    </form>
  );
}
```

Render that on page inside of `App`:

```js
 const addTodo = text => {
    const newTodos = [...todos, { text }];
    setTodos(newTodos);
  }

  return (
    <div className="app">
      <div className='todo-list'>
        {todos.map((todo, index) => (
          <Todo key={index} index={index} todo={todo} />
        ))}
        <TodoForm addTodo={addTodo} />
      </div>
    </div>
  );
}
```

Then just a feature of completion of a todo:

```js
function Todo({ todo, index, completeTodo}) {
  return(
    <div className="todo" style={{ textDecoration: todo.isCompleted ? 'line-through' : ''}}>
      { todo.text }
      <div>
        <button onClick={() => { completeTodo(index)}}>
          Complete
        </button>
      </div>
    </div>
  );
}
```

`completeTodo` is rendered in `App`: 

```js
const addTodo = text => {
    const newTodos = [...todos, { text }];
    setTodos(newTodos);
  }

  const completeTodo = index => {
    const newTodos = [...todos];
    newTodos[index].isCompleted = true;
    setTodos(newTodos);
  }
  return (
    <div className="app">
      <div className='todo-list'>
        {todos.map((todo, index) => (
          <Todo key={index} index={index} todo={todo} completeTodo={completeTodo} />
        ))}
        <TodoForm addTodo={addTodo} />
      </div>
    </div>
  );
}
```

And delete feature, pretty much the same thing:

```js
function Todo({ todo, index, completeTodo, deleteTodo}) {
  return(
    <div className="todo" style={{ textDecoration: todo.isCompleted ? 'line-through' : ''}}>
      { todo.text }
      <div>
        <button onClick={() => { completeTodo(index)}}>
          Complete
        </button>
        <button onClick={() => { deleteTodo(index)}}>
          x
        </button>
      </div>
    </div>
  );
}
```

And in `App`:

```js
 const deleteTo = index => {
    const newTodos = [...todos];
    newTodos.splice(index, 1);
    setTodos(newTodos);
  }

  return (
    <div className="app">
      <div className='todo-list'>
        {todos.map((todo, index) => (
          <Todo key={index} 
                index={index}
                todo={todo} 
                completeTodo={completeTodo} 
                deleteTodo={deleteTo} 
                />
        ))}
        <TodoForm addTodo={addTodo} />
      </div>
    </div>
  );
}

```

All in all we have this file:

```js
import React, {useState} from 'react';
import ReactDOM from 'react-dom';

function Todo({ todo, index, completeTodo, deleteTodo}) {
  return(
    <div className="todo" style={{ textDecoration: todo.isCompleted ? 'line-through' : ''}}>
      { todo.text }
      <div>
        <button onClick={() => { completeTodo(index)}}>
          Complete
        </button>
        <button onClick={() => { deleteTodo(index)}}>
          x
        </button>
      </div>
    </div>
  );
}

function TodoForm({addTodo}) {
  const [value, setValue] = useState('');

  const handleSubmit = event => {
    event.preventDefault();
    if(!value) return;
    addTodo(value);
    setValue('');
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type='text'
        className='input'
        value={value}
        placeholder="Insert a message"
        onChange={e => setValue(e.target.value)}
        />

    </form>
  );
}

function App() {
  const [todos, setTodos] = useState([
    {
      text: 'Learn Hooks',
      isCompleted: false
    }, {
      text: 'Build Todo with design, not like Yuras',
      isCompleted: false
    }, {
      text: 'Go to Yandex',
      isCompleted: false
    },
  ]);

  const addTodo = text => {
    const newTodos = [...todos, { text }];
    setTodos(newTodos);
  }

  const completeTodo = index => {
    const newTodos = [...todos];
    newTodos[index].isCompleted = true;
    setTodos(newTodos);
  }

  const deleteTo = index => {
    const newTodos = [...todos];
    newTodos.splice(index, 1);
    setTodos(newTodos);
  }

  return (
    <div className="app">
      <div className='todo-list'>
        {todos.map((todo, index) => (
          <Todo key={index} 
                index={index}
                todo={todo} 
                completeTodo={completeTodo} 
                deleteTodo={deleteTo} 
                />
        ))}
        <TodoForm addTodo={addTodo} />
      </div>
    </div>
  );
}



