# TypeScript: Generics

## Hello World

Without generics ping would look like:

```ts
const ping = (arg: number): number => arg;
```

And here is _type variable_ generic:

```ts
const ping<T> = (arg: T): T => (arg);
```

So `T` catches a type of user input in order to use it later on.

Once we’ve written the generic зштп function, we can call it in one of two ways. The first way is to pass all of the arguments, including the type argument, to the function:

```ts
let out = ping<string>(" PONG! ");
```

Or we can skip manual generic set:

```ts
let out = ping(" PONG! ");
```

Compiler will auto-detect type.

## More Generic Type Variables

`T` unlike any doesn't forgive your any-ness:

```ts
const aboutNumbers<T> = (arg: T): T => {
  console.log(arg.length); //ERROR: T doesn't have length
  return arg;
  }
```

Let's pretend we really want it to be arrays of `T` instead of T, then simple fix:

```ts
const aboutNumbers<T> = (arg: T[]): T[] => {
  console.log(arg.length);
  return arg;
}
```

Or like this:

```ts
const aboutNumbers<T> = (arg: Array<T>): Array<T> => {
  console.log(arg.length);
  return arg;
}
```

But hey, we can create our own generic types, like that `Array<T>`:

## Generic Types

So let's make our interface:

```ts
interface GenericPing{
    <T>(arg: T): T;
}

const ping<T> = (arg: T): T => (arg);

let pingpong: GenericPing = ping;
```

Or we can even add a generic type for our interface:

```ts
let pingpong: GenericPing<number> = ping;
```
