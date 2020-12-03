package service

import (
	"bufio"
	"log"
	"os"

	"KOHO/model"
	"KOHO/util"
)

// Input holds transaction sender structure
type Input struct {
	file        *os.File
	sendChannel chan<- *model.Deposit
}

// GenerateDeposits simulates deposit calls
func (i *Input) GenerateDeposits(inputFile *os.File) error {
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		req, err := ParseDepositRequest(scanner.Text())
		if err != nil {
			continue
		}
		// send to process thread - channel
		i.sendChannel <- req

		log.Println("Deposit request ", req.DepositRequest.ID, " sent successfuly.")
	}

	return nil
}

// StartCaller starts transaction sender process
func StartCaller(inputPath string, channel chan<- *model.Deposit) error {
	file := util.ReadFile(inputPath)
	i, err := initializeCaller(file, channel)

	if err != nil {
		return err
	}

	i.GenerateDeposits(file)
	i.cleanup(file)

	return nil
}

func initializeCaller(path *os.File, channel chan<- *model.Deposit) (*Input, error) {
	log.Println("Initializing caller ...")
	return &Input{
		file:        path,
		sendChannel: channel,
	}, nil
}

func (i *Input) cleanup(file *os.File) {
	file.Close()
	close(i.sendChannel)
}
