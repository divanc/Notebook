# Traversy — TypeScript Crash Course

TypeScript is a JS superset by Microsoft. It compiles to plain JS.

- Static type checking
- Class Based Objects
- Modular
- ES6

* Static typing is optional
* Helps to prevent bugs
* Makes code much more descriptive

## TS Types:

- String
- Number
- Boolean
- Array
- Any
- Void
- Null
- Undefined
- Tuple
- Enum
- Generics

## Class Based Objects

Cool feature for granddads

## TS Compiler

- Written in TS
- Compiles .ts to .js
- As NPM Package
- ES6

## Let's start!

Let's create `types.ts`:

```ts
console.log("Hello from TS!");
```

In orderto run that we need to compile it:

```console
tsc types.ts
```

Would create a js file, which would have same thing, as we didn't include any real TS into our file

Let's do some:

```ts
let str: String;
str = "Hello World";
console.log(str);
```

Would compile into:

```js
var str;
str = "Hello World";
console.log(str);
```

### Oh, by the way

You still compiling that ts every time? Actually, you can simply run compiler in autopilot:

```console
tsc types.ts -w
```

### Types

If we would assign, for example, `1` to `str`, we would get an error by the compiler, yet code would run anyway. TS is for developer, not for user

Simple things:

```ts
let str: string = "Hey";
let num: number = 1;
let bool: boolean = true;
let variable: any = 5;
```

### Arrays

For arrays `let arr: string[];` would demand `arr` to have only strings. Another way:

- `let arr: string[]`
- `let arr: Array<string>`

#### Tuple

```ts
let tuple: [string, number] = ["Hey", 1];
```

Once you've passed the pattern it may go on in any way

### Void, null, undefined

```ts
let emptyness: void = null;
let emptyness: void = undefined;
```

```ts
let zero: null = null;
let what: undefined = null;
```

## Functions

Typically in JS it works like this:

```js
const getSum = (num1, num2) => num1 + num2;
```

In TS we will set this:

```ts
const getSum = (num1: number, num2: number): number => num1 + num2;
```

Where last `number` is tracing the return type

### Optional

```ts
const setName = (firstName: string, secondName?: string): string =>
  firstName + " " + secondName;
```

Notice we use `secondName?:`, where `?` sets the var to optional;

If we would call it without `secondName`:

```ts
setName("John");
```

We would get `John undefined`. So we have to check the type within the function

### Returning void

Nice thing too

```ts
const alertMe = (message: string): void => {
  alert(message);
};
```

## Interfaces

```ts
const showTodo = (todo: { title: string; text: string }): void => {
  console.log(todo.title + " " + todo.text);
};

let myTodo = { title: "Trash", text: "Take out" };

showTodo(myTodo);
```

Yet there is a better and clearer way:

```ts
interface Todo {
  title: string;
  text: string;
}

const showTodo = (todo: Todo): void => {
  console.log(todo.title + " " + todo.text);
};

let myTodo: Todo = { title: "Trash", text: "Take out" };

showTodo(myTodo);
```

## Classes

```ts
class User {
  name: string;
  email: string;
  age: number;

  constructor(name, email, age) {
    this.name = name;
    this.email = email;
    this.age = age;
    console.log("Thank You," + name + "!");
  }
}

let John = new User("John Doe", "yuraist@icloud.com", 70);
```

Well. Just ES6 `¯\_(ツ)\_/¯`

But we can make props **private**, **public** and **protected**

### And Inheritance

### (but who wants to use those)

Same as in ES6

```ts
let Mike: User = new Member (...)
// User is a parent class
```

### Implementation

Classes use interfaces another way:

```ts
class User implements UserInterface {...}
```

And it's all. Thank you, pal
