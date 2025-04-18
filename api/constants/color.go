package constants

type TextColor string

const (
	TextReset   TextColor = "\033[0m"
	TextBlack   TextColor = "\033[30m"
	TextRed     TextColor = "\033[31m"
	TextGreen   TextColor = "\033[32m"
	TextYellow  TextColor = "\033[33m"
	TextBlue    TextColor = "\033[34m"
	TextMagenta TextColor = "\033[35m"
	TextCyan    TextColor = "\033[36m"
	TextWhite   TextColor = "\033[37m"
	TextGray    TextColor = "\033[90m"
)
