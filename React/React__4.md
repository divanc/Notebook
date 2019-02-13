
## 4. Components & Props

  Conceptually, components are like JavaScript functions. They accept arbitrary inputs (called “props”) and return React elements describing what should appear on the screen.

### Function And Class Components

Javascript function is an example of component:

```js
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
```

It is an example of a component cuz it takes single 'props' and returns a React element

ES6 Class also works:

```js
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```

### Rendering A Component

Previously we had DOM tags:

```js
const element = <div />;
```

Yet we can use our components as well

```js
const element = <Welcome name="Sarah" />;
```

So we can render `Hello, Sarah` on one page like this:

```js
function Welcome(props) {
    return <h1>Hello, {props.name}</h1>;
}
const element = <Welcome name="Sarah" />

ReactDOM.render(
    element,
    document.getElementById('root')
  );
```

  *Note: Always start component names with a capital letter.*

  React treats components starting with lowercase letters as DOM tags. For example, `<div />` represents an HTML div tag, but `<Welcome />` represents a component and requires Welcome to be in scope.

### Composing Components

We can insert components into components creating THE ULTIMATE APP

For example:

```js
function Welcome(props) {
    return <h1>Hello, {props.name}</h1>;
}

function App() {
  return (
    <div>
      <Welcome name="Sarah" />
      <Welcome name="Yura" />
      <Welcome name="Phil" />
    </div>
  );
}

ReactDOM.render(
    <App />,
    document.getElementById('root')
  );
```

### Extracting Components

  Don't be afraid to split components into smaller components

  For example let's consider `Comment` a component

  ```js
function Comment(props) {
  return (
    <div className="Comment">
      <div className="UserInfo">
        <img className="Avatar"
          src="props.author.AvatarUrl"
          alt="props.author.name"
        />
        <div className="Userinfo-Name">
          {props.author.name}
        </div>
      </div>
      <div className="Comment-text">
        {props.text}
      </div>
      <div className="Comment-date">
      {formatDate(props.date)}
      </div>
    </div>
  );
}
```

It accepts author (an object), text (a string), and date (a date) as props, and describes a comment on a social media website

This component is kinda tricky because of all the nesting and hard to reuse
So better solution would be:

`Avatar` component

```js
function Avatar(props) {
  return (
    <img
      src={props.user.avatarUrl}
      alt={props.user.name}
    />
  );
}
```

As we might use it now not only in comment section, we gave a prop more generic name — `user` instead of `author`

Next! Ah, `UserInfo`

```js
function UserInfo(props) {
  return (
    <div className="UserInfo">
      <Avatar user={props.user} />
      <div className="UserInfo-name">
        {props.user.name}
      </div>
    </div>
  );
}
```

So now we have two usable components and comment component will look like this

```js
function Comment(props) {
  return (
    <div className="Comment">
      <UserInfo user={props.author} />
      <div className="Comment-text">
      {props.text}
      </div>
      <div className="Comment-date">
      {formatDate(props.date)}
      </div>
    </div>
  );
}
```

**R E A D A B I L I T Y**

Let's save the Nature — reuse things

### Props & Read-only

  A component MUST NEVER modify it's own props

```js
function sum(a,b) {
  return a+b;
}
```

These are `pure` functions cuz they don't change their input values, unlike this bad boy

```js
function withdraw(account, amount) {
  account.total -= amount;
}
```

The first rule of ~~The Fighting Cl~~ React is to keep components pure. It doesn't mean React is static yet. So here we are with...
