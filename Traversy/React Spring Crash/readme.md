# Traversy: React Spring Crash Course

[This](https://www.youtube.com/watch?v=S8yn3-WpVV8) boy.

I already have big docs abstract [over here](https://github.com/Betra/Course-Abstract/tree/master/React/React%20Spring).

```console
npm install react-spring
```

First of all make a basic Component1:

```js
import React from 'react'
import { Spring } from 'react-spring/renderprops';

export default function Component1() {
  return (
    <div style={c1Style} >
      <h1> Component 1</h1>
      <p>Lorem30</p>
    </div>
  )
}


const c1Style = {
  backgroud: 'steelblue',
  color: 'white',
  padding: '1.5rem'
}
```

Then just add a `Spring`:

```js
export default function Component1() {
  return (
    <Spring
    from={{ opacity: 0, marginTop: -500}}
    to={{ opacity: 1, marginTop: 0}}
    >
      {props => (
        <div style={props}>
          <div style={c1Style} >
            <h1> Component 1</h1>
            <p>Lorem30</p>
          </div>
        </div>
      )}

    </Spring>

  );
}
```

`Component2` will render too, it would have state, so we use class for now:

```js
export class Component2 extends Component {
  render() {
    return (
      <Spring
      from={{ opacity: 0}}
      to={{ opacity: 1}}
      >
        {props => (
          <div style={props}>
            <div style={c2Style} >
              <h1>Component 2</h1>
              <p>Lorem40</p>
            </div>
          </div>
        )}
  
      </Spring>
    )
  }
}

const c2Style = {
  color: 'white',
  background: 'magneta',
  padding: '1.5rem'
}
```

```js
class App extends Component {
  render() {
    return( 
      <div>
        <Component1 />
        <Component2 />
      </div>
      );
  }
}
```

Then set the `config`:

```js
      <Spring
      from={{ opacity: 0}}
      to={{ opacity: 1}}
      config={{ delay: 1000, duration: 1000}}
      >
```

So now component 2 starts to fade in after a second waiting.

Let's add a counter in `Component1`:

```js
 <Spring
    from={{ opacity: 0, marginTop: -500}}
    to={{ opacity: 1, marginTop: 0}}
    >
      {props => (
        <div style={props}>
          <div style={c1Style} className="c1" >
            <h1>Component 1</h1>
            <p>Lorem30</p>

            <Spring
            from={{ number: 0}}
            to={{number: 10}}
            config={{duration: 10000}}
            >
            {props => (
              <div style={props}>
                <h1 style={counter}>
                  {props.number.toFixed()}
                </h1>
              </div>
            )}

            </Spring>
          </div>
        </div>
      )}

    </Spring>
```

## Transitions

Lets add a state to the `App` if we want to toggle transitions onClick:

```js
 state = {
    showComponent3: false
  }
```

Create `Component3` and add it into `App`.

```js
export default function Component3() {
  return (
    <div className="c3">
      <h1>Component 3</h1>
      <p>
        LoremeroL
      </p>
    </div>
  )
}
```

Add a button in C2:

```js
      <Spring
      from={{ opacity: 0}}
      to={{ opacity: 1}}
      config={{ delay: 1000, duration: 1000}}
      >
        {props => (
          <div style={props}>
            <div style={c2Style} className="c2" >
              <h1>Component 2</h1>
              <p>Lorem40</p>
              <button style={btn} onClick={this.props.toggle}>
                Toggle Component 3
              </button>
            </div>
          </div>
        )}
  
      </Spring>
```

```js
const btn = {
  background: '#333',
  color: '#fff',
  padding: '1rem 2rem',
  border: 'none',
  margin: '15px 0'
};
```

And handle it's click in the `App`:

```js
class App extends Component {
  state = {
    showComponent3: false
  }

  toggle = event => this.setState({showComponent3: !this.state.showComponent3})
  render() {
    return( 
      <div>
        <Component1 />
        <Component2  toggle={this.toggle} />
        <Component3 />
      </div>
      );
  }
}
```

Import those into app:

```js
import { Transition } from 'react-spring/renderprops';
import { animated } from 'react-spring';
```

Now instead of `<Component3 />` we use `<Transition>`

```js
class App extends Component {
  state = {
    showComponent3: false
  }

  toggle = event => this.setState({showComponent3: !this.state.showComponent3})
  render() {
    return( 
      <div>
        <Component1 />
        <Component2  toggle={this.toggle} />
        <Transition
          native
          items={this.state.showComponent3}
          from={{ opacity: 0 }}
          enter={{ opacity: 1 }}
          leave={{ opacity: 0}}
        >
         {show => show && (props => (
            <animated.div style={props}>
              <Component3 />
            </animated.div>
         ))} 
        </Transition>
      </div>
      );
  }
}
```

And that's it for basics!