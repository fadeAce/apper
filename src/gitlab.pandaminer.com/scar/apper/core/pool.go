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
// what returned are threads bond to selected task
// if no pipe is spared this invoke would be blocked
func (*task) MatchPIP() (map[int]*pipe) {
	return nil
}

func (*task) release() {

}

// task done involved a process that inject result from
// fragments to cache layer.
func (t *task) Done() string {
	t.release()
	return t.txID
}

func (*pipe) release() {

}

func (*pool) addPip(pip *pipe) {

}
