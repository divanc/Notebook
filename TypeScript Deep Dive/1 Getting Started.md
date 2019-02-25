[Back to TypeScript Deep Dive Book](https://github.com/Betra/Course-Abstract/tree/master/TypeScript%20Deep%20Dive)


## Getting Started

TS compiles into JS. So you are gonna need two tools: 

* TS compiler
* TS editor

### TS Version

We can write on unstable version to get more bugs, cuz it's fun, right?

```console
@ npm install -g typescript@next
```

We can ask vscode to use it by creating `.vscode/settings.json` with following:

```json
{
  "typescript.tsdk": "./node_modules/typescript/lib"
}
```

### Why TS?

There are two main goals:

* Provide an optional type system in JS
* Provide planned features from future JS editions into current JS engine

### The TS Type system

Why add types to JS?

  Types have proven ability to enhance code quality and understandability. Large teams (Google, Microsoft, Facebook) have continually arrived at this conclusion. Specifically:

* Types increase your agility when doing refactoring. It's better for the compiler to catch errors than to have things fail at runtime.
* Types are one of the best forms of documentation you can have. The function signature is a theorem and the function body is the proof.

### Your JS Is TS

TS is intentional sctrict superset of JS with optional type checking. You can rename your `.js` into `.ts` and they would still  compile as nice.

### Types Can Be Implicit

TS tries to get as much info about types, as possible, in order to make our code much more safe, yet not loading the memory.

```ts
var foo = 123;
foo = '456'; // Error: can't assign string to number
```

300 lines below in JS you would just wonder is foo a string or a number? TS won't let you worry, worring about you.

### Types Can Be Explicit

You can use annotations, so:

* Compiler would be sure of type, document will be clear, future you would understand it
* Make sure you see the code the same way compiler does

```ts
var foo: number = 123;
```

So nobody can mess with foo now!

```ts
var foo: number = '123'; // Error: can't assign string to number
```

### Types are structural

In most languages static typing is a ceremony. You have to do it, though you know code would work just as fine. However, in TS types are structural, meaning the **duck typing** is a first class language constructed

```ts
interface Point2D {
  x: number;
  y: number;
}

interface Point3D {
  x: number;
  y: number;
  z: number;
}

var Point2D: Point2D = {x:0, y:10};
var Point3D: Point3D = {x:0, y:10, z:20};

function iTakePoint(point: Point2D) { /* some magic */ }

iTakePoint2D(point2D); // exact match okay
iTakePoint2D(point3D); // extra information okay
iTakePoint2D({ x: 0 }); // Error: missing information `y`
```

### Type errors do not prevent JS emit

To make it easy to move to TS, even if there are compilation errors, by default TypeScript will emit valid JavaScript the best that it can

```ts
var foo = 123;
foo = '456'; // Error: cannot assign a `string` to a `number`
```

Will parse this:

```js
var foo = 123;
foo = '456';
```

So you can upgrade to JS instantly, TS would just warn where problematic places are.


### Types can be ambient

  A major design goal of TypeScript was to make it possible for you to safely and easily use existing JavaScript libraries in TypeScript. TypeScript does this by means of declaration. TypeScript provides you with a sliding scale of how much or how little effort you want to put in your declarations, the more effort you put the more type safety + code intelligence you get.

Defenitions for most popular libs were already written for us [here](https://github.com/borisyankov/DefinitelyTyped). So for most purposes the defenition file already  exists or you have a nice list of declaration templates.

For example, jquery. TS wants you to declare variable before using it: 

```ts
$('.awesome').show(); // Error: cannot find name `$`
```

quick fix:

```ts
declare var $: any;
$('.awesome').show(); // Okay!
```

Or you can enhance declarations to get rid of possible errors:

```ts
declare var $: {
    (selector:string): any;
};
$('.awesome').show(); // Okay!
$(123).show(); // Error: selector needs to be a string
```

### Future JS => Now

TS provides a number of features of ES6, working on today's engines of ES5. **[Then Why Babel?]** 

Here is just an example of a class:

```ts
class Point {
    constructor(public x: number, public y: number) {
    }
    add(point: Point) {
        return new Point(this.x + point.x, this.y + point.y);
    }
}

var p1 = new Point(0, 10);
var p2 = new Point(10, 20);
var p3 = p1.add(p2); // { x: 10, y: 30 }

var inc = x => x+1;
```

### Summary

  In this section we have provided you with the motivation and design goals of TypeScript. With this out of the way we can dig into the nitty gritty details of TypeScript.