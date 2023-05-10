package main

type Device interface {
  setName(name string);
  setBrand(brand string);
  setRAM(ram string);
  setPrice(price string);
  getName() string;
  getBrand() string;
  getRAM() string;
  getPrice() string;
}
