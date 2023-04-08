## Discord Call DDoS Protection Bypass Script

This script allows you to bypass the DDoS protection implemented by Discord for voice calls.

### Installation

Before you can use the script, you'll need to install Go on your system. Here's how you can do it:

```shell
apt install snapd -y
snap install go --classic
go build -o discord main.go
```


### Usage

To use the script, run the following command in your terminal:

```shell
./discord <discord-call-ip> <discord-call-port> <time>
```

### Requirements

To successfully carry out the attack, you'll need a high volume of network traffic. Specifically, you'll need approximately 10 million packets per second (10 Mpps) and 5 gigabits per second (5 Gbps) of bandwidth to bring down a Discord call.

You can use this script on many cloud servers like Google Cloud, Azure or Amazon
