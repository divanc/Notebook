
# `useChain`

```js
import { useChain, animated } from 'react-spring'
```

Set the execution order of previously defined animation-hooks, where one animation starts after the other in sequence. You need to collect refs off the animations you want to chain, which blocks the animation from starting on its own. The order can be changed in subsequent render passes.

```js
// Build a spring and catch its ref
const springRef = useRef()
const props = useSpring({ ...values, ref: springRef })
// Build a transition and catch its ref
const transitionRef = useRef()
const transitions = useTransition({ ...values, ref: transitionRef })
// First run the spring, when it concludes run the transition
useChain([springRef, transitionRef])
// Use the animated props like always
return (
  <animated.div style={props}>
    {transitions.map(({ item, key, props }) =>
      <animated.div key={key} style={props} />)}
  </animated.div>
)
```

## Demos

### [Chain Animation](https://codesandbox.io/s/2v716k56pr?from-embed)

```js
export default function App() {
  const [open, set] = useState(false)

  const springRef = useRef()
  const { size, opacity, ...rest } = useSpring({
    ref: springRef,
    config: config.stiff,
    from: { size: '20%', background: 'hotpink' },
    to: { size: open ? '100%' : '20%', background: open ? 'white' : 'hotpink' }
  })

  const transRef = useRef()
  const transitions = useTransition(open ? data : [], item => item.name, {
    ref: transRef,
    unique: true,
    trail: 400 / data.length,
    from: { opacity: 0, transform: 'scale(0)' },
    enter: { opacity: 1, transform: 'scale(1)' },
    leave: { opacity: 0, transform: 'scale(0)' }
  })

  // This will orchestrate the two animations above, comment the last arg and it creates a sequence
  useChain(open ? [springRef, transRef] : [transRef, springRef], [0, open ? 0.1 : 0.6])

  return (
    <>
      <Global />
      <Container style={{ ...rest, width: size, height: size }} onClick={() => set(open => !open)}>
        {transitions.map(({ item, key, props }) => (
          <Item key={key} style={{ ...props, background: item.css }} />
        ))}
      </Container>
    </>
  )
}
```