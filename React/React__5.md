
## 5. State and Lifecycle

  Consider the ticking clock example from one of the previous sections. In Rendering Elements, we have only learned one way to update the UI. We call ReactDOM.render() to change the rendered output:

```js
function tick() {
    const element = (
        <div>
            <h1>Hello, World!</h1>
            <h2>It is {new Date().toLocaleString()}.</h2>
        </div>
    )

    ReactDOM.render(
        element,
        document.getElementById('root')
      );
}

setInterval(tick,1000);
```

So in this section we will learn how to make this `clock` component truly reusable

Starting with encapsulation:

```js
function Clock(props) {
  return (
    <div>
      <h1>Hello, World!</h1>
      <h2>It is {props.date.toLocaleString()}.</h2>
    </div>
  );
}

function tick() {
  ReactDOM.render(
    <Clock date={new Date()} />,
    document.getElementById('root')
  );
}

setInterval(tick,1000);
```

However the update-feature should be the part of the clock, not external, cuz clock ticks.

Ideally we want to get this:

```js
ReactDOM.render(
  <Clock />,
  document.getElementById('root')
);
```

In order to do this, we need to add `state` to the `Clock`

State is kinda like props, but is being controlled from the inside. There is yet one problem.

**State is a feature of class-component**

### Easy converting Function >> Class

1. ES6 Class, that extends React.Component
2. Add an empty method `render()`
3. The body of the function in to the body of the render()
4. props >> this.props

```js
class Clock extends React.Component {
  render() {
    return (
      <div>
        <h1>Hello, World!</h1>
        <h2>It is {this.props.date.toLocaleString()}.</h2>
      </div>
    );
  }
}
```

`render()` will be called automatically each time an update happens. Ain't that cool?

### Adding Local State To A Class

1. Replace `this.`__props__`.date` with `this.`__state__`.date` in the `render()` method
2. Add a class constructor, that assigns the initial `this.state`

```js
class Clock extends React.Component {
  constructor(props) {
    super(props); //important
    this.state = {date: new Date()}
  }

  render() {
    return (
      <div>
        <h1>Hello, World!</h1>
        <h2>It is {this.state.date.toLocaleString()}.</h2>
      </div>
    );
  }
}

//And our ideal model
ReactDOM.render(
  <Clock />,
  document.getElementById('root')
);
```

And now we need to add to the `Clock` it's own timer

### Adding Lifecycle Methods To The Class

  In applications with many components, it’s very important to free up resources taken by the components when they are destroyed.
  
  We want to set up a timer whenever the Clock is rendered to the DOM for the first time. This is called “mounting” in React.

  We also want to clear that timer whenever the DOM produced by the Clock is removed. This is called “unmounting” in React.

We can declare specials methods to mount & unmount timer:

```js
class Clock extends React.Component {
  constructor(props) {
    super(props); //important
    this.state = {date: new Date()}
  }

  // This boy
  componentDidMount() {

  }

  // And This boy
  componentWillUnmount() {

  }

  render() {
    return (
      <div>
        <h1>Hello, World!</h1>
        <h2>It is {this.state.date.toLocaleString()}.</h2>
      </div>
    );
  }
}
```

These are `lifecycle methods`

The `componentDidMount()` works after component outputs to the DOM (nice place to setup a timer)

```js
componentDidMount() {
  this.timerID = setInterval(
    () => this.tick(),
    1000
  );
}
```

The `componentWillUnmount()` will break the timer

```js
componentWillUnmount() {
  clearInterval(this.timerID);
}
```

Finally, we set `tick()` method to represent the date

```js
tick() {
  this.setState({
    date: new Date()
  });
}
```

Done! It ticks every second

```js
class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {date: new Date()};
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => this.tick(),
      1000
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  tick() {
    this.setState({
      date: new Date()
    });
  }

  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
      </div>
    );
  }
}

ReactDOM.render(
  <Clock />,
  document.getElementById('root')
);
```

1. When `<Clock />` is passed to ReactDOM.render(), React calls the constructor of the Clock component. Since Clock needs to display the current time, it initializes this.state with an object including the current time. We will later update this state.

2. React then calls the Clock component’s `render()` method. This is how React learns what should be displayed on the screen. React then updates the DOM to match the Clock’s render output.

3. When the Clock output is inserted in the DOM, React calls the `componentDidMount()` lifecycle method. Inside it, the Clock component asks the browser to set up a timer to call the component’s `tick()` method once a second.

4. Every second the browser calls the `tick()` method. Inside it, the Clock component schedules a UI update by calling `setState()` with an object containing the current time. Thanks to the `setState()` call, React knows the state has changed, and calls the `render()` method again to learn what should be on the screen. This time, `this.state.date` in the `render()` method will be different, and so the render output will include the updated time. React updates the DOM accordingly.

5. If the `Clock` component is ever removed from the DOM, React calls the `componentWillUnmount()` lifecycle method so the timer is stopped.

### Using State Correctly

Three things to know about `setState()`

#### Do Not Modify State Directly

```js
//Won't work
this.state.comment = "Hello";

//Superb
this.setState({comment: 'Hello'});

```

Setting `this.state` can be done only in `constructor`.

#### State Updates Can Be Asynchronous

React may batch multiple setState() calls into a single update for performance.

  Because `this.props` and `this.state` may be updated asynchronously, you should not rely on their values for calculating the next state.

```js
//MAY FAIL
this.setState({
  counter: this.state.counter + this.props.increment
});
```

To fix it, use a second form of `setState()` that accepts a function rather than an object. That function will receive the previous state as the first argument, and the props at the time the update is applied as the second argument:

```js
// Correct
this.setState((state, props) => ({
  counter: state.counter + props.increment
}));

//Or without arrow func
// Correct
this.setState(function(state, props) {
  return {
    counter: state.counter + props.increment
  };
});
```

#### State Updates are Merged

When you call `setState()`, React merges the object you provide into the current state.

```js
constructor(props) {
    super(props);
    this.state = {
      posts: [],
      comments: []
    };
  }
```

Then you can update them independently with separate `setState()` calls:

```js
componentDidMount() {
    fetchPosts().then(response => {
      this.setState({
        posts: response.posts
      });
    });

    fetchComments().then(response => {
      this.setState({
        comments: response.comments
      });
    });
}
```

### The Data Flows Down

  Neither parent nor child components can know if a certain component is stateful or stateless, and they shouldn’t care whether it is defined as a function or a class.

  This is why state is often called local or encapsulated. It is not accessible to any component other than the one that owns and sets it.

  This is commonly called a “top-down” or “unidirectional” data flow. Any state is always owned by some specific component, and any data or UI derived from that state can only affect components “below” them in the tree.

  If you imagine a component tree as a waterfall of props, each component’s state is like an additional water source that joins it at an arbitrary point but also flows down.