# React JS

    Nice thing to have back in 2000

I am studying [official Docs](https://reactjs.org/docs/).
<hr>

## 1. Hello, World!

Fantastic, I mean, have you seen this? 

```js
ReactDOM.render(
  <h1>Hello, world!</h1>,
  document.getElementById('root')
);
```

<hr>

## 2. JSX

We are free to mix HTML & JS, as a dino, I wanted this forever

Curvies help js to be found in html-like string

```js
const element = (
    <h1>
        Hello, {formatName(user)}!
    </h1>
);
ReactDOM.render(
    element,
    document.getElementById('root')
  );
  ```

### JSX is expression

Oof

```js
function getGreeting(user) {
    if (user) {
        return <h1>Hello, {formatName(user)}!</h1>
    }
    return <h1>Hello, Stranger</h1> // WE CAN DO THAT
}
```

Quotes to specify html attrs:

```js
const element = <div tabIndex="0"></div>;
```

Curlies embeds JS into an attr:

```js
const element = <img src={user.avatarUrl}></img>;
```

  Use just one at a time

### JSX Specifics

* Closing empty tags like in XML `/>`
* JSX Tags may contain children:

```js
const element = (
  <div>
  <h1>Hello!</h1>
  <h2>Good to see you reading the manual of the manual</h2>
  </div>
);
```

### JSX Prevents Injection Attacks

```js
const title = response.potentiallyMaliciousInput;
// Totally safe:
const element = <h1>{title}</h1>;
```

  By default, React DOM escapes any values embedded in JSX before rendering them. Thus it ensures that you can never inject anything that’s not explicitly written in your application. Everything is converted to a string before being rendered. This helps prevent XSS (cross-site-scripting) attacks.

### JSX Represents Objects

Babel compiles JSX down to `React.createElement()` calls

So identical would be

```js
//2019
const element = (
  <h1 className="greeting">
    Hello, Github!
  </h1>
);
```

and

```js
//Meh
const element = React.createElement(
  'h1',
  {className: 'greeting'},
  'Hello, Github!'
);
```

`React.createElement()` creates objects (React Elements) like this:

```js
// Note: this structure is simplified
const element = {
  type: 'h1',
  props: {
    className: 'greeting',
    children: 'Hello, world!'
  }
};
```

<hr>

## 3. Rendering Elements

  Element describes what I want to see on screen, much more plain than browser DOM objects

### Rendering into the DOM

If we have a `<div id="root"></div>` somewhere in our HTML, we call this a "root" DOM node because everything inside will be managed by React DOM.

Apps built with just React usually have single root DOM Node

In order to render a React element into a root DOM node,  we should pass both to `ReactDOM.render()`:

```js
const element = <h1>Hello, World</h1>;
ReactDOM.render(element, document.getElementById('root'));
```

### Updating the Rendering Element

  React elements are immutable. You can't change it's children or attributes.

So far, the only way to update the UI is to create a new element, and pass it to `ReactDOM.render()`

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

It calls `ReactDOM.render()` every second from a `SetInterval()` callback.

  In practice, most React apps only call ReactDOM.render() once. In the next sections we will learn how such code gets encapsulated into stateful components.

  We recommend that you don’t skip topics because they build on each other.

### React Only Updates What's Necessary

Even though we create an element describing the whole UI tree on every tick, only the text node whose contents has changed gets updated by React DOM.

In our experience, thinking about how the UI should look at any given moment rather than how to change it over time eliminates a whole class of bugs.