
[< Previous abstract](React__1.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__3.md)
----------------------- | ----------------------------|-----------------------------


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
