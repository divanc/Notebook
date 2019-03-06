const posts = [
  { title:'Post One', body: 'This is post one'},
  { title:'Post Two', body: 'This is post Two'},
  { title:'Post Three', body: 'This is post Three'},
];

function getPosts() {
  setTimeout(() => {
    let output = '';
    posts.forEach((post,index) => {
      output += `<li>${post.title}</li>`;
    });
    document.body.innerHTML = output;
  }, 1000);
}

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

async function init() {
  await createPost({title:'Post 4', body: '444'});
  getPosts();
}

init();