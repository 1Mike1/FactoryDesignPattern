package main

type Mobile struct {
  name string
  brand string
  ram string
  price string
}

func (m *Mobile) setName(name string){
  m.name = name
}

func (m *Mobile) setBrand(brand string){
  m.brand = brand
}

func (m *Mobile) setRAM(ram string){
  m.ram = ram
}

func (m *Mobile) setPrice(price string){
  m.price = price
}

func (m *Mobile) getName() string {
  return m.name;
}
func (m *Mobile) getBrand() string {
  return m.brand;
}

func (m *Mobile) getRAM() string {
  return m.ram;
}

func (m *Mobile) getPrice() string {
  return m.price;
}


