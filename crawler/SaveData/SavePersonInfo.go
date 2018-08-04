package SaveData
type Person struct{
	items chan interface{}
	saveQueue chan chan interface{}
}

func (p *Person) CreateSaver (item interface{}) {
	p.items <- item
}
func (p * Person) Run() {

}
