
[< Previous abstract](React__2.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__4.md)
----------------------- | ----------------------------|-----------------------------


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
