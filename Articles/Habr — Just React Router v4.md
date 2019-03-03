# Prostoi Tutorial React Router v4

[This](https://habr.com/ru/post/329996/) little boy.

**Router v4** is now all components.

Now it is split in 3 packages:

* `react-router`
* `router-dom`
* `react-router-native`


`react-router` provides two areas: Browser and native

Our site is in browser, so let's use `react-router-dom`.

```console
npm install --save react-router-dom
```

## Router

On project init you should choose one router. For browsering there are two components: `BrowserRouter` & `HashRouter`

`BrowserRouter` for dynamic purposes, otherwise `Hash..`

## History

Each Router creates `history` which contains path  `location[1]`

## Rendering Router

`Router` wants just one child element. You proobably want just render `App` inside:

```js
import { BrowserRouter } from 'react-router-dom';

ReactDOM.render((
  <BrowserRouter>
    <App />
  </BrowserRouter>
), document.getElementById('root'))
```

## App

This is gonna be rendered

```js
const App = () => (
  <div>
    <Header />
    <Main />
  </div>
)
```

## Routes

Main building block of roter is `Route`. It renders depending in URL:

### Path

Route eats path as a prop with the location

```js
<Route path='/roster' />
```

```js
<Route path='/roster'/>
// location.pathname == '/' ? prop path doesn't match
// location.pathname == '/roster' || '/roster/2' prop path matches
// exact prop: matches on === '/roster',not '/roster/2'
<Route exact path='/roster'/>
```

### Path match

`path-to-regexp` package generates regexp and match it to location.pathname

`match` object has props:

* `url` — what is being matched
* `path` — path in `Route`
* `isExact` — for exact ===
* `params`

## Creating Routes

`<Switch />` for several routes within router. Lets pretend we have these paths:

* **/**
* **/roster** — Game teams
* **/roster/:number** — Page of a team
* **/schedule** — schedule

```js
<Switch>
  <Route exact path='/' component={Home}/>
  {/* Both roster and number starts with a roster */}
  <Route path='/roster' component={Roster}/>
  <Route path='/schedule' component={Schedule}/>
</Switch>
```

## What does rendering a `Route`?

Route has 3 props:

* **component**: uses React.createElement
* **render**: pretty much like `component`, yet extra features, inline
* **children**: always showed, no matter the path

```js
<Route path='/page' component={Page} />
const extraProps = { color: 'red' }
<Route path='/page' render={(props) => (
  <Page {...props} data={extraProps}/>
)}/>

<Route path='/page' children={(props) => (
  props.match
    ? <Page {...props}/>
    : <EmptyPage {...props}/>
)}/>
```


## Main

Main Component:

```js
import { Switch, Route } from 'react-router-dom'
const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route path='/roster' component={Roster}/>
      <Route path='/schedule' component={Schedule}/>
    </Switch>
  </main>
)
```

/roster/:number is excluded from `<Switch/>`: it would be rendered by `Roster` element, which is rendered for any path starts from `/roster/`.

In Roster we create:

* **/roster** — with prop exact
* **/roster/:number** — which catches numbers

```js
const Roster = () => (
  <Switch>
    <Route exact path='/roster' component={FullRoster}/>
    <Route path='/roster/:number' component={Player}/>
  </Switch>
)
```

We can preload some shared content beforehand:

```js
const Roster = () => (
  <div>
    <h2>This is a roster page!</h2>
    <Switch>
      <Route exact path='/roster' component={FullRoster}/>
      <Route path='/roster/:number' component={Player}/>
    </Switch>
  </div>
)
```

### Path params

`:number` means that this part of the path would be saved in `match.params.number`. For ex. **/roster/6** would give:

```js
{ number: '6'
```

`<Player />` would use `props.match.params.number` for render info.

```js
// API возращает информацию об игроке в виде объекта
import PlayerAPI from './PlayerAPI'
const Player = (props) => {
  const player = PlayerAPI.get(
    parseInt(props.match.params.number, 10)
  )
  if (!player) {
    return <div>Sorry, but the player was not found</div>
  }
  return (
    <div>
      <h1>{player.name} (#{player.number})</h1>
      <h2>{player.position}</h2>
    </div>
)
```

Some extra components will be: 

```js
const FullRoster = () => (
  <div>
    <ul>
      {
        PlayerAPI.all().map(p => (
          <li key={p.number}>
            <Link to={`/roster/${p.number}`}>{p.name}</Link>
          </li>
        ))
      }
    </ul>
  </div>
)
const Schedule = () => (
  <div>
    <ul>
      <li>6/5 @ Спартак</li>
      <li>6/8 vs Зенит</li>
      <li>6/14 @ Рубин</li>
    </ul>
  </div>
)
const Home = () => (
  <div>
    <h1>Добро пожаловать на наш сайт!</h1>
  </div>
)
```

## Links

To do links without external loading process, there is a `<Link>.

```js
import { Link } from 'react-router-dom'
const Header = () => (
  <header>
    <nav>
      <ul>
        <li><Link to='/'>Home</Link></li>
        <li><Link to='/roster'>Roster</Link></li>
        <li><Link to='/schedule'>Schedule</Link></li>
      </ul>
    </nav>
  </header>
)
```

Paths must be absolute. They could be string or location [it's hard, let it go]

