# HOSTZ

## PROJECT UNDER DEVELOPMENT

You know, sometimes is a pain to updates the hosts files, sometimes you need to add some new host mapping to test some project or even to avoid your browsing to go somewhere else without your consent. The good news is that your problems just ended with the new `hostz`, your command-line tool to quickly edit and improve your hosts file.

The `hostz` make easy to control your `/etc/hosts` file based on profiles, let's say you're working on a project `mycoolsite.coolz` and sometimes needs to add the `mycoolsite.coolz` on your `/etc/hosts`, well you can create a profile and load that profile when needed, and back to the old default profile again.

## Example Of Usage

```bash
# Copy the current /etc/hosts to a profile named default
hostz profile copy default /etc/hosts

# Create a new hosts file
echo "127.0.0.1 mycoolsite.coolz" > myhosts
echo "127.0.0.1 localhost" > myhosts
echo "::1       localhost" > myhosts

# Copy the new hosts file to a profile named developer
hostz profile copy developer myhosts

# Load the developer profile and do your work
hostz host use developer

# Load the default profile when you finish your work
hostz host use default
```

## Installation

Via the `go get` tool

```bash
go get -u github.com/eduardonunesp/hostz
```
