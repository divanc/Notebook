
# `useSpring`

```js
import { useSpring, animated } from 'react-spring'
```

Turns values into animated-values.

## Either: overwrite values to change the animation

If you re-render with changed props, the animation will update

```js
const props = useSpring({ opacity: toggle ? 1 : 0 })
```

## Or: pass a function that returns values, and update using "set"

```js
const [props, set, stop] = useSpring(() => ({ opacity: 1 }))

// Update spring with new props
set({ opacity: toggle ? 1 : 0 })
// Stop animation
stop()
```

## Finally: distribute animated props among the view

```js
return <animated.div style={props}>i will fade</animated.div>
```

## Properties

All from previous API

## Additional Notes

### To-prop shortcuts

Any property `useSpring` does not recognize will go to `to`

```js
// This ...
const props = useSpring({ opacity: 1, color: 'red' })
// is a shortcut for this ...
const props = useSpring({ to: { opacity: 1, color: 'red' } })
```

### Async scripts/chains

The `to`-property also allows you to either script your animation, or chain multiple animations together. Since these animations will execute asynchroneously make sure to provide a `from` property for base values, otherwise `props` will be empty.

#### This is a script

```js
const props = useSpring({ 
  to: async (next, cancel) => {
    await next({ opacity: 1, color: '#ffaaee' })
    await next({ opacity: 0, color: 'rgb(14,26,19)' })
  },
  from: { opacity: 0, color: 'red' } 
})
// ...
return <animated.div style={props}>I will fade in and out</animated.div>
```

#### This is a chain

```js
const props = useSpring({ 
  to: [{ opacity: 1, color: '#ffaaee' }, { opacity: 0, color: 'rgb(14,26,19)' }],
  from: { opacity: 0, color: 'red' } 
})
// ...
return <animated.div style={props}>I will fade in and out</animated.div>
```

## Demos

### [3D Card](https://codesandbox.io/s/rj998k4vmm?from-embed)

```js
const calc = (x, y) => [-(y - window.innerHeight / 2) / 20, (x - window.innerWidth / 2) / 20, 1.1]
const trans = (x, y, s) => `perspective(600px) rotateX(${x}deg) rotateY(${y}deg) scale(${s})`

function Card() {
  const [props, set] = useSpring(() => ({ xys: [0, 0, 1], config: { mass: 5, tension: 350, friction: 40 } }))
  return (
    <animated.div
      class="card"
      onMouseMove={({ clientX: x, clientY: y }) => set({ xys: calc(x, y) })}
      onMouseLeave={() => set({ xys: [0, 0, 1] })}
      style={{ transform: props.xys.interpolate(trans) }}
    />
  )
}
```

### [Flip Card](https://codesandbox.io/s/01yl7knw70?from-embed)

```js
function Card() {
  const [flipped, set] = useState(false)
  const { transform, opacity } = useSpring({
    opacity: flipped ? 1 : 0,
    transform: `perspective(600px) rotateX(${flipped ? 180 : 0}deg)`,
    config: { mass: 5, tension: 500, friction: 80 }
  })
  return (
    <div onClick={() => set(state => !state)}>
      <a.div class="c back" style={{ opacity: opacity.interpolate(o => 1 - o), transform }} />
      <a.div class="c front" style={{ opacity, transform: transform.interpolate(t => `${t} rotateX(180deg)`) }} />
    </div>
  )
}
```

### [Gesture Slider](https://codesandbox.io/s/zrj66y8714?from-embed)

```js
function Slider({ children }) {
  const [bind, { delta, down }] = useGesture()
  const { x, bg, size } = useSpring({
    x: down ? delta[0] : 0,
    bg: `linear-gradient(120deg, ${delta[0] < 0 ? '#f093fb 0%, #f5576c' : '#96fbc4 0%, #f9f586'} 100%)`,
    size: down ? 1.1 : 1,
    immediate: name => down && name === 'x'
  })
  const avSize = x.interpolate({ map: Math.abs, range: [50, 300], output: ['scale(0.5)', 'scale(1)'], extrapolate: 'clamp' })
  return (
    <animated.div {...bind()} class="item" style={{ background: bg }}>
      <animated.div class="av" style={{ transform: avSize, justifySelf: delta[0] < 0 ? 'end' : 'start' }} />
      <animated.div class="fg" style={{ transform: interpolate([x, size], (x, s) => `translate3d(${x}px,0,0) scale(${s})`) }}>
        {children}
      </animated.div>
    </animated.div>
  )
}
```

### [Mouse Parallax](https://codesandbox.io/s/r5x34869vq?from-embed)

```js
const calc = (x, y) => [x - window.innerWidth / 2, y - window.innerHeight / 2]
const trans1 = (x, y) => `translate3d(${x / 10}px,${y / 10}px,0)`
const trans2 = (x, y) => `translate3d(${x / 8 + 35}px,${y / 8 - 230}px,0)`
const trans3 = (x, y) => `translate3d(${x / 6 - 250}px,${y / 6 - 200}px,0)`
const trans4 = (x, y) => `translate3d(${x / 3.5}px,${y / 3.5}px,0)`

function Card() {
  const [props, set] = useSpring(() => ({ xy: [0, 0], config: { mass: 10, tension: 550, friction: 140 } }))
  return (
    <div class="container" onMouseMove={({ clientX: x, clientY: y }) => set({ xy: calc(x, y) })}>
      <animated.div class="card1" style={{ transform: props.xy.interpolate(trans1) }} />
      <animated.div class="card2" style={{ transform: props.xy.interpolate(trans2) }} />
      <animated.div class="card3" style={{ transform: props.xy.interpolate(trans3) }} />
      <animated.div class="card4" style={{ transform: props.xy.interpolate(trans4) }} />
    </div>
  )
}
```

### [Scroll Parallax](https://codesandbox.io/s/py912w5k6m?from-embed)

```js
function App() {
  const [{ st, xy }, set] = useSpring(() => ({ st: 0, xy: [0, 0] }))
  const interpBg = xy.interpolate((x, y) => `perspective(400px) rotateY(${x / 60}deg) rotateX(${-y / 60}deg) translate3d(-50%, -50%, 0)`)
  const interpFace = st.interpolate(o => `translate(90,${105 + o / 4})`)
  const interpEye = interpolate([st, xy], (o, xy) => `translate(${xy[0] / 30 + 157},${xy[1] / 30 + 80 + o / 2}) scale(0.8)`)
  const interpIris = interpolate([st, xy], (o, xy) => `translate(${xy[0] / 30},${xy[1] / 30 + -10 + o / 8})`)
  const interpPupil = interpolate([st, xy], (o, xy) => `translate(${xy[0] / 25},${xy[1] / 25 + -10 + o / 8})`)
  const interpSpot = interpolate([st, xy], (o, xy) => `translate(${8 + -xy[0] / 80},${-xy[1] / 80 + -10 + o / 8})`)
  const interpMouth = interpolate([st, xy], (o, xy) => `translate(${xy[0] / 18 + 188},${xy[1] / 20 + 230 + o / 1.7}) scale(0.8)`)
  const interpHair = st.interpolate(o => `translate(79,${o / 4})`)
  const onMove = useCallback(({ clientX: x, clientY: y }) => set({ xy: [x - window.innerWidth / 2, y - window.innerHeight / 2] }), [])
  const onScroll = useCallback(e => set({ st: e.target.scrollTop / 30 }), [])
  return (
    <div class="container" onMouseMove={onMove} onScroll={onScroll}>
      <div style={{ height: '700%', overflow: 'hidden' }}>{lorem({ count: 200 })}</div>
      <a.svg style={{ transform: interpBg }} viewBox="0 0 490 512">
        <g id="bg">
          <path d=".." />
        </g>
        <g id="sweater" transform="translate(94.000000, 361.000000)">
          <path d=".." />
          <path d=".." />
          <path d=".." />
        </g>
        <a.g id="face" transform={interpFace}>
          <path d=".." />
          <path d=".." />
        </a.g>
        <a.g id="eye" transform={interpEye}>
          <circle fill="#FFFFFF" cx="105" cy="104" r="104" />
          <path d=".." />
          <a.g transform={interpIris}>
            <path d=".." />
            <path d=".." />
          </a.g>
          <a.g transform={interpPupil} fill="#FFFFFF">
            <circle fill="#333031" cx="105" cy="104" r="36" />
            <path d=".." />
            <a.path
              transform={interpSpot}
              d=".."
            />
          </a.g>
        </a.g>
        <a.g id="mouth" transform={interpMouth}>
          <path d=".." />
          <path d=".." />
          <path d=".." />
          <path d=".." />
        </a.g>
        <a.g id="hair" transform={interpHair}>
          <g id="ears" transform="translate(-20.000000, 203.000000)" fill="#EFB06C">
            <path d=".." />
            <path d=".." />
          </g>
          <g fill="#794091">
            <path d=".." />
            <path d=".." />
            <path d=".." />
          </g>
          <path d=".." />
          <path d=".." />
          <path d=".." />
        </a.g>
      </a.svg>
    </div>
  )
}
```

### [Animated Burger](https://codesandbox.io/s/8ypj5vq6m9?from-embed)

```js
const items = range(4)
const interp = i => r => `translate3d(0, ${15 * Math.sin(r + (i * 2 * Math.PI) / 1.6)}px, 0)`

export default function App() {
  const { radians } = useSpring({
    to: async next => {
      while (1) await next({ radians: 2 * Math.PI })
    },
    from: { radians: 0 },
    config: { duration: 3500 },
    reset: true,
  })
  return items.map(i => <animated.div key={i} className="script-bf-box" style={{ transform: radians.interpolate(interp(i)) }} />)
}
```


### [Click](https://codesandbox.io/s/88lmnl6w88?from-embed)

```js
/*
0 % { transform: scale(1); }
25 % { transform: scale(.97); }
35 % { transform: scale(.9); }
45 % { transform: scale(1.1); }
55 % { transform: scale(.9); }
65 % { transform: scale(1.1); }
75 % { transform: scale(1.03); }
100 % { transform: scale(1); }
`*/

function Demo() {
  const [state, toggle] = useState(true)
  const { x } = useSpring({ from: { x: 0 }, x: state ? 1 : 0, config: { duration: 1000 } })
  return (
    <div onClick={() => toggle(!state)}>
      <animated.div
        style={{
          opacity: x.interpolate({ output: [0.3, 1] }),
          transform: x
            .interpolate({
              range: [0, 0.25, 0.35, 0.45, 0.55, 0.65, 0.75, 1],
              output: [1, 0.97, 0.9, 1.1, 0.9, 1.1, 1.03, 1]
            })
            .interpolate(x => `scale(${x})`)
        }}>
        click
      </animated.div>
    </div>
  )
}
```

### [Animating auto](https://codesandbox.io/s/q3ypxr5yp4?from-embed)

```js
function App() {
  const [open, toggle] = useState(false)
  const [bind, { width }] = useMeasure()
  const props = useSpring({ width: open ? width : 0 })

  return (
    <div {...bind} class="main" onClick={() => toggle(!open)}>
      <animated.div class="fill" style={props} />
      <animated.div class="content">{props.width.interpolate(x => x.toFixed(0))}</animated.div>
    </div>
  )
}
```

### [Pull & Release](https://codesandbox.io/s/r24mzvo3q?from-embed)

```js
function Pull() {
  const [{ xy }, set] = useSpring(() => ({ xy: [0, 0] }))
  const bind = useGesture(({ down, delta, velocity }) => {
    velocity = clamp(velocity, 1, 8)
    set({ xy: down ? delta : [0, 0], config: { mass: velocity, tension: 500 * velocity, friction: 50 } })
  })
  return <animated.div {...bind()} style={{ transform: xy.interpolate((x, y) => `translate3d(${x}px,${y}px,0)`) }} />
}
```

### [SVG Filter](https://codesandbox.io/s/rloj7nw3pn?from-embed)

```js
const AnimFeTurbulence = animated('feTurbulence')
const AnimFeDisplacementMap = animated('feDisplacementMap')

function App() {
  const [open, toggle] = useState(false)
  const { freq, scale, transform, opacity } = useSpring({
    reverse: open,
    from: { scale: 10, opacity: 0, transform: 'scale(0.9)', freq: '0.0175, 0.0' },
    to: { scale: 150, opacity: 1, transform: 'scale(1)', freq: '0.0, 0.0' },
    config: { duration: 3000 }
  })

  return (
    <div onClick={() => toggle(!open)}>
      <animated.svg style={{ transform, opacity }} viewBox="0 0 1278 446">
        <defs>
          <filter id="water">
            <AnimFeTurbulence type="fractalNoise" baseFrequency={freq} numOctaves="1.5" result="TURB" seed="8" />
            <AnimFeDisplacementMap xChannelSelector="R" yChannelSelector="G" in="SourceGraphic" in2="TURB" result="DISP" scale={scale} />
          </filter>
        </defs>
        <g filter="url(#water)">
          <animated.path
            d="M179.53551,113.735463 C239.115435,113.735463 292.796357,157.388081 292.796357,245.873118 L292.796357,415.764388 L198.412318,415.764388 L198.412318,255.311521 C198.412318,208.119502 171.866807,198.681098 151.220299,198.681098 C131.753591,198.681098 94.5898754,211.658904 94.5898754,264.749925 L94.5898754,415.764388 L0.205836552,415.764388 L0.205836552,0.474616471 L94.5898754,0.474616471 L94.5898754,151.489079 C114.646484,127.893069 145.321296,113.735463 179.53551,113.735463 Z M627.269795,269.469127 C627.269795,275.95803 626.679895,285.396434 626.089994,293.065137 L424.344111,293.065137 C432.012815,320.790448 457.378525,340.257156 496.901841,340.257156 C520.497851,340.257156 554.712065,333.768254 582.437376,322.560149 L608.392987,397.47748 C608.392987,397.47748 567.09997,425.202792 494.54224,425.202792 C376.562192,425.202792 325.240871,354.414762 325.240871,269.469127 C325.240871,183.343692 377.152092,113.735463 480.974535,113.735463 C574.178773,113.735463 627.269795,171.545687 627.269795,269.469127 Z M424.344111,236.434714 L528.166554,236.434714 C528.166554,216.378105 511.649347,189.242694 476.255333,189.242694 C446.17042,189.242694 424.344111,216.378105 424.344111,236.434714 Z M659.714308,0.474616471 L754.098347,0.474616471 L754.098347,415.764388 L659.714308,415.764388 L659.714308,0.474616471 Z M810.13887,0.474616471 L904.522909,0.474616471 L904.522909,415.764388 L810.13887,415.764388 L810.13887,0.474616471 Z M1097.42029,113.735463 C1191.80433,113.735463 1257.87315,183.343692 1257.87315,269.469127 C1257.87315,355.594563 1192.98413,425.202792 1097.42029,425.202792 C997.727148,425.202792 936.967423,355.594563 936.967423,269.469127 C936.967423,183.343692 996.547347,113.735463 1097.42029,113.735463 Z M1097.42029,340.257156 C1133.9941,340.257156 1163.48912,308.402543 1163.48912,269.469127 C1163.48912,230.535711 1133.9941,198.681098 1097.42029,198.681098 C1060.84647,198.681098 1031.35146,230.535711 1031.35146,269.469127 C1031.35146,308.402543 1060.84647,340.257156 1097.42029,340.257156 Z"
            fill="lightblue"
          />
        </g>
      </animated.svg>
    </div>
  )
}
```