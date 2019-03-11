# Refacturing The Refactured Todo!

## FilterLink & Footer

We have to pass a lot of components in the tree :c

Check it out yourself:

```jsx
export const FilterLink = ({ filter, children, currentFilter, onClick }) => {
  if (currentFilter === filter) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick(filter);
      }}
    >
      {children}
    </a>
  );
};
```

And father-`Footer` needs to know all that and his papa needs to know everything was mentioned, kinda a lot.

## Let's add more containers!

Cut all props from the `Footer`:

```jsx
export const Footer = () => (
  <p>
    Show: {"   "}
    <FilterLink filter="SHOW_ALL">All</FilterLink> {"  "}
    <FilterLink filter="SHOW_ACTIVE">Active</FilterLink> {"  "}
    <FilterLink filter="SHOW_COMPLETED">Completed</FilterLink> {"  "}
  </p>
);
```

And our `FilterLink` to be honest not as presentational as it could be:

It was:

```jsx
export const FilterLink = ({ filter, children, currentFilter, onClick }) => {
  if (currentFilter === filter) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick(filter);
      }}
    >
      {children}
    </a>
  );
};
```

Let's cut it's functions and make it just `Link`:

```jsx
export const Link = ({ active, children, onClick }) => {
  if (active) return <span>{children}</span>;
  return (
    <a
      href="#"
      onClick={e => {
        e.preventDefault();
        onClick();
      }}
    >
      {children}
    </a>
  );
};
```

Well, it looks presentational as hell

However, container for link would be more complex:

```jsx
export class FilterLink extends Component {
  componentDidMount() {
    this.unsubscribe = store.subscribe(() => this.forceUpdate());
  }

  componentWillUnmount() {
    this.unsubscribe();
  }
  render() {
    const props = this.props;
    const state = store.getState();

    return (
      <Link
        active={props.filter === state.visibilityFilter}
        onClick={() => {
          store.dispatch({
            type: "SET_VISIBILITY_FILTER",
            filter: props.filter
          });
        }}
      >
        {props.children}
      </Link>
    );
  }
}
```
