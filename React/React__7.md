

<p style="text-align:center;">[< Previous abstract](React__6.md)
[Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React)
[Next abstract >](React__8.md)
</p>

## Conditional Rendering

  In React, you can create distinct components that encapsulate behavior you need. Then, you can render only some of them, depending on the state of your application.

Conditional rendering in React = Conditional rendering in JS

```js
function UserGreeting(props) {
  return <h1>Hello, {props.name}!</h1>;
}

function GuestGreeting(props) {
  return <h1>Please sign up</h1>;
}
```

Let's create a summing component, that will display either of these

```js
function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return <UserGreeting />
  }
  return <GuestGreeting />
}

ReactDOM.render (
  <Greeting isLoggedIn={true} name='Yura' />,
  document.getElementById('root')
);
```

### Element Variables

```js
function LoginButton(props) {
  return(
    <button onClick={props.onClick}>
      Login
    </button>
  );
}

function LogoutButton(props) {
  return(
    <button onClick={props.onClick}>
      Logout
    </button>
  );
}
```

Let's create stateful component `LoginControl`

```js
class LoginControl extends React.Component {
  constructor(props) {
    super(props);
    this.handleLoginClick = this.handleLoginClick.bind(this);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
    this.state = {isLoggedIn: false};
  }
  
  handleLoginClick() {
    this.setState({isLoggedIn: true});
  }
  
  handleLogoutClick() {
    this.setState({isLoggedIn: false});
  }

  render() {
    const isLoggedIn = this.state.isLoggedIn;
    let button;

    if (isLoggedIn) {
      button = <LogoutButton onClick={this.handleLogoutClick} />;
    } else {
      button = <LoginButton onClick={this.handleLoginClick} />;
    }

    return (
      <div>
      <Greeting isLoggedIn={isLoggedIn} />
      {button}
      </div>
    );
  }
}

ReactDOM.render(
  <LoginControl />,
  document.getElementById('root')
);
```

The fact is we can shorten this logic furthermore

### Inline If With && Operator

JS `&&` and JSX `{}` may work together perfectly in conditions

```js
function Mailbox(props) {
  const unreadMessages = props.unreadMessages
  return (
    <div>
      <h1>Hello!</h1>
      {unreadMessages.length > 0 && 
      <h2>You have {unreadMessages.length} unread messages.</h2>
      }
    </div>
  );
}

const messages = ['React','Re:React','Re:Re:React'];
ReactDOM.render(
<Mailbox unreadMessages={messages} />,
document.getElementById('root')
);

```

This works, because JS renders this as `expression && false == false`

### Inline If-Else And Conditional Operator

  `condition ? true : false`

```js
render() {
  const isLoggedIn = this.state.isLoggedIn;
  return (
    <div>
    <h1> This user is {isLoggedIn ? 'currently' : 'not'} logged in </h1>;
    </div>
  );
}
```

Works for bigger expressions as well in case

```js
condition ? (
  true) : (
  false)
```

### Preventing From Rendering

Sometimes you might want component to hide itself. It can be done via `return null`, so component wont render

```js
function WarnningBanner(props) {
  if (!props.warn) {
    return null;
  }

  return (
    <div className="Warning">
    Warning
    {alert("Warning!")}
    </div>
  );
}

class Page extends React.Component {
  constructor(props) {
    super(props);
    this.state = {showWarning: true};
    this.handleToggleClick = this.handleToggleClick.bind(this);
  }

  handleToggleClick() {
    this.setState(state => ({
      showWarning: !state.showWarning
    }));
  }

  render() {
    return (
      <div>
        <WarningBanner warn={this.state.showWarning} />
        <button onClick={this.handleToggleClick}>
          {this.state.showWarning ? 'Hide' : 'Show'}
        </button>
      </div>
    );
  }
}

ReactDOM.render(
  <Page />,
  document.getElementById('root')
);
```

  Returning null from a component’s render method does not affect the firing of the component’s lifecycle methods. For instance componentDidUpdate will still be called.
