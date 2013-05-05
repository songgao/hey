package main

var (
	TEXT_greeting                  []byte
	TEXT_heuristic_not_implemented []byte
	TEXT_about                     []byte
	TEXT_help                      []byte
	TEXT_message_logged            []byte
	TEXT_message_failed            []byte
	TEXT_message_start             []byte
)

func init() {
	TEXT_greeting = []byte("\n\n\n" + line_to_print("Hey, what's up? I'm "+ROBOT_NAME+", "+USER_NAME+"'s robot.") + line_to_print("What can I do for you? Type /help for help."))
	TEXT_heuristic_not_implemented = []byte(line_to_print("Heuristic not implemented. Falling back to /help."))
	TEXT_about = []byte(line_to_print("Hey! I'm open sourced on https://github.com/songgao/hey"))
	TEXT_help = []byte(
		line_to_print("Available commands:") +
			line_to_print("/help            show this message") +
			line_to_print("/msg [content]   leave "+USER_NAME+" a message") +
			line_to_print("/msgstart        start a multi line message for "+USER_NAME+"; type /msgend to finish it") +
			line_to_print("/about           something about me, the robot"))
	TEXT_message_logged = []byte(line_to_print("Sure. I'll let " + correct_gender("him") + " know."))
	TEXT_message_failed = []byte(line_to_print("Something's wrong; " + correct_gender("he") + "'s probably not gonna get the message. Sorry!"))
	TEXT_message_start = []byte(line_to_print("You've started a multi-line message for " + USER_NAME + ". Leave whatever you want " + correct_gender("him") + " to know and type /msgend to end the message."))
}
