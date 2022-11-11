package gui

type state struct {
	flagIsRunning bool // 当前应用是否正在运行
}

// GenDefault 生成默认状态
func (s state) GenDefault() *state {
	newState := new(state)
	newState.flagIsRunning = false
	return newState
}
