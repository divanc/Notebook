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
