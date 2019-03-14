| [< Previous abstract](5%20React%20Counter.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/blob/master/Egghead/Dan%20Abramov%20-%20Redux/) | [Next abstract >](6%20Avoiding%20Object%20Mutations.md) |
| --------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |


# Avoiding Array Mutations

### `concat()` is good

If we want to add zero in the past array, we never want to use `append()`: it mutates!

```js
const addCounter = list => {
  return list.concat([0]); // makes a new array
};
```

Another way to write it in ES6:

```js
const addCounter = list => {
  return [...list, 0];
};
```

### `splice()` is bad, sorry

Because it mutates an array, instead we would use:

```js
const removeCounter = (list, index) => {
  return list.slice(0, index).concat(list.slice(index + 1));
};
```

In ES6:

```js
const removeCounter = (list, index) => {
  return [...list.slice(0, index), ...list.slice(index + 1)];
};
```

### Adding

This is mutating our list, not what we really want

```js
const addIncrement = (list, index) => {
  list[index]++;
  return list;
};
```

Instead, let's try it out:

```js
const addIncrement = (list, index) => {
  return list
    .slice(0, index) // We want to take our list just before the position
    .concat([list[index] + 1]) // Insert a value
    .concat(list.slice(index + 1)); // And add the rest
};
```

In ES6:

```js
const addIncrement = (list, index) => {
  return [...list.slice(0, index), list[index] + 1, ...list.slice(index + 1)];
};
```
