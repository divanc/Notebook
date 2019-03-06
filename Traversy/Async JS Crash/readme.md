# Async JS Crash Course: Callbacks, Promises, Async Await

Imagine we have some posts:

```js
const posts = [
  { title:'Post One', body: 'This is post one'},
  { title:'Post Two', body: 'This is post Two'},
  { title:'Post Three', body: 'This is post Three'},
];
```

Dealing with servers takes time to move data between. Let's artificially recreate the server delay with `setTimeout()`:

```js
function getPosts() {
  setTimeout(() => {
    let output = '';
    posts.forEach((post,index) => {
      output += `<li>${post.title}</li>`;
    });
    document.body.innerHTML = output;
  }, 1000);
}

getPosts();
```

A second to get posts from the server? That's harsh!

What about sending to server?

```js
function createPost(post) {
  setTimeout(() => {
    posts.push(post);
  }, 2000);
}

createPost({title: 'Post Four', body: 'Post four this is'});
```

`createPost` took longer than `getPosts`, that is why we can't see anything even if we write `getPosts()` after `createPost()`.

  Let's fix it

## Callbacks

The most simple solution would be to create a **callback** function within the `createPost`. It could be done like this:

```js
function createPost(post, callback) {
  setTimeout(() => {
    posts.push(post);
    callback();
  }, 2000);
}

createPost({title: 'Post Four', body: 'Post four this is'}, getPosts)
```

That causes `getPosts` to be called right after pushing a post.

## Promises

We call `getPosts` in `then` which occurs after promise is resolved, which is successfully creating a post:

```js
function createPost(post) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      posts.push(post);
    }, 2000);

    const error = false;

    if (!error) {
      resolve();
    } else {
      reject('Something went wrong!');
    }
  });
}

createPost({title: 'Post Four', body: 'Post four this is'})
  .then(getPosts);
```

We also can manage all promises at once:

```js
const prom1 = Promise.resolve("Hello World");
const prom2 = Promise.resolve("Hello Yand");
const prom3 = new Promise((resolve, reject) => setTimeout(resolve,2000,'Goodbye'));

Promise.all([prom1,prom2,prom3])
.then((values) => console.log(values))
```

Fetch also returns promise. We have to use two `then`s so far.

## Async Await

We can also wait until `createPost` is done using this:

```js
async function init() {
  await createPost({title:'Post 4', body: '444'});
  getPosts();
}

init();
```

We can also do fetch using async/await:

```js
async function fetchData() {
  const res = await fetch(...);
  const data = await res.json();

  console.log(data);
}
```

**^ THAT IS SUPER HELPFUL, GUYS ^**