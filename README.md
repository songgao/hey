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
\* replace 127.0.0.1 with your actual IP address running the robot

## Develop
`fsm.go` is where implements the finite state machine for processing messages. Implement more transitions to enrich functionalities of `hey`!
