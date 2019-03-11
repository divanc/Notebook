# Redux: The Immutable State Tree

**First principle of Redux: no matter how hard your app is, you save all your states in one JS object**.

All mutations are explicit

## Example

With one counter we just hold it's value in state and it's fine, however with several counters we want to give each one an ID in state to make sure which is being changed

If we have a todo app, we want to store all aspects in one object, so if a task is done, it is cheched within an array of this tash within an array of all tasks \*it is called a **state tree\***

### We Describe State Changes with _Actions_

It is a js object, the minimum representation of state:

```js
[object Object] { /// THIS IS STATE TREE
  type: 'INCREMENT' // THIS IS ONE ACTION
}
```

for counter we need only `increment` and `decrement` for adding or substracting.

For several we have to add an `id` to every counter.

**Second dogma**: Redux can only change state via _actions_

### Pure and Impure functions

Pure functions return values depending only on arguments. They are predictable. They do not overwrite arguments.

Impure can call dbs, overwrite and do stuff.

In Redux we want pure funcs
