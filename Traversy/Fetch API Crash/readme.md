# Fetch API Introduction

## Part 1: txt

In basic html, lets construct a fundament:

```html
  <button id="getText">Get Text</button>

  <script>
    document.getElementById('getText').addEventListener('click',getText);
    function getText() {
      alert();
    }
  </script>
  ```

Replacing `alert` with a `fetch`, we have to acknoledge, that fetch returns a promise.

```html
  <script>
    document.getElementById('getText').addEventListener('click',getText);
    function getText() {
      fetch('sample.txt')
      .then(function(res) {
        console.log(res)
      })
    }
  </script>
```

What we get in console is response:

```json
Response {type: "basic", url: "http://127.0.0.1:5500/sample.txt", redirected: false, status: 200, ok: true, …}
body: (...)
bodyUsed: false
headers: Headers {}
ok: true
redirected: false
status: 200
statusText: "OK"
type: "basic"
url: "http://127.0.0.1:5500/sample.txt"
__proto__: Response
```

With `.txt` `res.text()` will give us the filedata.

Best thing about promises — we can build ultimate function chain

```html
  <script>
    document.getElementById('getText').addEventListener('click',getText);
    function getText() {
      fetch('sample.txt')
      .then(function(res) {
        return res.text();
      })
      .then(function(data) {
        console.log(data)
      })
    }
  </script>
```

We can clean up, using arrow functions:

```js
      fetch('sample.txt')
      .then((res) => res.text())
      .then((data) => console.log(data));
```

Let's place file data in the DOM, inside *outside* div

```js
    document.getElementById('getText').addEventListener('click',getText);
    function getText() {
      fetch('sample.txt')
      .then((res) => res.text())
      .then((data) => {
        document.getElementById('output').innerHTML = data;
      });
    }
```

For catching errors, just place cather in the end:

```js
fetch('sample.txt')
.then((res) => res.text())
.then((data) => {
  document.getElementById('output').innerHTML = data;
})
.catch((error) => aler(error));
```

## Part 2: JSON

Building same ol' button

```js
  document.getElementById('getUsers').addEventListener('click',getUsers);
```

```js
function getUsers() {
  fetch("users.json")
  .then((res) => res.json())
  .then((data) => {
    let output = '<h2>Users</h2>';
    data.forEach(function(user){
      output +=`
      <ul>
        <li>ID: ${user.id}</li>
        <li>Name: ${user.name}</li>
        <li>Email: ${user.email}</li>
      </ul>
      `;
    });
    document.getElementById('output').innerHTML = output;
  })
}
```

## Part 3: API


The very same button and event listener, function is:

```js
function getPosts() {
  fetch("https://jsonplaceholder.typicode.com/posts")
  .then((res) => res.json())
  .then((data) => {
    let output = '<h2>Posts</h2>';
    data.forEach(function(post){
      output +=`
        <section>
          <h2> ${post.title} </h2>
          <p> ${post.body} </p>
      `;
    });
    document.getElementById('output').innerHTML = output;
  });
}
```

## Part 4: Add to API

Let's add the feature of sending to api. HTML is:

```html
  <button id="getText">Get Text</button>
  <button id="getUsers">Get JSON</button>
  <button id="getPosts">Get API Data</button>
  <hr>
  <div id="output"></div>
  <hr>
  <form id="addPost">
    <div>
      <input type='text' id='title' placeholder="title" />
    </div>
    <div>
      <textarea placeholder="Body" id="body">

      </textarea>
    </div>
    <input type="submit" />
  </form>
  
```

Adding a listener

```js
document.getElementById('addPost').addEventListener('submit',addPost);
```

And function will look like:

```js
function addPost(event) {
  event.preventDefault();

  let title = document.getElementById('title').value;
  let body = document.getElementById('body').body;

  fetch('https://jsonplaceholder.typicode.com/posts', {
    method: 'POST',
    headers: {
      'Accept': 'application/json, text/plain, */*',
      'Content-type': 'application/json'
    },
    body: JSON.stringify({title:title, body:body})
  })
  .then((res) => res.json())
  .then((data) => console.log(data))
  .catch((error) => console.log(error))
}
```