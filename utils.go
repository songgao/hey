package main

func line_to_print(line string) []byte {
	return []byte("\x1b[38;05;4m--- \x1b[0m" + line + "\n")
}

func blank_line() []byte {
	return []byte("\n")
}

func correct_gender(word string) string {
	switch word {
	case "him":
		if USER_GENDER == MALE {
			return "him"
		} else {
			return "her"
		}
	case "his":
		if USER_GENDER == MALE {
			return "his"
		} else {
			return "her"
		}
	case "he":
		if USER_GENDER == MALE {
			return "he"
		} else {
			return "she"
		}
	}
	return ""
}
