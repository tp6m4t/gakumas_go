package turn

type field interface {
	AddScore(value int)
	SubEnergy(value int)
	SubHealth(value int)
	AddEnergy(value int)
	AddHealth(value int)
}

type Turn struct {
}

func (t *Turn) Start(f *field) {

}
