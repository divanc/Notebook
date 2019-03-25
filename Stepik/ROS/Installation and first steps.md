# Installation. First Steps.

Instead of repeating the course, let's install ROS to Ubuntu using [official installation guide](http://wiki.ros.org/melodic/Installation/Ubuntu)

Again, let's cut the long story short, installation is pretty simple, for some reason, official guide uses `apt`, instead of `apt-get`

1. Set up ypur antennas to listen to packages of ROS:

```console
sudo sh -c 'echo "deb http://packages.ros.org/ros/ubuntu $(lsb_release -sc) main" > /etc/apt/sources.list.d/ros-latest.list'
```

2. Set up your private channel. We don't want intruders there, do we?

```console
sudo apt-key adv --keyserver hkp://ha.pool.sks-keyservers.net:80 --recv-key 421C365BD9FF1F717815A3895523BAEEB01FA116
```

3. Make sure you are up-to-date enough to handle the power of chaos emeralds:

```console
sudo apt-get update
```

4. Dive deeply into the full version of chaos:

```console
sudo apt install ros-melodic-desktop-full
```

5. Know, where the roots of the tree are. Let your machine know them:

```console
sudo rosdep init
rosdep update
```

6. You brush your teeth daily, let your device write ROS enviroment vars into bash daily:

```console
echo "source /opt/ros/melodic/setup.bash" >> ~/.bashrc
source ~/.bashrc
```

    You may want to use zsh here, if you want

7. Now you are all set, but wait, want some DLCs?

```console
sudo apt install python-rosinstall python-rosinstall-generator python-wstool build-essential
```

## Nodes + Topics

ROS > NODE > TOPIC

Node calls messages via topics.

Let's see how two nodes can communicate with each other:

```console
roscore
```

Choosing the node, for example, `turtlesim`

Runs the node:

```console
rosrun `NODENAME` `PKGNAME`

rosrun turtlesim turtlesim_node
```

Then another one:

```console
rosrun turtlesim turtle_teleop_key
```

## How ROS Sees it

ROS sees it as a proccess list

```console
rosnode list
```

We will get list of active nodes ROS currently use.

- `/rosout` is used to exit ROS env
- then our nodes

### Variables

If we run another `turtlesim_node`, it would close first instance of the node, as that name is already used

Yet we can give another name, so ROS can distinguish and use both of these nodes:

```console
rosrun turtlesim turtlesim_node __name:=turtwo
```

Now we have two TURTLES!

### Node Info

```console
rosnpde info `/NODENAME`
```

We will get info about what packages and what nodes are subscribed to the node

Topic has a name and type: how we can call it and what it stores

Some nodes write info in the topic, whereas other get info from it. Moreover nodes which write into topics, don't know about which nodes would read it

That's why when we send info to `turtle_teleop_key` it has effect on both turtles: those turtles just subscribed to listen to the topic, to which `teleop` writes

## Topic info

In each topic in brackets we can see a type of msg, which it takes into itself.

We can read that info using

```console
rostopic info `TOPICNAME`
```

We can manually write into topic:

```
rostopic pub `TOPICNAME` `TYPE`
```

we can set the daemon to read topic updates:

```console
rostopic echo `TOPICNAME`
```

## But there is a way

package called `rqt_graph` shows all the info

```console
rosrun rqt_graph rqt_graph
```

We can see all the connections between nodes and topics

#### Or `service` — from 1 to 1:

```console
rosservice list
```

We will get maaaaany services

```console
ressoervice info `SERVICENAME`
```

```console
rossrv show `SERVICENAME`
```
