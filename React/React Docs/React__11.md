[< Previous abstract](React__10.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__12.md)
----------------------- | ----------------------------|-----------------------------

## 11. Composition vs Inheritance

  React has a powerful composition model, and we recommend using composition instead of inheritance to reuse code between components.

### Containment

Some components don't know their children ahead of time. Docs recommend to use special `children` prop to pass children elements directly in the output:

```js
function FancyBorder(props) {
  return (
    <div className={"FancyBorder FancyBorder-"+props.color}>
      {props.children}
    </div>
  )
}
```

```js
function WelcomeDialog(props) {
  return (
    <FancyBorder color="blue">
      <h1 className="Dialog-title">
        Welcome
      </h1>
      <p className="Dialog-message">
        Thank you for joining
      </p>
    </FancyBorder>
  )
}
```

While this is less common, sometimes you might need multiple “holes” in a component. In such cases you may come up with your own convention instead of using children:

```js
function SplitPane(props) {
  return (
    <div className="SplitPane">
      <div className="SplitPane-left">
        {props.left}
      </div>
      <div className="SplitPane-right">
        {props.right}
      </div>
    </div>
  );
}

function App() {
  return (
    <SplitPane
      left={
        <Contacts />
      }
      right={
        <Chat />
      } />
  );
}
```

### Specialization

  Sometimes we think about components as being “special cases” of other components. For example, we might say that a WelcomeDialog is a special case of Dialog.

In React  more “specific” component renders a more “generic” one and configures it with props

```js
function Dialog(props) {
  return (
    <FancyBorder color="blue">
      <h1 className="Dialog-title">
        {props.title}
      </h1>
      <p className="Dialog-message">
        {props.message}
      </p>
    </FancyBorder>
  );
}

function WelcomeDialog(props) {
  return (
    <Dialog
      title="Welcome"
      message="To The family, son" />
  );
}
```

Composition works just as well for classes.

```js
function Dialog(props) {
  return (
    <FancyBorder color="blue">
      <h1 className="Dialog-title">
        {props.title}
      </h1>
      <p className="Dialog-message">
        {props.message}
      </p>
      {props.children}
    </FancyBorder>
  );
}

class SignUpDialog extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this)
    this.handleSignUp = this.handleSignUp.bind(this)
    this.state = {login: ''}
  }

  render() {
    return (
      <Dialog title="Mars Exploration Program"
        message="How we should refer to you?">
      <input value={this.state.login} onChange={this.handleChange} />
      <button onClick={this.handleSignUp} > Sign Me Up! </button>
      </Dialog>
    );
  }

  handleChange(event) {
    this.setState({login: event.target.value});
  }
  handleSignUp() {
    alert('Welcome aboard, ${this.state.login}!');
  }
}
```

### Inheritance??

  At Facebook, we use React in thousands of components, and we haven’t found any use cases where we would recommend creating component inheritance hierarchies.

  Props and composition give you all the flexibility you need to customize a component’s look and behavior in an explicit and safe way. Remember that components may accept arbitrary props, including primitive values, React elements, or functions.

  If you want to reuse non-UI functionality between components, we suggest extracting it into a separate JavaScript module. The components may import it and use that function, object, or a class, without extending it.