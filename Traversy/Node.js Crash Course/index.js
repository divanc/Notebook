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