# TypeScript: Basic Types

There are:

- `Boolean`
- `Number`
- `String`
- `Array`
- `Tuple`
- `Enum`
- `Any`
- `Void`
- `Null & Undefined`
- `Never`
- `Object`

#### Boolean: true/false

```ts
let isDone: boolean = false;
```

#### Number: every kind you wanted as a frontend engineer

```ts
let decimal: number = 6;
let hex: number = 0xf00d;
let binary: number = 0b1010;
let octal: number = 0o744;
```

#### String: Hello, World!

```ts
let color: string = "blue";
color = "red";

let sentence: string = `my name is ${name}`;
sentence = "Hello, my name is " + fullName + ".";
```

#### Array

```ts
let list: number[] = [1, 2, 3];
let list: Array<number> = [1, 2, 3];
```

#### Tuple: not even like in Python

```ts
let x: [string, number];
x = [10, "hello"]; // Error
x = ["hello", 10]; // OK

console.log(x[0].substr(1)); // OK
console.log(x[1].substr(1)); // Error, 'number' does not have 'substr'
```

#### Enum: ABCDEFG

```ts
enum Color {
  Red,
  Green,
  Blue
}
let c: Color = Color.Green;

let colorName: string = Color[1];

console.log(colorName); // Displays 'Green' as its value is 1 above
```

#### Any: I don't care

```ts
let notSure: any = 4;
notSure.ifItExists(); // okay, ifItExists might exist at runtime
notSure.toFixed(); // okay, toFixed exists (but the compiler doesn't check)

let prettySure: Object = 4;
prettySure.toFixed(); // Error: Property 'toFixed' doesn't exist on type 'Object'.
```

#### Void: nothing

```ts
function warnUser(): void {
  console.log("This is my warning message");
}
```

#### Undefined and Null

```ts
let u: undefined = undefined;
let n: null = null;
```

    However, when using the --strictNullChecks flag, null and undefined are only assignable to void and their respective types. This helps avoid many common errors. In cases where you want to pass in either a string or null or undefined, you can use the union type string | null | undefined. Once again, more on union types later on

#### Never: never to occur

```ts
// Function returning never must have unreachable end point
function error(message: string): never {
  throw new Error(message);
}

// Inferred return type is never
function fail() {
  return error("Something failed");
}

// Function returning never must have unreachable end point
function infiniteLoop(): never {
  while (true) {}
}
```

#### Object

```ts
declare function create(o: object | null): void;

create({ prop: 0 }); // OK
create(null); // OK

create(42); // Error
create("string"); // Error
create(false); // Error
create(undefined); // Error
```

### Type assertions

```ts
let someValue: any = "this is a string";

let strLength: number = (<string>someValue).length;

strLength = (someValue as string).length;
```
