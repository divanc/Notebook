# ES6 for humans/dinos

* [ES6 по-человечески](https://habr.com/ru/post/305900/)
* And [ES-2015 Modern features](https://learn.javascript.ru/es-modern) lecture series, almost repeats in broader terms

## Table Of Contents

1. [ES6 For Humans — Habr](#es6ForHumans)
   1. [`let`,`const` and namespace](#11)
   2. [Arrow functions](#12)
   3. [Default arguments](#13)
   4. [Spread/Rest](#14)
   5. [Literals Extended](#15)
   6. [Octal & Binary Literals](#16)
   7. [Destructing array & objects](#17)
   8. [`super`](#18)
   9. [Inline separators](#19)
   10. [`for...of` vs `for..in`](#110)
   11. [`Map` & `WeakMap`](#111)
   12. [`Set` & `WeakSet`](#112)
   13. [ES6 Classes](#113)
   14. [Symbols type](#114)
   15. [Iterators](#115)
   16. [Generators](#116)
   17. [Promises](#117)
2. [Modern features of ES-2015](#2)


<a name="es6ForHumans"></a>

## 1. ES6 For Humans — Habr

There are 15 fundamental features in ES6, let's go through each one of these:

<a name="11"></a>

### `let`,`const` and namespace

`let` sets the variable within the block `{...}` — block namespace.

```js
let a = 2;
function foo() {  let a = 3;  console.log(a) } //3
console.log(a); //2
```

`const` set the variable unchangeble, however it's content's value may be changed.

```js
const arr = [1]; //1
arr.push(2); //1, 2
arr = 10; // ERR
arr[0] = 4; // 4, 2
```

`let` & `const` variables *do not exist before they are set*. These variables are block namespace.

<a name="12"></a>

### Arrow functions

Short way to write down functions

```js
let add = function(a, b) {  return a+b; }
let add = a,b => a+b;

let arr = ['apple', 'banana', 'orange'];
let breakfast = arr.map(fruit => {
    return fruit + 's';
});
```

Every function in JS sets their own `this` context, yet arrow function doesn't have their on `this`. Hence, they share outer `this`.

<a name="13"></a>

### Default arguments

ES6 lets one set default value of arguments.

```js
let getFinalPrice = (price, tax = 0.7) => price + price * tax;
getFinalPrice(500); // 850, так как значение tax не задано

```

<a name="14"></a>

### Spread/Rest

One operator, different names depending on context. While iterable, this operator spread object on seperate elements:

```js
function foo(x, y, z) {
    console.log(x, y, z);
}

let arr = [1, 2, 3];
foo(...arr); // 1 2 3
```

... or merging elements in an array:

```js
foo = ...args => return (args);
foo(1,2,3,4,5); //[1,2,3,4,5]
```

<a name="15"></a>

### Literals Extended

```js
function getCar(make, model, value) {
    return {
        make,  // === make: make
        model, // === model: model
        value, // === value: value

        // we can do maths
        ['make' + make]: true,

        // depreciate: function() {} ===
        depreciate() {
            this.value -= 2500;
        }
    };
}

let car = getCar('Kia', 'Sorento', 40000);
console.log(car);
// {
//     make: 'Kia',
//     model:'Sorento',
//     value: 40000,
//     makeKia: true,
//     depreciate: function()
// }
```

<a name="16"></a>

### Octal & Binary Literals

Adding `0o` or `0O` before integer would make it octal, adding `0b` or `0B` would make it binary

```js
let oValue = 0o10;
console.log(oValue); // 8

let bValue = 0b10;
console.log(bValue); // 2
```

<a name="17"></a>

### Destructing array & objects

It helps not to use extra variables.

```js
function foo() {  return [1, 2, 3]; }
let arr = foo(); // [1,2,3]
let [a, b, c] = foo();
console.log(a, b, c); // 1 2 3

function bar() {
    return {
        x: 4,
        y: 5,
        z: 6
    };
}
let { x: a, y: b, z: c } = bar();
console.log(a, b, c); // 4 5 6
```

<a name="18"></a>

### `super`

ES6 allows to use `super` in objects with prototypes:

```js
var parent = {
    foo() {
        console.log("Привет от Родителя!");
    }
}

var child = {
    foo() {
        super.foo();
        console.log("Привет от Ребёнка!");
    }
}

Object.setPrototypeOf(child, parent);
child.foo(); // Привет от Родителя!
             // Привет от Ребёнка!
```


<a name="19"></a>

### Inline separators

```js
let user = 'Alisa';
console.log(`Hey, ${user}!`);
```


<a name="110"></a>

### `for...of` vs `for..in`

`for...of` is being used in case you want to go through the cycle of iterable objects.

```js
let nicknames = ['di', 'boo', 'punkeye'];
nicknames.size = 3;
for (let nickname of nicknames) {
    console.log(nickname);
}
// di
// boo
// punkeye
```

`for...in` for enumerable objects:

```js
let nicknames = ['di', 'boo', 'punkeye'];
nicknames.size = 3;
for (let nickname in nicknames) {
    console.log(nickname);
}
// 0
// 1
// 2
// size
```

<a name="111"></a>

### `Map` & `WeakMap`

#### `Map`

New type of objects in JS. Every object is an example of `Map`. Every object consists of key & value, in `Map` we can use any data in these:

```js
var myMap = new Map();

var keyString = "str",
    keyObj = {},
    keyFunc = function() {};

// setting vars
myMap.set(keyString, "value, connected to 'str'");
myMap.set(keyObj, "value, connected to keyObj");
myMap.set(keyFunc, "value, connected to keyFunc");

myMap.size; // 3

// getting values
myMap.get(keyString);    // "value, connected to 'str'"
myMap.get(keyObj);       // "value, connected to keyObj"
myMap.get(keyFunc);      // "value, connected to keyFunc"
```

#### `WeakMap`

`WeakMap` is `Map` with uncertain connection. This means memory is not gonna leak. In `WeakMap` every key must be object

`WeakMap` methods:

* `delete(key)`
* `has(key)`
* `get(key)`
* `set(key, value)`

```js
let w = new WeakMap();
w.set('a', 'b');
// Uncaught TypeError: Invalid value used as weak map key

var o1 = {},
    o2 = function(){},
    o3 = window;

w.set(o1, 37);
w.set(o2, "azerty");
w.set(o3, undefined);

w.get(o3); // undefined, потому что это заданное значение

w.has(o1); // true
w.delete(o1);
w.has(o1); // false
```

<a name="112"></a>

### `Set` & `WeakSet`

#### `Set`

`Set` are object with unique elements inside;

```js
let mySet = new Set([1, 1, 2, 2, 3, 3]);
mySet.size; // 3
mySet.has(1); // true
mySet.add('string');
mySet.add({ a: 1, b:2 });

mySet.forEach((item) => {
    console.log(item);
    // 1
    // 2
    // 3
    // 'string'
    // Object { a: 1, b: 2 }
});

for (let value of mySet) {
    console.log(value);
    // 1
    // 2
    // 3
    // 'string'
    // Object { a: 1, b: 2 }
}
```

Has `delete` & `clear` methods

#### `WeakSet`

Analogy with `WeakMap`

```js
var ws = new WeakSet();
var obj = {};
var foo = {};

ws.add(window);
ws.add(obj);

ws.has(window); // true
ws.has(foo);    // false, foo is not in collection

ws.delete(window); // deletes
ws.has(window);    // false
```

<a name="113"></a>

### ES6 Classes

#### New syntax

```js
class Task {
    constructor() {
        console.log("new task created!");
    }

    showId() {
        console.log(23);
    }

    static loadAll() {
        console.log("loading tasks...");
    }
}

console.log(typeof Task); // function
let task = new Task(); // "new task created!"
task.showId(); // 23
Task.loadAll(); // "loading tasks..."
```

#### `extends` & `super`

```js
class Car {
    constructor() {
        console.log("New Auto");
    }
}

class Porsche extends Car {
    constructor() {
        super();
        console.log("new Porsche");
    }
}

let c = new Porsche();
// New Auto
// new Porsche
```

Child-class have to use `super` to use parent. Classes are seeable only after they are initted.

<a name="114"></a>

### Symbols type

Unique constant data type. Creates unique ID one cannot read. You can't do `new` with `Symbol`.

Setted in array it is hidden from view. In order to get symbol property use `Object.getOwnPropertySymbols(o)`.

```js
var sym = Symbol("symbol?");
console.log(typeof sym); // symbol

var o = {
    val: 10,
    [Symbol("random")]: "R - sy",
};

console.log(Object.getOwnPropertyNames(o)); // val
```

<a name="115"></a>

### Iterators

Iterator calls elements of collection one by one, saving its position.

```js
var arr = [11,12,13];
var itr = arr[Symbol.iterator]();

itr.next(); // { value: 11, done: false }
itr.next(); // { value: 12, done: false }
itr.next(); // { value: 13, done: false }

itr.next(); // { value: undefined, done: true }
```

<a name="116"></a>

### Generators

```js
function *infiniteNumbers() {
    var n = 1;
    while (true) {
        yield n++;
    }
}

var numbers = infiniteNumbers(); 

numbers.next(); // { value: 1, done: false }
numbers.next(); // { value: 2, done: false }
numbers.next(); // { value: 3, done: false }
```

Generator return iterable object on call. It is set with `*`, in body there must be `yield`. On call `yield` raises function

<a name="117"></a>

### Promises

Promise is new type. It is an object, which waits async operation to complete, after which returns either `fulfilled` or `rejected`

```js
var p = new Promise(function(resolve, reject) {  
    if (/* condition */) {
        resolve(/* value */);  // fulfilled successfully 
    } else {
        reject(/* reason */);  // rejected 
    }
});
```

Promises has method `then` with 2 callbacks: resolved & rejected:

```js
p.then((val) => "Success",val),
       (err) => "Error", err));
```

```js
var hello = new Promise(function(resolve, reject) {  
    resolve("Hello");
});

hello.then((str) => `${str} World`)
     .then((str) => `${str}!`)
     .then((str) => console.log(str)) // Hello World!
```

The value of promise will automatically move forward in order to avoid callbeck hell

```js
var p = new Promise(function(resolve, reject) {  
    resolve(1);
});

var eventuallyAdd1 = (val) => {
    return new Promise(function(resolve, reject){
        resolve(val + 1);
    });
}

p.then(eventuallyAdd1)
 .then(eventuallyAdd1)
 .then((val) => console.log(val)) // 3
```



## Examples

```js
function delay(ms) {
  return new Promise((resolve, reject) => {
    setTimeout(resolve, ms);
  });
}
```

would return resolved promise after `ms` ms.