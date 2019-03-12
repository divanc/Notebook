# Generating Containers with connect()

## VisibleTodoList

1.

```jsx
const mapStateToProps = state => {};
```

For `VisibleTodoList`:

```jsx
const mapStateToProps = state => {
  return {
    todos: getVisibleTodos(state.todos, state.visibilityFilter)
  };
};
```

2.

```jsx
const mapDispatchToProps = dispatch => {
  return {};
};
```

For `VisibleTodoList`:

```jsx
const mapDispatchToProps = dispatch => {
  return {
    onTodoClick: id =>
      dispatch({
        type: "TOGGLE_TODO",
        id
      })
  };
};
```

3.

```jsx
import { connect } from "react-redux";
```

## Focus

Now instead of this:

```jsx
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
        todos={}
        onTodoClick={
        }
      />
    );
  }
}
VisibleTodoList.contextTypes = {
  store: PropTypes.object
};

```

It would be ok to have this:

```jsx
const VisibleTodoList = connect(
  mapStateToProps,
  mapDispatchToProps
)(TodoList);
```

## AddTodo

We had this

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

### Focus

```jsx
let nextTodoId = 0;

export let AddTodo = ({ dispatch }) => {
  let input;
  const ref = useRef(null);

  return (
    <div>
      <input ref={r => (ref.current = r)} />

      <button
        onClick={() => {
          dispatch({
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

AddTodo = connect(
  null,
  null
)(AddTodo);
```

`null`s mean id does not subscribe to a store. This is fine as it is a container-component and wraps itself

## FilterLink

We had:

```jsx
export class FilterLink extends Component {
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
FilterLink.contextTypes = {
  store: PropTypes.object
};
```

### Focus:

```jsx
const mapStateToProps = (state, ownProps) => {
  return {
    active: ownProps.filter === state.visibilityFilter
  };
};

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    onClick: () => {
      dispatch({
        type: "SET_VISIBILITY_FILTER",
        filter: ownProps.filter
      });
    }
  };
};

export const FilterLink = connect(
  mapStateToProps,
  mapDispatchToProps
)(Link);
```
