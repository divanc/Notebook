
# Common API

## Configs

|**Property**   | **Default**   | **Description**                                                         |
|:-:            |---            |---                                                                      |
|`mass`         |1              |spring mass                                                              |
|`tension`      |170            |spring energetic load                                                    |
|`friction`     |26             |spring resistance                                                        |
|`clamp`        |false          |stops the spring once it overshoots its boundaries                       |
|`precision`    |0.01           |precision                                                                |
|`velocity`     |0              |initial velocity                                                         |
|`duration`     |undefined      |>0 will switch to a duration-based animation instead of spring physics   |

## Presets

```js
import { ..., config } from 'react-spring'

useSpring({ ..., config: config.default })
```

|**Property**        |**Value**                                   |
|:-:                 |---                                         |
|`config.default`    |`{ mass: 1, tension: 170, friction: 26 }`   |
|`config.gentle`     |`{ mass: 1, tension: 120, friction: 14 }`   |
|`config.wobbly`     |`{ mass: 1, tension: 180, friction: 12 }`   |
|`config.stiff`      |`{ mass: 1, tension: 210, friction: 20 }`   |
|`config.slow`       |`{ mass: 1, tension: 280, friction: 60 }`   |
|`config.molasses`   |`{ mass: 1, tension: 280, friction: 120 }`  |

## Properties

```js
useSpring({ from: { ... }, to: { ... }, delay: 100, onRest: () => ... })
```

All primitives share these:

|**Property**   |**Type**            |**Object**                                                                             |
|:-:            |---                 |---|
|`from`         | obj                |Base values, optional   |
|`to`           | obj/fn/array(obj)  | Animates to...  |
|`delay`        | number/fn          | Delay before anim starts. Also valid as a function for individual keys: key => delay  |
|`immediate`    | bool/fn            | Prevents animation if true. Or key => immediate  |
|`config`       | obj/fn             | Spring config. Or key => config  |
|`reset`        | bool               | Spring starts to animate from scratch  |
|`reverse`      | bool               | "from" and "to" are switched if set true, this will only make sense in combination with the "reset" flag  |
|`onStart`      | fn                 |  Callback |
|`onRest`       | fn                 | Callback  |
|`onFrame`      | fn                 | Frame-by-frame callback  |

## Interpolations

|**Value**          |**default**  |**Description**
|:-:                |---          |---|
|`extrapolateLeft`  |extend       |Left extrapolate. Can be: identity/clamp/extend
|`extrapolateRight` |extend       |Right extrapolate. Can be: identity/clamp/extend
|`extrapolate`      |extend       |Shortcut to set left and right-extrapolate
|`range`            |[0,1]        |Array of input ranges
|`output`           |undefined    |Array of mapped output ranges
|`map`              |undefined    |Value filter applied to input value
