package log

import (
	"log"
	"os"
)

var (
	Debug *log.Logger = log.New(
		os.Stdout,
		"đ DEBUG: ",
		log.Lshortfile,
	)
	Info *log.Logger = log.New(
		os.Stdout,
		"âšī¸ INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Warning *log.Logger = log.New(
		os.Stdout,
		"â ī¸ WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Error *log.Logger = log.New(
		os.Stderr,
		"đĨ ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)
)
