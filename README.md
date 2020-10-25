# HOSTZ

## PROJECT UNDER DEVELOPMENT

You know, sometimes is a pain to updates the hosts files, sometimes you need to add some new host mapping to test some project or even to avoid your browsing to go somewhere else without your consent. The good news is that your problems just ended with the new `hostz`, your command-line tool to quickly edit and improve your hosts file.

The `hostz` make easy to control your `/etc/hosts` file based on profiles, let's say you're working on a project `mycoolsite.coolz` and sometimes needs to add the `mycoolsite.coolz` on your `/etc/hosts`, well you can create a profile and load that profile when needed, and back to the old default profile again.

## Example Of Usage

```bash
# Load the default profile
hostz profile copy default /etc/hosts

# Create a new hosts file
echo "127.0.0.1 mycoolsite.coolz" > /etc/hosts
echo "127.0.0.1 localhost" >> /etc/hosts
echo "::1       localhost" >> /etc/hosts

# Load the new profile
hostz profile copy developer /etc/hosts

# Load the old profile back
hostz host generate default > /etc/hosts
```

## Installation

Install via `homebrew`

```bash
brew install hostz
```

Or install via the `go get` tool

```bash
go get -u github.com/eduardonunesp/hostz
```
