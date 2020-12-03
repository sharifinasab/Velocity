package service

import (
	"log"

	"KOHO/model"
	"KOHO/repository"
)

// Manager holds structure to process transaction
type Manager struct {
	storeRepository repository.IStoreRepository
	receiveChannel  <-chan *model.Deposit
	sendChannel     chan<- string
}

func (m *Manager) executeDeposite(d *model.Deposit) (bool, error) {
	success := false
	log.Println("Manager: received a deposit to execute ...", d.DepositRequest.ID)

	if m.storeRepository.IsDuplicatedTransaction(d.DepositRequest.ID, d.DepositRequest.CustomerID) {
		log.Println("Manager: Duplicated deposit ", d.DepositRequest.ID, " ignored.")
		return success, nil
	}

	acc := m.storeRepository.GetAccount(d.DepositRequest.CustomerID)
	// check if rules still effective
	acc.SetAccountQuota(d.ParsedTime)

	acc.PrintRemainingDepositLimits()

	if acc.ValidateDeposit(d) {
		acc.Deposit(d)
		m.storeRepository.AddTransaction(d.DepositRequest.ID, d.DepositRequest.CustomerID)
		success = true
	}

	log.Println("Manager: deposit process finished. Sending result ...")

	// changes to account can be persisted here
	// {
	// }

	return success, nil
}

func (m *Manager) sendResponse(d *model.Deposit, flag bool) {
	response :=
		FormatDepositResponse(d.DepositRequest.ID, d.DepositRequest.CustomerID, flag)

	// send result to another process to write it in output - channel
	m.sendChannel <- response
}

// StartManager starts transaction manager
func StartManager(reqChannel <-chan *model.Deposit, resChannel chan<- string) error {
	m, err := initializeManager(reqChannel, resChannel)

	if err != nil {
		return err
	}

	m.process()
	m.cleanup()

	return nil
}

func initializeManager(reqChannel <-chan *model.Deposit, resChannel chan<- string) (*Manager, error) {
	log.Println("Initializing manager ...")
	return &Manager{
		storeRepository: repository.NewStore(),
		receiveChannel:  reqChannel,
		sendChannel:     resChannel,
	}, nil
}

func (m *Manager) process() {
	log.Println("Manager-Process started")
	for deposit := range m.receiveChannel {
		result, err := m.executeDeposite(deposit)

		if err != nil {
			continue
		}
		m.sendResponse(deposit, result)
	}
}

func (m *Manager) cleanup() {
	close(m.sendChannel)
}
