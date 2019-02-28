# Spring

```js
import { Spring } from 'react-spring/renderprops'
```

The primary task of a Spring is to move data from one state to another. The optional from-prop only plays a role when the component renders first, use the to-prop to update the spring with new values. Springs are

* accumulative, they remember all the values you ever pass
* physics based, you will never have to break your head about durations and curves again

```js
<Spring
  from={{ opacity: 0 }}
  to={{ opacity: 1 }}>
  {props => <div style={props}>hello</div>}
</Spring>
```

```js
<Spring
  from={{ number: 0 }}
  to={{ number: 1 }}>
  {props => <div>{props.number}</div>}
</Spring>
```

```js
<Spring
  from={{ value: 0 }}
  to={{ value: 100 }}>
  {props => <Donut value={props.value} />}
</Spring>
```

Otherwise works like [hooks version](../Hooks/useSpring.md).

## Demos

### [Native Spring](https://github.com/react-spring/react-spring-examples/tree/master/demos/renderprops/nativespring)

```js
export default class NativeSpringExample extends React.Component {
  state = { toggle: true }
  toggle = () => this.setState(state => ({ toggle: !state.toggle }))
  componentDidMount() {
    //setInterval(() => this.forceUpdate(), 1000)
  }
  render() {
    const toggle = this.state.toggle
    return (
      <Spring
        native
        from={{ fill: 'black' }}
        to={{
          fill: toggle ? '#247BA0' : '#70C1B3',
          backgroundColor: toggle ? '#A29B7F' : '#F3FFBD',
          rotate: toggle ? '0deg' : '180deg',
          scale: toggle ? 0.3 : 0.7,
          shape: toggle ? TRIANGLE : RECTANGLE,
        }}
        toggle={this.toggle}
        onRest={() => console.log('done')}>
        {({ toggle, backgroundColor, fill, rotate, scale, shape }) => (
          <animated.div style={{ ...styles.container, backgroundColor }}>
            <animated.svg
              style={{
                ...styles.shape,
                fill,
                transform: interpolate(
                  [rotate, scale],
                  (r, s) => `rotate3d(0,1,0,${r}) scale(${s})`
                ),
              }}
              version="1.1"
              viewBox="0 0 400 400">
              <g
                style={{ cursor: 'pointer' }}
                fillRule="evenodd"
                onClick={this.toggle}>
                <animated.path id="path-1" d={shape} />
              </g>
            </animated.svg>
          </animated.div>
        )}
      </Spring>
    )
  }
}
```

### [Morhing SVG](https://github.com/react-spring/react-spring-examples/blob/master/demos/renderprops/morph/index.js)

```js
const paths = [
  'M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z',
  'M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z',
  'M21 16v-2l-8-5V3.5c0-.83-.67-1.5-1.5-1.5S10 2.67 10 3.5V9l-8 5v2l8-2.5V19l-2 1.5V22l3.5-1 3.5 1v-1.5L13 19v-5.5l8 2.5z',
  'M7 2v11h3v9l7-12h-4l4-8z',
  'M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z',
]

const interpolators = []
for (let i = 0; i < paths.length; i++) {
  interpolators.push(
    interpolate(paths[i], paths[i + 1] || paths[0], { maxSegmentLength: 0.1 })
  )
}

export default class App extends React.Component {
  state = { interpolators, index: 0 }
  goNext = () =>
    this.setState(({ index, interpolators }) => ({
      index: index + 1 >= interpolators.length ? 0 : index + 1,
    }))
  render() {
    const { interpolators, index } = this.state
    const interpolator = interpolators[index]
    return (
      <div
        style={{
          background: '#F3FFBD',
          width: '100%',
          height: '100%',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}>
        <svg width="180" viewBox="0 0 22 22">
          <Gradient id="gradient-morph" />
          <g fill="url(#gradient-morph)">
            <Spring
              reset
              native
              from={{ t: 0 }}
              to={{ t: 1 }}
              onRest={this.goNext}>
              {({ t }) => <animated.path d={t.interpolate(interpolator)} />}
            </Spring>
          </g>
        </svg>
      </div>
    )
  }
}
```

### [Panel Keyframes](https://github.com/react-spring/react-spring-examples/tree/master/demos/renderprops/keyframes)

```js
// Creates a spring with predefined animation slots
const Sidebar = Keyframes.Spring({
  // Slots can take arrays/chains,
  peek: [
    { x: 0, from: { x: -100 }, delay: 500 },
    { x: -100, delay: 800 },
  ],
  // single items,
  open: { delay: 0, x: 0 },
  // or async functions with side-effects
  close: async call => {
    await delay(400)
    await call({ delay: 0, x: -100 })
  },
})

// Creates a keyframed trail
const Content = Keyframes.Trail({
  peek: [
    { x: 0, opacity: 1, from: { x: -100, opacity: 0 }, delay: 600 },
    { x: -100, opacity: 0, delay: 0 },
  ],
  open: { x: 0, opacity: 1, delay: 100 },
  close: { x: -100, opacity: 0, delay: 0 },
})

const items = [
  <Avatar src="https://semantic-ui.com/images/avatar2/large/elyse.png" />,
  <Input
    size="small"
    prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
    placeholder="Username"
  />,
  <Input
    size="small"
    prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
    type="password"
    placeholder="Password"
  />,
  <Fragment>
    <Checkbox size="small">Remember me</Checkbox>
    <a className="login-form-forgot" href="#" children="Forgot password" />
    <Button
      size="small"
      type="primary"
      htmlType="submit"
      className="login-form-button"
      children="Log in"
    />
  </Fragment>,
]

export default class App extends React.Component {
  state = { open: undefined }
  toggle = () => this.setState(state => ({ open: !state.open }))
  render() {
    const state =
      this.state.open === undefined
        ? 'peek'
        : this.state.open
          ? 'open'
          : 'close'
    const icon = this.state.open ? 'fold' : 'unfold'
    return (
      <div style={{ background: 'lightblue', width: '100%', height: '100%' }}>
        <Icon
          type={`menu-${icon}`}
          className="sidebar-toggle"
          onClick={this.toggle}
        />
        <Sidebar native state={state}>
          {({ x }) => (
            <animated.div
              className="sidebar"
              style={{
                transform: x.interpolate(x => `translate3d(${x}%,0,0)`),
              }}>
              <Content
                native
                items={items}
                keys={items.map((_, i) => i)}
                reverse={!this.state.open}
                state={state}>
                {(item, i) => ({ x, ...props }) => (
                  <animated.div
                    style={{
                      transform: x.interpolate(x => `translate3d(${x}%,0,0)`),
                      ...props,
                    }}>
                    <Form.Item className={i === 0 ? 'middle' : ''}>
                      {item}
                    </Form.Item>
                  </animated.div>
                )}
              </Content>
            </animated.div>
          )}
        </Sidebar>
      </div>
    )
  }
}
```

### [Auto](https://github.com/react-spring/react-spring-examples/tree/master/demos/renderprops/auto)

```js

const LOREM = `Hello world`

export default class App extends React.Component {
  state = { toggle: true, text: [LOREM] }
  onToggle = () => this.setState(state => ({ toggle: !state.toggle }))
  onAddText = () =>
    this.setState(state => ({ toggle: true, text: [...state.text, LOREM] }))
  onRemoveText = () =>
    this.setState(state => ({ toggle: true, text: state.text.slice(1) }))
  render() {
    const { toggle, text } = this.state
    return (
      <div className="auto-main">
        <button onClick={this.onToggle}>Toggle</button>
        <button onClick={this.onAddText}>Add text</button>
        <button onClick={this.onRemoveText}>Remove text</button>
        <div className="content">
          <Spring
            native
            force
            config={{ tension: 2000, friction: 100, precision: 1 }}
            from={{ height: toggle ? 0 : 'auto' }}
            to={{ height: toggle ? 'auto' : 0 }}>
            {props => (
              <animated.div className="item" style={props}>
                {text.map((t, i) => (
                  <p key={i}>{t}</p>
                ))}
              </animated.div>
            )}
          </Spring>
        </div>
      </div>
    )
  }
}
```