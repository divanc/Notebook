
[< Previous abstract](React__5.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__7.md)
----------------------- | ----------------------------|-----------------------------

## Handling Events

Handling is similar to HTML, yet you have to use camelCase:

HTML | JSX
------------ | -------------
```<button onclick="activateLasers" />```| ```<button onClick={activateLasers} />```

Another difference is you cannot `return false` in React:

HTML would be:

```html
<a href="#" onclick="console.log('Click!'); return false">Click</a>
```

Whereas React would use `preventDefault` method:

```js
function ActionLink() {
  function handleClick(e) {
      e.preventDefault();
      console.log('Click');
  }

  return (
    <a href="#" onClick={handleClick}> Click </a>
  );
}
```

Here, `e` is an event

Event handler can be typically a method in ES6 class like this example of `Toggle` component

```js
class Toggle extends React.Component {
  constructor(props) {
    super(props);
    this.state = {isToggleOn: true};

    //This is necassary in order button to work in a callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    this.setState(state => ({
      isToggleOn: !state.isToggleOn
    }));
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        {this.state.isToggleOn ? "ON" : "OFF"}
      </button>
    );
  }
}

ReactDOM.render(
  <Toggle />,
  document.getElementById('root');
);
```

In JS class methods *are not bound by default*. If you forget to bind `handleClick` and pass it to `onClick` it would be undefined.

There are alternatives, which are experimental:

```js
class LoggingButton extends React.Component {
  // This syntax ensures `this` is bound within handleClick.
  // Warning: this is *experimental* syntax.
  handleClick = () => {
    console.log('this is:', this);
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        Click me
      </button>
    );
  }
}
```

Or arrowfunction in the callback:

```js
class LoggingButton extends React.Component {
  handleClick() {
    console.log('this is:', this);
  }

  render() {
    // This syntax ensures `this` is bound within handleClick
    return (
      <button onClick={(e) => this.handleClick(e)}>
        Click me
      </button>
    );
  }
}
```

This is fine for most cases, whilst this method generates different callback every time. If you need to pass this callback lower, this won't be your case. **Docs** recommend first two methods.

### Passing Agrguments to Event Handlers

Both of these would work:

```js
<button onClick={(e) => this.deleteRow(id, e)} />
<button onClick={this.deleteRow.bind(this, id)} />
```

Both cases would pass `e` as an event after the ID