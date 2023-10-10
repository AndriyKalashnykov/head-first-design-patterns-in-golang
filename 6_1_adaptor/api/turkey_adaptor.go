package api

/**
 * We need to get a reference to the
 * object we are adapting.
 */
type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdaptor(t Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{
		turkey: t,
	}
}

/**
 * Implement the interface of the
 * type you’re adapting to. This is the
 * interface your client expects to see
 */
func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 1; i <= 5; i++ {
		t.turkey.Fly()
	}
}
