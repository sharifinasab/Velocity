package service

import (
	"bufio"
	"log"
	"os"

	"KOHO/util"
)

// Output holds transactions results structure
type Output struct {
	file           *os.File
	writer         *bufio.Writer
	receiveChannel <-chan string
}

// WriteResponse writes the deposit response to file
func (o *Output) writeResponse(response string) error {
	// write the response
	if _, err := o.writer.WriteString(response + "\n"); err != nil {
		return err
	}

	return nil
}

// GetResponse get the response messages from transaction manager
// and writes it to output
func (o *Output) getResponse() {
	for response := range o.receiveChannel {
		log.Println("Producer: got a response to write")
		if err := o.writeResponse(response); err != nil {
			continue
		}
	}
}

// Cleanup frees up the resourses
func (o *Output) cleanup() {
	o.writer.Flush()
	o.file.Close()
}

// StartProducer starts the process
func StartProducer(ch <-chan string) error {
	file := util.CreateOutputFile()
	output, err := initializeProducer(file, ch)

	if err != nil {
		return err
	}

	// channel
	output.getResponse()
	output.cleanup()

	return nil
}

// initialize the output to write into
func initializeProducer(path *os.File, ch <-chan string) (*Output, error) {
	log.Println("Initializing producer ...")
	writer := bufio.NewWriter(path)

	return &Output{
		file:           path,
		writer:         writer,
		receiveChannel: ch,
	}, nil
}
