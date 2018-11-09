package core

type pool struct {
	pips map[int]*pipe
}

type pipe struct {
}

var PipPool = func() *pool {
	p := &pool{}
	return p
}()

// algorithm for matching the num of pipes to a given task
func (*task) matchPIP() ([]*pipe) {
	return nil
}

func (*pipe) release() {

}

func (*pool) addPip(pip *pipe) {

}
