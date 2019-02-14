[< Previous abstract](React__7.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__9.md)
----------------------- | ----------------------------|-----------------------------

## 8. Lists & Keys

Let's review how we transformed lists in JS

```js
const numbers = [1,2,3,4,5];
const doubled = numbers.map((number) => number*2);
console.log(doubled)
```

```console
>> [2, 4, 6, 8, 10]
```

In React it is similar:

### Rendering Multiple Components

`{}` 4 LIFE

```js
const  number = [1,2,3,4,5]
const liNumbers = numbers.map((number) => <li>{number}</li>);

ReactDOM.render(
  <ul>{liNumbers}</li>,
  document.getElementById('root')
);
```

### Basic List Component

Usually you want to render list inside a component

```js
function NumberList(props) {
  const numbers = props.numbers;
  const listItems = numbers.map((number) => 
    <li>{number}</li>
    );
  return (
    <ul>{listItems}</ul>
  );
}

const numbers = [1,2,3,4,5];
ReactDOM.render (
  <NumberList numbers={numbers} />,
  document.getElementById('root')
);
```

But we get a warning, that `key` is missing. Fixing this in one step

```js
// Warning here
const listItems = numbers.map((number) => 
    <li>{number}</li>
    );

// Key here
const listItems = numbers.map((number) => 
    <li key={number.toString()}>
    {number}
    </li>
    );
```

### Keys 

Keys help React to identify, which element was touched. Keys should be given to the elements of an array

The best way to pick a key is to use a string that uniquely identifies a list item among its siblings. Most often you would use IDs from your data as keys:

```js
const todoItems = todos.map((todo) =>
    <li key={todo.id}>
    {todo.text}
    </li>
  );
```

### Extracting Components With Keys

If you extract `ListItem` component, you should keep the key on the `<ListItem />` elements in the array, not on the `<li>` at all

#### Bad Example

```js
function ListItems(props) {
  const value = props.value;
  return (
    //WRONG
    <li key={value.toString()}>
    {value}
    </li>
  );
}

function NumberList(props) {
  const numbers = props.numbers;
  const listItems = numbers.map((number) => 
    //WRONG! Key should be placed here
    <ListItems value={number} />
    );
  return (
    <ul>
      {listItems}
    </ul>
  );
}

```

#### Correct Example

```js
//CORRECT! No need to specify here
function ListItems = props => <li>{props.value}</li>;

function NumberList(props) {
  const numbers = prop.numbers;
  const listItems = numbers.map((number) =>
    //Correct!
    <ListItems key={number.toString()}
                value = {number} />
    );
  return (
    <ul>
    {listItems}
    </ul>
  );
}
```

  A good rule of thumb is that elements inside the `map()` call need keys

### Keys Must Only Be Unique Among Siblings

Keys used within arrays should be unique among their siblings. However they don’t need to be globally unique. We can use the same keys when we produce two different arrays.

Keys serve as a hint to React but they don’t get passed to your components. If you need the same value in your component, pass it explicitly as a prop with a different name:

```js
const content = posts.map((post) =>
  <Post 
    key={post.id}
    id={post.id}
    title={post.title} />
    );
```

Component can read `post.id`, but can't `post.key` — it is for React purposes

### Embedding map() in JSX

Recently we did

```js
function NumberList(props) {
  const numbers = prop.numbers;
  const listItems = numbers.map((number) =>
    <ListItems key={number.toString()}
                value = {number} />
    );
  return (
    <ul>
    {listItems}
    </ul>
  );
}
```

Yet JSX allows embedding an expression into curly `{}`, so we could inline `map()` like:

```js
function NumberList(props) {
  return (
    <ul>
      {numbers.map((number) =>
      <ListItems key={number.toString()}
        value = {number} />)}
    </ul>
  );
}
// The same as 
function NumberList = props => 
  <ul>
    {numbers.map((number) =>
    <ListItems key={number.toString()}
      value = {number} />)} 
  </ul>
```

  Keep in mind that if the `map()` body is too nested, it might be a good time to extract a component.

