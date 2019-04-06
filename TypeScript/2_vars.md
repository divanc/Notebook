# TypeScript: Variable Declarations

## `var`

Grandmas of JS use `var`

```js
// As usual
var a = 10;

//Inside funcs
function f() {
  var message = "Hello, world!";
  return message;
}

// Use only inside
function f() {
  var a = 10;
  return function g() {
    var b = a + 1;
    return b;
  };
}

// Reassign funcs
var g = f();
g();
```

```js
function f() {
  var a = 1;
  a = 2;
  var b = g();
  a = 3;
  return b;
  function g() {
    return a;
  }
}

f(); // returns '2'

// x exists everywhere
function f(shouldInitialize: boolean) {
  if (shouldInitialize) var x = 10;
  return x;
}

f(true); // returns '10'
f(false); // returns 'undefined'
```

## `let`

```js
let hello = "Hello!";

function f(input: boolean) {
  let a = 100;

  if (input) {
    // Still okay to reference 'a'
    let b = a + 1;
    return b;
  }

  // Error: 'b' doesn't exist here
  return b;
}

a++; // illegal to use 'a' before it's declared;
let a;
```

```js
function foo() {
  // okay to capture 'a'
  return a;
}

// illegal call 'foo' before 'a' is declared
// runtimes should throw an error here
foo();

let a;
```

```js
// with `var` it is fine:
function f(x) {
  var x;
  var x;

  if (true) var x;
}

//with `let` – no
let x = 10;
let x = 20; // error: can't re-declare 'x' in the same scope

function f(x) {
  let x = 100; // error: interferes with parameter declaration
}

function g() {
  let x = 100;
  var x = 100; // error: can't have both declarations of 'x'
}
```

## `const`

```js
const numLivesForCat = 9; //This should not be confused with the idea that the values they refer to are immutable.

const kitty = {
  name: "Aurora",
  numLives: numLivesForCat
};

// Error
kitty = {
  name: "Danielle",
  numLives: numLivesForCat
};

// all "okay"
kitty.name = "Rory";
kitty.name = "Kitty";
kitty.name = "Cat";
kitty.numLives--;
```

You can modify `const` in TS, making insides `readonly`

## Destructuring

```ts
let input = [1, 2];
let [first, second] = input;
console.log(first); // outputs 1
console.log(second); // outputs 2

// swap variables
[first, second] = [second, first];

function f([first, second]: [number, number]) {
  console.log(first);
  console.log(second);
}
f([1, 2]);
```

```js
let [first, ...rest] = [1, 2, 3, 4];
console.log(first); // outputs 1
console.log(rest); // outputs [ 2, 3, 4 ]

let o = {
  a: "foo",
  b: 12,
  c: "bar"
};
let { a, b } = o;
```
