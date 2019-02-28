
# `useTrail`

```js
import { useTrail, animated } from 'react-spring'
```

Creates multiple springs with a single config, each spring will follow the previous one. Use it for staggered animations.

## Either: overwrite values to change the animation

```js
const trail = useTrail(number, { opacity: 1 })
```

Update on re-render

## Or: pass a function that returns values, and update using "set"

```js
const [trail, set, stop] = useTrail(number, () => ({ opacity: 1 }))

// Update trail
set({ opacity: 0 })
// Stop trail
stop()
```

## Finally: distribute animated props among the view

```js
return trail.map(props => <animated.div style={props} />) //Type: Array
```

## Demos

### [Trails](https://codesandbox.io/s/zn2q57vn13?from-embed)

```js
const items = ['Lorem', 'ipsum', 'dolor', 'sit']
const config = { mass: 5, tension: 2000, friction: 200 }

function App() {
  const [toggle, set] = useState(true)
  const trail = useTrail(items.length, {
    config,
    opacity: toggle ? 1 : 0,
    x: toggle ? 0 : 20,
    height: toggle ? 80 : 0,
    from: { opacity: 0, x: 20, height: 0 },
  })

  return (
    <div className="trails-main" onClick={() => set(state => !state)}>
      <div>
        {trail.map(({ x, height, ...rest }, index) => (
          <animated.div
            key={items[index]}
            className="trails-text"
            style={{ ...rest, transform: x.interpolate(x => `translate3d(0,${x}px,0)`) }}>
            <animated.div style={{ height }}>{items[index]}</animated.div>
          </animated.div>
        ))}
      </div>
    </div>
  )
}
```

### [World Of Goo](https://codesandbox.io/s/8zx4ppk01l?from-embed)

```js
const fast = { tension: 1200, friction: 40 }
const slow = { mass: 10, tension: 200, friction: 50 }
const trans = (x, y) => `translate3d(${x}px,${y}px,0) translate3d(-50%,-50%,0)`

export default function Goo() {
  const [trail, set] = useTrail(3, () => ({ xy: [0, 0], config: i => (i === 0 ? fast : slow) }))
  return (
    <>
      <svg style={{ position: 'absolute', width: 0, height: 0 }}>
        <filter id="goo">
          <feGaussianBlur in="SourceGraphic" result="blur" stdDeviation="30" />
          <feColorMatrix in="blur" values="1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 30 -7" />
        </filter>
      </svg>
      <div className="hooks-main" onMouseMove={e => set({ xy: [e.clientX, e.clientY] })}>
        {trail.map((props, index) => (
          <animated.div key={index} style={{ transform: props.xy.interpolate(trans) }} />
        ))}
      </div>
    </>
  )
}
```