package lexer

//

type Lexer struct {
	input        string
	position     int  //	current	position	in	input	(points	to	current	char)
	readPosition int  //	current	reading	position	in	input	(the next char)
	ch           byte // current char under examination
}

//init the string input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //reached EOF or nothing has been read yet
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
