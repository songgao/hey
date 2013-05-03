# hey
`hey` is a robot implemented in Go with finite state machine that has potential to work as an automated secretary for you. It's currently not quite useful except letting others leave messages for you.

## Installation
```bash
git clone git://github.com/songgao/hey.git
```

## Usage
Edit config.go to meet your needs.

Build the robot
```bash
cd hey
go build
```

Bring up the robot
```bash
./hey
```

Ask your friend to contact your robot through `openssl` or `telnet`:
```bash
openssl s_client -no_ssl2 -connect 127.0.0.1:56789
```
```bash
telnet 127.0.0.1:56788
```
```bash
bash-4.2$ telnet localhost 56788
Trying ::1...
Connected to localhost.
Escape character is '^]'.



--- Hey, what's up? I'm Oag Gnos, Song Gao's robot.
--- What can I do for you? Type /help for help.
/help
--- Available commands:
--- /help              show this message
--- /msg [content]     leave Song Gao a message
/msg Hey
--- Sure. I'll let him know.
/msg How's going?
--- Sure. I'll let him know.
^]
telnet> Connection closed.
bash-4.2$ cat ~/.hey/logs/2013-05-03T15\:23\:17-05\:00-\[\:\:1\]\:63987
Message: Hey

Message: How's going?

```
\* replace 127.0.0.1 with your actual IP address running the robot

## Develop
`fsm.go` is where implements the finite state machine for processing messages. Implement more transitions to enrich functionalities of `hey`!
