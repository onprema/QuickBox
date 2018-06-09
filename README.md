![QuickBox Screenshot](https://image.ibb.co/dZZy58/qq.png)


QuickBox is a web application that allows users to connect to a Linux
environment in a matter of seconds. You might find this service useful if you:
- Want to test some code without it affecting your machine
- Want an isolated development environment that can sync with your GitHub repo
- Want to get a taste of Linux without having to install it on your machine

### Get Started
Getting started is as easy as going to [QuickBox](http://138.68.8.138/) and
providing your GitHub username. You will be given instructions on how to access
your environment via ssh.

### Features
- **Import your code** by providing the URL for the GitHub repository where it
  resides. It will baked into your environment (located at the root directory `/example_repo`).

### Notes
- The containers will only be destroyed after 12 hours. Any content that you write
  in your environment should be backed up regularly. Do not rely on these
containers to be persistent.
