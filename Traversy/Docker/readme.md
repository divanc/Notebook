# Docker 

`docker info` gives info

## Starting

Let's create our first nginx container

```console
docker container run -it -p 80:80 nginx
```

- `-it` stands for interactive 
- `-p` stands for publish
- `80` is default nginx port
- `nginx` — image name

Second `80` is always `80` for nginx and apache, for MySQL, for example, it would be `3306`

Now, as it can't find `nginx` locally, it would download it from *DockerHub* by itself.

That's it! now `localhost:80` is running!

## Moving around

You can get your hands on the [hub](https://hub.docker.com/search?q=nginx&type=image) and look for images there.

Searching for `nginx` we may find a [**Dockerfile**](https://github.com/nginxinc/docker-nginx/blob/e5123eea0d29c8d13df17d782f15679458ff899e/mainline/stretch/Dockerfile) there.

So what we did in the *Starting* was simply pulling Dockerfile with nginx instructions and executing it in the image.

- `docker container ls` would show currently running containers, however, with flag `-a` would show all containers in the system.
- `docker container rm @first 3 symbols of ID@` would erase container
- `docker images` even if we deleted nginx container, image is still in our system, `docker images` would show that
- `docker pull @REPO@` would download image manually from hub

## New try

```console
docker container run -d -p 8080:80 --name mynginx nginx
```

`-d` stands for detached, would run in the background

`docker ps` === `docker container ls`

Then, let's create an Apache container

```console
docker container run -d -p 8081:80 --name myapache httpd
```

And MySQL

```console
docker container run -d -p 3306:3306 --name mysql --env MYSQL_ROOT_PASSWORD=123456 mysql
```

`docker container stop` would cease process of the container without deleting.

To delete running container we should force it via flag `-f`.

### exec 

In order to change files in the container we should open it via bash

```terminal
docker container exec -it mynginx bash
```

Here, moving to `usr/share/nginx/html` we can find our index.html

Well, we don't need to edit it inside bash, that would be madness.

For now, let's remove all containers: `docker rm $(docker ps -aq) -f`

## Third attempt

```terminal
docker container run -d -p 8080:80 -v $(pwd):/usr/share/nginx/html --name nginx-site nginx
```

What we did by that is we binded local current folder to `/nginx/html` folder of the image. Now we can do stuff in that local folder and it would be sent to the container.

Saving `index.html` in that directory would overwrite container's one and *voila!*

## Dockerfile

In order to create image from that application we need to create `Dockerfile`:

```Dockerfile
FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY . .
```

Now do

```terminal
docker image build -t divando/nginx-site .
```

after `t` *divando* is DockerHub username and `nginx-site` is name of the container
`.` is important, it tells where the **Dockerfile** is.

In order to build such images we don't actually need to create what we did before. Now lets junk it out.


It's time to run our image:

```terminal
docker container run -d -p 8082:80 divando/nginx-site
```

And let's push it to our DockerHub!

```terminal
docker push divando/nginx-site
```

# Lection 2: Real Deal

We download mongo + node app from prepared repo.

In the root let's create Dockerfile:

```Dockerfile
FROM node:10

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm","start"]
```

That should work! Now, Mongo is still lacking here and we probably should create another container for db. In order to connect those we use **docker-compose.yml** . Create that!

```yaml
version: '3'
services:
  app:
    container_name:  docker_node_mongo
    restart: always
    build: .
    ports:
      - '80:3000'
    links:
      - mongo
  mongo:
    container_name: mongo
    image: mongo
    ports:
      - '27017:27017'
```

- `restart: always` means it would restart every time on fail
- `build: .` means it would be built by the Dockerfile in this directory
- `links: ` connects containers
- `image: mongo` as mongo doesn't have a Dockerfile, we have to show, which image it should use, mongo will be pulled from DockerHub


Last would be `.dockerignore`

```
node_modules
npm-debug.log 
```

Than let's build!

```terminal
docker-compose up
```

If all set, you can deconstruct the image

```terminal
docker-compose down
```

## Dropleting

Let's upload that to a droplet!

There are many ways to do that. However, let's do that using Github.
Make a repo with our project and push it.


Now create a droplet, git clone into home our repo and just do `docker-compose-up`


    Insight: to delete all images you can use `docker rmi $(docker images -q)`

That's it!