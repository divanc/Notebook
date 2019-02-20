# Node.js Crash Course

[This magic video](https://www.youtube.com/watch?v=fBNz5xF-Kx4) is today's topic

## What is Node?

It is **Javascript Runtime**. It is the same JS as always, yet not in browser, but in server. Built on *V8 Javascript*.

### It would help to know:

* You guess it, JS
* HTTP status
* JSON
* Arrow Functions
* Promises
* MVC Patterns

### Why Node?

* Extremly fast
* Runs on a single loop, event-driven
* non-blocking I/O model
* Popular in the industry
* Same language on the front and back

#### Non-blocking I/O

Php is synchronous, each time new thread is born, which takes CPU

* Works on a single thread using non-blocking I/O
* Supports thousands concurrent connection
* Optimizes throughout & scalability in apps with may I/O

## Node's Event Loop

* Single Threaded
* Supports concurrency via events & callbacks
* `EventEmitter` calss is used to bind events and listeners 

Node doesn't need to wait until process is complete!

### Best Types Of Projects For Node

Usually, those are not CPU intensive.

* Rest API & Microservices
* Real Time Services
* CRUD Apps
* Tools & Utils

## Installing

  NPM — Node Package Manager

* Install 3rd party packages
* Packages are stored into folder `node_modules`
* All dependencies are listed in a `package.json`
* NPM scripts can run certain tasks like run a server

```console
npm init                # Grenerates a package.json
npm install express     # Installs lockally
npm install -g nodemon  # Install globally
```

## Node Modules

* Node Core Mordules
* 3rd party modues
* Custom modules

```js
const path = require('path');
const myFile = require('./myFile')
```

# Let's Jump In!

`npm init` to start

Then package.json is set: 

```json
{
  "name": "node-js-crash",
  "version": "1.0.0",
  "description": "node js crash course tutorial",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "Ivancik",
  "license": "ISC"
}

```

`npm install uuid`  adds package, adds dependency in package.json

```json
++++++++++++++++
  "dependencies": {
    "uuid": "^3.3.2"
  }
+++++++++++++
```

Adding in index.js `console.log("Hello World")` will result in success, using a command `node index.js`.

### Moving between files

We can export modules like in `person.js`:

```js
const person = {
  name: 'John Doe',
  age: 30
}

module.exports = person;
```

And import like this, using `./` since it is not a package at this point:

```js
const person = require('./person');
console.log(person);
```

When we import node sees it in a strange way:

```js
(function (exports, require, module,__filename,__dirname){

})
```

We have access to any of these!

Importing files this way is called `CommonJS`:

```js 
const person = require('./person');
```

Whereas this is `ES6`:

```js
import { person } from './person'
```

### Path module example

```js
const path = require('path');

// Base file name
console.log(path.basename(__filename));

//Directory name
console.log(path.dirname(__filename));

// File ext
console.log(path.extname(__filename));

// Create path object
console.log(path.parse(__filename));

// Concatenate paths
console.log(path.join(__dirname,'test','hello.html'))
```

### Fs module example

File manager Fs

```js
const fs = require('fs');
const path = require('path');

// Create folder (is asynchronous) << It doesnt wait until the proccess is complete
fs.mkdir(path.join(__dirname,'/test', {}, function(err){
  console.log("Folder created");
}));

// The same with arrow
//fs.mkdir(path.join(__dirname,'/test', {}, err => {
//  if(err) throw err;
//  console.log("Folder created");
//}));

// Create folder and write to file
fs.writeFile(
  path.join(__dirname,'/test','Hello World', function(err){
    console.log("File written");
  })
);


// Create folder and write to the end of the file
fs.appendFile(
  path.join(__dirname,'/test','hello.txt'),'Hello World', function(err){
    console.log("File appended");
  });

// Read file
fs.readFile(path.join(__dirname,'/test','hello.txt'),'utf8', function(err,data){
  console.log(data);
});

// Rename file
fs.rename(path.join(__dirname,'/test','hello.txt'),path.join(__dirname,'/test','hello.txt'), function(err){
  console.log("File renamed");
});
```

### OS Module

```js
const os = require('os');

//Platform
console.log(os.platform());

//CPU Arch
console.log(os.arch());

// CPU Core Info
console.log(os.cpus());

// Free memory
console.log(os.freemem());

// Total Memory
console.log(os.totalmem());

// Home dir
console.log(os.homedir());

// Uptime
console.log(os.uptime());
```

## Url Module

```js
const url = require('url');

const myUrl = new URL('https://mysite.com/hello.html?id=100&status=active');

//Serialized URL
console.log(myUrl.href);

//Host
console.log(myUrl.host);

//Hostname (does not give port)
console.log(myUrl.hostname);

//Pathname
console.log(myUrl.pathname);

//Serialized query (after ?)
console.log(myUrl.search);

//Params object json
console.log(myUrl.searchParams);

//Add param
myUrl.searchParams.append('abc','1223');

//Loop through params
myUrl.searchParams.forEach((value, name) => console.log(`${name}: ${value}`));
```

## Events

```js
const EventEmitter = require('events');

//Create class
class MyEmmiter extends EventEmitter {}

// Init obj
const myEmmiter = new MyEmmiter();

//Event listener
myEmmiter.on('event', () => console.log('event happened!'));

//Init event
myEmmiter.EventEmitter('event');
```

### Logger

`logger.js`

```js
const EventEmitter = require('events');
const uuid = require('uuid');

class Logger extends EventEmitter {
  log(msg) {
    //call event
    this.emit('message',{id: uuid.v4(), msg });
  }
}

module.exports = Logger;
```

`index.js`

```js
const Logger = require('./logger');

const logger = new Logger();

logger.on('message', (data) => console.log('Called Listener:',data));

logger.log('Hello World');
logger.log('Hey');
logger.log('HeWoop');
```

## HTTP Server

```js
const http = require('http');

//Create server obj
http.createServer((req,res) => {
  //Write response
  res.write('Hello World');
  res.end();
}).listen(5000, () => console.log("Running..."));
```

Then on `localhost:5000` We would see that

Creating actual server is a bit harder:

```js 
const http = require('http');
const path = require('path');
const fs = require('fs');

const server = http.createServer((req, res) => {
  console.log(req.url);
});

const PORT = process.env.PORT || 5000;

server.listen(PORT, () => console.log(`Server running on port: ${PORT}`));
```

Would show browser path we are at in the console! In order to move through the pages and add HTML, we can do that:

```js
const http = require('http');
const path = require('path');
const fs = require('fs');

const server = http.createServer((req, res) => {
  if(req.url === '/') {
    res.end('<h1>Home</h1>');
  }
});

const PORT = process.env.PORT || 5000;

server.listen(PORT, () => console.log(`Server running on port: ${PORT}`));
```

In order not to restart server every update we need `nodemon` module.

We need to run `scripts` in `package.json` and fix it a little:

```json
  "scripts": {
    "start": "node index",
    "dev": "nodemon index"
  },
```

Then we could run a server with `npm run dev`....

Now we don't need to reload a server!

Yet browser doesn't consider a page as HTML, we can fix that using:

```js
const server = http.createServer((req, res) => {
  if(req.url === '/') {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.end('<h1>Homde</h1>');
  }
});
```

Passing all the HTML code in those brackets is inconvinient, instead we wanna load HTML files in folder called `public`.

Let's create `index.html` & `about.html` and fill with some typical code:

```js
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Homepage</title>
</head>
<body>
  Welcome to the Hobbit Home!
</body>
</html>
```

Same with `about`.

We want to make, if the url is `/` it adresses to `index.html`:

```js
const http = require('http');
const path = require('path');
const fs = require('fs');

const server = http.createServer((req, res) => {
  if(req.url === '/') {
    fs.readFile(path.join(__dirname, 'public', 'index.html'), (err, content) => {
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.end(content);
      }
    );
  }
  if(req.url === '/about') {
    fs.readFile(path.join(__dirname, 'public', 'about.html'), (err, content) => {
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.end(content);
      }
    );
  }
});

const PORT = process.env.PORT || 5000;

server.listen(PORT, () => console.log(`Server running on port: ${PORT}`));
```

It works with JSON like that:

```js
  if(req.url === '/api/users') {
    const users = [
    {name: 'Bob Smith', age: 40},
    {name: 'John Dow', age: 30}
    ];

    res.writeHead(200, {'Content-Type': 'application/json'});
    res.end(JSON.stringify(users));
  }
});
```

## This isn't gonna work

We can't add infinite files and css, this is a deadend, unless...

```js
  // We build a dynamic page parser
  // This builds a path to  the file
  let filePath = path.join(__dirname, 'public', req.url === '/' ? 'index.html': req.url);

  // Extension of the file to propper HTML Code 200
  let extname = path.extname(filePath);

  // Initial content type
  let contentType = 'text/html';

  // Check ext and set content type 
  switch(extname) {
    case '.js':
      contentType = 'text/javascript';
      break;
    case '.css':
      contentType = 'text/css';
      break;
    case '.json':
      contentType = 'application/json';
      break;
    case '.png':
      contentType = 'image/png';
      break;
    case '.jpeg':
      contentType = 'image/jpeg';
      break;
  }
```

And page loader with stylish 404 page you can make up yourself:

```js
  //Read File
  fs.readFile(filePath, (err, content) => {
    if (err) {
      if(err.code == "ENOENT") {
        // Page not find 404
        fs.readFile(path.join(__dirname,'public','404.html'),(err, content) => {
          res.writeHead(200, { 'Content-Type':'text/html'});
          res.end(content, 'utf8');
        });
      }
      // Some server error (500?)
      res.writeHead(500);
      res.end(`Server Error: ${err.code}`);
    } else {
      // Success
      res.writeHead(200, { 'Content-Type':'text/html'});
      res.end(content, 'utf8');
    }
  });
```

All the code would be:


```js
const http = require('http');
const path = require('path');
const fs = require('fs');

const server = http.createServer((req, res) => {
  
  // This builds a path to  the file
  let filePath = path.join(__dirname, 'public', req.url === '/' ? 'index.html': req.url);

  // Extension of the file to propper HTML Code 200
  let extname = path.extname(filePath);

  // Initial content type
  let contentType = 'text/html';

  // Check ext and set content type 
  switch(extname) {
    case '.js':
      contentType = 'text/javascript';
      break;
    case '.css':
      contentType = 'text/css';
      break;
    case '.json':
      contentType = 'application/json';
      break;
    case '.png':
      contentType = 'image/png';
      break;
    case '.jpeg':
      contentType = 'image/jpeg';
      break;
  }

  //Read File
  fs.readFile(filePath, (err, content) => {
    if (err) {
      if(err.code == "ENOENT") {
        // Page not find 404
        fs.readFile(path.join(__dirname,'public','404.html'),(err, content) => {
          res.writeHead(200, { 'Content-Type':'text/html'});
          res.end(content, 'utf8');
        });
      }
      // Some server error (500?)
      res.writeHead(500);
      res.end(`Server Error: ${err.code}`);
    } else {
      // Success
      res.writeHead(200, { 'Content-Type':contentType});
      res.end(content, 'utf8');
    }
  });
});

const PORT = process.env.PORT || 5000;

server.listen(PORT, () => console.log(`Server running on port: ${PORT}`));
```

**Now we are all set, we can do CSS, js, whatever you want, man**

## Deploy in Heroku!

You gonna need `Heroku CLI`. Then do `Heroku login`.

Commit all folder via `git` and do `heroku create`

Then go to [Dashboard](https://dashboard.heroku.com/) >> App >> Deploy and copy `heroku git:remote -a lit-springs-49763` kind of line

Do `git push`, then `heroku open`

# [TA DA](https://lit-springs-49763.herokuapp.com/)
