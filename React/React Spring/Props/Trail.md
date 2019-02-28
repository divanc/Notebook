# Trail

```js
import { Trail } from 'react-spring/renderprops'
```

Trail animates the first item of a list of elements, the rest form a natural trail and follow their previous sibling.

```js
<Trail
  items={items} keys={item => item.key}
  from={{ transform: 'translate3d(0,-40px,0)' }}
  to={{ transform: 'translate3d(0,0px,0)' }}>
  {item => props =>
    <span style={props}>{item.text}</span>
  }
</Trail>
```

## Demos

### [Trails](https://github.com/react-spring/react-spring-examples/blob/master/demos/renderprops/trails/index.js/)

```js
export default class TrailsExample extends React.PureComponent {
  state = { toggle: true, items: ['item1', 'item2', 'item3', 'item4', 'item5'] }
  toggle = () => this.setState(state => ({ toggle: !state.toggle }))
  render() {
    const { toggle, items } = this.state
    return (
      <div
        style={{
          backgroundColor: '#247BA0',
          position: 'relative',
          width: '100%',
          height: '100%',
        }}>
        <Trail
          native
          reverse={toggle}
          initial={null}
          items={items}
          from={{ opacity: 0, x: -100 }}
          to={{ opacity: toggle ? 1 : 0.25, x: toggle ? 0 : 100 }}>
          {item =>({ x, opacity }) => (
            <animated.div
              className="box"
              onClick={this.toggle}
              style={{
                opacity,
                transform: x.interpolate(x => `translate3d(${x}%,0,0)`),
              }}
            />
          )}
        </Trail>
      </div>
    )
  }
}
```

## [Grid](https://github.com/react-spring/react-spring-examples/tree/master/demos/renderprops/grid)

File is big :c

