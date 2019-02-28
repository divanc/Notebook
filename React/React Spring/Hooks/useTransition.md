
# `useTransition`

```js
import { useTransition, animated } from 'react-spring'
```

An animated TransitionGroup. Feed it your items, keys (which can be null if items are atomic), and lifecycles. Whenever items are added or removed, it will animate these changes.

## Properties

|**Property** |**Type**   |**Description**|
|:-:          |---         |---|
|initial	    |obj/fn	     |Initial (first time) base values, optional (can be null)|
|from         |	obj/fn     |	Base values, optional|
|enter        |	obj/fn/array(obj)	|Styles apply for entering elements|
|update	      |obj/fn/array(obj)	|Styles apply for updating elements (you can update the hook itself with new values)|
|leave        |	obj/fn/array(obj)	|Styles apply for leaving elements|
|trail         |number            |	Delay in ms before the animation starts, adds up for each enter/update and leave|
|unique        |	bool/fn	        |If this is true, items going in and out with the same key will be re-used|
|reset	        |bool/fn	        |Used in combination with "unique" and makes entering items start from scratch|
|onDestroyed	  |fn	              |Called when objects have disappeared for good|

## Additional Notes

### Multi-staged transitions

The initial/from/enter/update/leave lifecycles can be objects, arrays or functions. When you provide a function you have access to individual items. The function is allowed to return plain objects, or either an array or a function for multi-stage transitions. When you provide a plain array you also can form a basic multi-stage transition (without access to the item).

```js
useTransition(items, items => items.id, {
  enter: item => [
    { opacity: item.opacity, height: item.height },
    { life: '100%' },
  ],
  leave: item => async (next, cancel) => {
    await next({ life: '0%' })
    await next({ opacity: 0 })
    await next({ height: 0 })
  },
  from: { life: '0%', opacity: 0, height: 0 },
})
```

### Transitioning routes

```js
const { location } = useRouter()
const transitions = useTransition(location, location => location.pathname, { ... })
return transitions.map(({ item, props, key }) => (
  <animated.div key={key} style={props}>
    <Switch location={item}>
      <Route path="/a" component={A} />
      <Route path="/b" component={B} />
      <Route path="/c" component={C} />
    </Switch>
  </animated.div>
))
```

## Demos

### [Simple Transition](https://codesandbox.io/s/1y3yyqpq7q?from-embed)

```js
const pages = [
  ({ style }) => <animated.div style={{ ...style, background: 'lightpink' }}>A</animated.div>,
  ({ style }) => <animated.div style={{ ...style, background: 'lightblue' }}>B</animated.div>,
  ({ style }) => <animated.div style={{ ...style, background: 'lightgreen' }}>C</animated.div>,
]

export default function App() {
  const [index, set] = useState(0)
  const onClick = useCallback(() => set(state => (state + 1) % 3), [])
  const transitions = useTransition(index, p => p, {
    from: { opacity: 0, transform: 'translate3d(100%,0,0)' },
    enter: { opacity: 1, transform: 'translate3d(0%,0,0)' },
    leave: { opacity: 0, transform: 'translate3d(-50%,0,0)' },
  })
  return (
    <div className="simple-trans-main" onClick={onClick}>
      {transitions.map(({ item, props, key }) => {
        const Page = pages[item]
        return <Page key={key} style={props} />
      })}
    </div>
  )
}
```

### [Image Fade](https://codesandbox.io/s/morr206pv8?from-embed)

```js
const slides = [
  { id: 0, url: 'photo-1544511916-0148ccdeb877?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1901&q=80i' },
  { id: 1, url: 'photo-1544572571-ab94fd872ce4?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1534&q=80' },
  { id: 2, url: 'reserve/bnW1TuTV2YGcoh1HyWNQ_IMG_0207.JPG?ixlib=rb-1.2.1&w=1534&q=80' },
  { id: 3, url: 'photo-1540206395-68808572332f?ixlib=rb-1.2.1&w=1181&q=80' },
]

const App = () => {
  const [index, set] = useState(0)
  const transitions = useTransition(slides[index], item => item.id, {
    from: { opacity: 0 },
    enter: { opacity: 1 },
    leave: { opacity: 0 },
    config: config.molasses,
  })
  useEffect(() => void setInterval(() => set(state => (state + 1) % 4), 2000), [])
  return transitions.map(({ item, props, key }) => (
    <animated.div
      key={key}
      class="bg"
      style={{ ...props, backgroundImage: `url(https://images.unsplash.com/${item.url}&auto=format&fit=crop)` }}
    />
  ))
}
```

### [Multistaged transition](https://codesandbox.io/s/vqpqx5vrl0?from-embed)

```js
n App() {
  const ref = useRef([])
  const [items, set] = useState([])
  const transitions = useTransition(items, null, {
    from: { opacity: 0, height: 0, innerHeight: 0, transform: 'perspective(600px) rotateX(0deg)', color: '#8fa5b6' },
    enter: [
      { opacity: 1, height: 80, innerHeight: 80 },
      { transform: 'perspective(600px) rotateX(180deg)', color: '#28d79f' },
      { transform: 'perspective(600px) rotateX(0deg)' },
    ],
    leave: [{ color: '#c23369' }, { innerHeight: 0 }, { opacity: 0, height: 0 }],
    update: { color: '#28b4d7' },
  })

  const reset = useCallback(() => {
    ref.current.map(clearTimeout)
    ref.current = []
    set([])
    ref.current.push(setTimeout(() => set(['Apples', 'Oranges', 'Kiwis']), 2000))
    ref.current.push(setTimeout(() => set(['Apples', 'Kiwis']), 5000))
    ref.current.push(setTimeout(() => set(['Apples', 'Bananas', 'Kiwis']), 8000))
  }, [])

  useEffect(() => void reset(), [])

  return (
    <div>
      {transitions.map(({ item, props: { innerHeight, ...rest }, key }) => (
        <animated.div className="transitions-item" key={key} style={rest} onClick={reset}>
          <animated.div style={{ overflow: 'hidden', height: innerHeight }}>{item}</animated.div>
        </animated.div>
      ))}
    </div>
  )
}
```

### [Notification](https://codesandbox.io/s/7mqy09jyq?from-embed)

```js
let id = 0

function MessageHub({ config = { tension: 125, friction: 20, precision: 0.1 }, timeout = 3000, children }) {
  const [refMap] = useState(() => new WeakMap())
  const [cancelMap] = useState(() => new WeakMap())
  const [items, setItems] = useState([])
  const transitions = useTransition(items, item => item.key, {
    from: { opacity: 0, height: 0, life: '100%' },
    enter: item => async next => await next({ opacity: 1, height: refMap.get(item).offsetHeight }),
    leave: item => async (next, cancel) => {
      cancelMap.set(item, cancel)
      await next({ life: '0%' })
      await next({ opacity: 0 })
      await next({ height: 0 })
    },
    onRest: item => setItems(state => state.filter(i => i.key !== item.key)),
    config: (item, state) => (state === 'leave' ? [{ duration: timeout }, config, config] : config),
  })

  useEffect(() => void children(msg => setItems(state => [...state, { key: id++, msg }])), [])
  return (
    <Container>
      {transitions.map(({ key, item, props: { life, ...style } }) => (
        <Message key={key} style={style}>
          <Content ref={ref => ref && refMap.set(item, ref)}>
            <Life style={{ right: life }} />
            <p>{item.msg}</p>
            <Button
              onClick={e => {
                e.stopPropagation()
                cancelMap.has(item) && cancelMap.get(item)()
              }}>
              <X size={18} />
            </Button>
          </Content>
        </Message>
      ))}
    </Container>
  )
}

export default function App() {
  const ref = useRef(null)
  return (
    <Main onClick={() => ref.current(lorem())}>
      <GlobalStyle />
      Click here to create notifications
      <MessageHub children={add => (ref.current = add)} />
    </Main>
  )
}
```

### [Mansory Grid](https://codesandbox.io/s/26mjowzpr?from-embed)

```js
function App() {
  const columns = useMedia(['(min-width: 1500px)', '(min-width: 1000px)', '(min-width: 600px)'], [5, 4, 3], 2)
  const [bind, { width }] = useMeasure()
  const [items, set] = useState(data)
  useEffect(() => void setInterval(() => set(shuffle), 2000), [])

  let heights = new Array(columns).fill(0) // Each column gets a height starting with zero
  let gridItems = items.map((child, i) => {
    const column = heights.indexOf(Math.min(...heights)) // Basic masonry-grid placing, puts tile into the smallest column using Math.min
    const xy = [(width / columns) * column, (heights[column] += child.height / 2) - child.height / 2] // X = container width / number of columns * column index, Y = it's just the height of the current column
    return { ...child, xy, width: width / columns, height: child.height / 2 }
  })

  // This turns gridItems into transitions, any addition, removal or change will be animated
  const transitions = useTransition(gridItems, item => item.css, {
    from: ({ xy, width, height }) => ({ xy, width, height, opacity: 0 }),
    enter: ({ xy, width, height }) => ({ xy, width, height, opacity: 1 }),
    update: ({ xy, width, height }) => ({ xy, width, height }),
    leave: { height: 0, opacity: 0 },
    config: { mass: 5, tension: 500, friction: 100 },
    trail: 25
  })

  return (
    <div {...bind} class="list" style={{ height: Math.max(...heights) }}>
      {transitions.map(({ item, props: { xy, ...rest }, key }) => (
        <a.div key={key} style={{ transform: xy.interpolate((x, y) => `translate3d(${x}px,${y}px,0)`), ...rest }}>
          <div style={{ backgroundImage: item.css }} />
        </a.div>
      ))}
    </div>
  )
}
```