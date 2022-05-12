package facade

type engine struct {
	maxRpm int
	rpm    int
}

//NewEngine создаём новый инстанс двигателя
func (e *engine) NewEngine(maxRpm int) *engine {
	return &engine{
		maxRpm: maxRpm,
		rpm:    0,
	}
}

//SetRpm установка нового количества оборотов двигателя
func (e *engine) SetRpm(rpm int) {
	if rpm > e.maxRpm {
		e.rpm = e.maxRpm
	} else {
		e.rpm = rpm
	}
}
