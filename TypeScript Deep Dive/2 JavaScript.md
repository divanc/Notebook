[Back to TypeScript Deep Dive Book](https://github.com/Betra/Course-Abstract/tree/master/TypeScript%20Deep%20Dive)

# JavaScript

It is pretty big, so here is table of contents:

1. [Your JS is TS](#1)
2. [Equality](#2)
3. [References](#3)
4. [Null vs undefined](#4)
5. [this](#5)
6. [Closure](#6)
7. [Number](#7)
8. [Thruthy](#8)

<a name='1'></a>

## Your JS is TS

![](https://raw.githubusercontent.com/basarat/typescript-book/master/images/venn.png)

  TypeScript is just standardizing all the ways you provide good documentation on JavaScript.

TS is just JS with docs.

### Make JS great again

TS would protect you from JS snippets, which never work:

```ts
[] + []; // JavaScript will give you "" (which makes little sense), TypeScript will error

//
// other things that are nonsensical in JavaScript
// - don't give a runtime error (making debugging hard)
// - but TypeScript will give a compile time error (making debugging unnecessary)
//
{} + []; // JS : 0, TS Error
[] + {}; // JS : "[object Object]", TS Error
{} + {}; // JS : NaN or [object Object][object Object] depending upon browser, TS Error
"hello" - 1; // JS : NaN, TS Error

function add(a,b) {
  return
    a + b; // JS : undefined, TS Error 'unreachable code detected'
}
```

### You still need to learn JS

There are things about JS you still need to know. Let's take a look!

<a name='2'></a>

### Equality

`==` / `===` is a nightmare in JS

```js
console.log(5 == "5"); // true   , TS Error
console.log(5 === "5"); // false , TS Error
```

Choices JS make are lack meaning sometimes:

```js
console.log("" == "0"); // false
console.log(0 == ""); // true // Wonder why is that

console.log("" === "0"); // false
console.log(0 === ""); // false
```

*Note that string == number and string === number are both compile time errors in TypeScript, so you don't normally need to worry about this.*

Also there is `!=` and `!==`

**Protip**: always use `===` and `!==`, except for null checks


#### Structural Equality

Checking for structural equality those are not sufficient:

```js
console.log({a:123} == {a:123}); // False
console.log({a:123} === {a:123}); // False
```

To do so, we have to use [deep-equal](https://www.npmjs.com/package/deep-equal) npm package

```js
import * as deepEqual from "deep-equal";

console.log(deepEqual({a:123},{a:123})); // True
```

Yet in most cases it is enough to check by ids

```js
type IdDisplay = {
  id: string,
  display: string
}
const list: IdDisplay[] = [
  {
    id: 'foo',
    display: 'Foo Select'
  },
  {
    id: 'bar',
    display: 'Bar Select'
  },
]

const fooIndex = list.map(i => i.id).indexOf('foo');
console.log(fooIndex); // 0
```

<a name='3'></a>

### References

Beyond literals, any Object in JS is reference, meaning following:

#### Mutations are across all refs

```js
var foo = {};
var bar = foo; // bar is a reference to the same object

foo.baz = 123;
console.log(bar.baz); // 123
```

#### Equality is for refs

```js
var foo = {};
var bar = foo; // bar is a reference
var baz = {}; // baz is a *new object* distinct from `foo`

console.log(foo === bar); // true
console.log(foo === baz); // false
```

<a name='4'></a>

### Null vs Undefined

[Video on subject](https://www.youtube.com/watch?v=kaUfBNzuUAI)

JS (TS) has two bottom types: `null` and `undefined`. They are different:

* is not initialized: `undefined`
* unavailable: `null`

#### Checking for either

```js
/// Imagine you are doing `foo.bar == undefined` where bar can be one of:
console.log(undefined == undefined); // true
console.log(null == undefined); // true

// You don't have to worry about falsy values making through this check
console.log(0 == undefined); // false
console.log('' == undefined); // false
console.log(false == undefined); // false
```

It is recommended `==null` to check for both.

```js
function foo(arg: string | null | undefined) {
  if (arg != null) {
    // arg must be a string as `!=` rules out both null and undefined. 
  }
}
```

#### Checking for root level undefined

Don't use `==null` for root level things. 

That case, if you check for `foo` and `foo` is undefined, you'd get a `ReferenceError`.

So to check on root level you need to use `typeof`:

```js
if (typeof someglobal !== 'undefined') {
  // someglobal is now safe to use
  console.log(someglobal);
}
```

#### Limit explicit use of `undefined`

Because TS can document your structures separately. Instead of:

```js
function foo(){
  // if Something
  return {a:1,b:2};
  // else
  return {a:1,b:undefined};
}
```

type annotation is applied:

```ts
function foo():{a:number,b?:number}{
  // if Something
  return {a:1,b:2};
  // else
  return {a:1};
}
```

#### Node style callbacks

Node callbacks are typically sent with `err = null` if there are no errors. Just use it anyway:

```js
fs.readFile('someFile', 'utf8', (err,data) => {
  if (err) {
    // do something
  } else {
    // no error
  }
});
```

#### Don't use `undefined` as a means of denoting validity

An awful function:

```ts
function toInt(str:string) {
  return str ? parseInt(str) : undefined;
}
```

Can be much better:

```ts
function toInt(str: string): { valid: boolean, int?: number } {
  const int = parseInt(str);
  if (isNaN(int)) {
    return { valid: false };
  }
  else {
    return { valid: true, int };
  }
}
```

#### JSON and serialization

JSON has standard for `null`, but not `undefined`. Attribute with `undefined` would be excluded entirely:

```js
JSON.stringify({willStay: null, willBeGone: undefined}); // {"willStay":null}
```

#### Final thoughts

TS team doesn't use `null`: [TS Guidelines](https://github.com/Microsoft/TypeScript/wiki/Coding-guidelines#null-and-undefined). Yet it is not that strict.


<a name='5'></a>

### this

It all depends  on calling context.

```js
function foo() {
  console.log(this);
}

foo(); // logs out the global e.g. `window` in browsers
let bar = {
  foo
}
bar.foo(); // Logs out `bar` as `foo` was called on `bar`
```

  So be mindful of your usage of this. If you want to disconnect this in a class from the calling context use an arrow function, more on that later.

<a name='6'></a>

### Closure

Closures = <3

```js
function outerFunction(arg) {
    var variableInOuterFunction = arg;

    function bar() {
        console.log(variableInOuterFunction); // Access a variable from the outer scope
    }

    // Call the local function to demonstrate that it has access to arg
    bar();
}

outerFunction("hello closure"); // logs hello closure!
```

Inner function has access to outer variable. It is pretty intuitive

  Now the awesome part: The inner function can access the variables from the outer scope even after the outer function has returned.

```js
function outerFunction(arg) {
    var variableInOuterFunction = arg;
    return function() {
        console.log(variableInOuterFunction);
    }
}

var innerFunction = outerFunction("hello closure!");

// Note the outerFunction has returned
innerFunction(); // logs hello closure!
```

#### Reason why it's awesome

```js
function createCounter() {
    let val = 0;
    return {
        increment() { val++ },
        getVal() { return val }
    }
}

let counter = createCounter();
counter.increment();
console.log(counter.getVal()); // 1
counter.increment();
console.log(counter.getVal()); // 2
```

It makes things like __Node.js__ possible

```js
server.on(function handler(req, res) {
    loadData(req.id).then(function(data) {
        // the `res` has been closed over and is available
        res.send(data);
    })
});
```

<a name='7'></a>

### Numbers

#### Types

JavaScript has only one number type. It is a `double-precision 64-bit Number`

#### Decimal

These are not mapped correctly

```js
console.log(.1 + .2); // 0.30000000000000004
```

For **true decimal experience** use `big.js`

#### Integer

Limits are:

```js
console.log({max: Number.MAX_SAFE_INTEGER, min: Number.MIN_SAFE_INTEGER});
// {max: 9007199254740991, min: -9007199254740991}
```

*Safe* refers to the fact number can't be rounded to that value

```js
console.log(Number.MAX_SAFE_INTEGER + 1 === Number.MAX_SAFE_INTEGER + 2); // true!
console.log(Number.MIN_SAFE_INTEGER - 1 === Number.MIN_SAFE_INTEGER - 2); // true!

console.log(Number.MAX_SAFE_INTEGER);      // 9007199254740991
console.log(Number.MAX_SAFE_INTEGER + 1);  // 9007199254740992 - Correct
console.log(Number.MAX_SAFE_INTEGER + 2);  // 9007199254740992 - Rounded!
console.log(Number.MAX_SAFE_INTEGER + 3);  // 9007199254740994 - Rounded - correct by luck
console.log(Number.MAX_SAFE_INTEGER + 4);  // 9007199254740996 - Rounded!
```

in `ES6`:

```js
// Safe value
console.log(Number.isSafeInteger(Number.MAX_SAFE_INTEGER)); // true

// Unsafe value
console.log(Number.isSafeInteger(Number.MAX_SAFE_INTEGER + 1)); // false

// Because it might have been rounded to it due to overflow
console.log(Number.isSafeInteger(Number.MAX_SAFE_INTEGER + 10)); // false

```

#### big.js

  Whenever you use math for financial calculations (e.g. GST calculation, money with cents, addition etc) use a library like big.js which is designed for

* Perfect decimal math
* Safe out of bounds

```zsh
npm install big.js @types/big.js
```

```ts
import { Big } from 'big.js';

export const foo = new Big('111.11111111111111111111');
export const bar = foo.plus(new Big('0.00000000000000000001'));

// To get a number:
const x: number = Number(bar.toString()); // Loses the precision
```

#### NaN

`NaN` is returned when number calculation is not representable.

```js
console.log(Math.sqrt(-1)); // NaN
```

Instead of equality checks use `Number.isNaN`:

```js
// Don't do this
console.log(NaN === NaN); // false!!

// Do this
console.log(Number.isNaN(NaN)); // true
```

#### Infinity

```js
console.log(Number.MAX_VALUE);  // 1.7976931348623157e+308
console.log(-Number.MAX_VALUE); // -1.7976931348623157e+308

console.log(Number.MAX_VALUE + 1 == Number.MAX_VALUE);   // true!
console.log(-Number.MAX_VALUE - 1 == -Number.MAX_VALUE); // true!

console.log(Number.MAX_VALUE + 10**1000);  // Infinity
console.log(-Number.MAX_VALUE - 10**1000); // -Infinity

console.log(Number.POSITIVE_INFINITY === Infinity);  // true
console.log(Number.NEGATIVE_INFINITY === -Infinity); // true

console.log( Infinity >  1); // true
console.log(-Infinity < -1); // true
```

#### Infinitesimal

The smallest number after zero.

```js
console.log(Number.MIN_VALUE);  // 5e-324

console.log(Number.MIN_VALUE / 10);  // 0
```

<a name='8'></a>

### Truthy

```js
if (123) { // Will be treated like `true`
  console.log('Any number other than 0 is truthy');
}
```

Something that isn't truthy is called falsy.

#### Being explicit

`!!` converts variable into a Boolean

```js
// Direct variables
const hasName = !!name;

// As members of objects
const someObj = {
  hasName: !!name
}

// e.g. in ReactJS JSX
{!!someName && <div>{someName}</div>}
```