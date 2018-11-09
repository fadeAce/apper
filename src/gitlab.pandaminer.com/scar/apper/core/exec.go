package core

import "gitlab.pandaminer.com/scar/apper/client"

// this package unit is for executors to exec pipes
// then assemble all data and make a transaction to store them

// important interface for starting s pool
func StartPool(sum int, notifier *client.Notifier) {
	// create pipes step.1
	//

}

func (*task) run() map[*pipe]*DataUnit {
	return nil
}

// when task is attempt to store a assemble
// it will fetch them from dataAssemble by pipes it's currently using
// then release all pipes after store finished
func (*task) store() {

}
