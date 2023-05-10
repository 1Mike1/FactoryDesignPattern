package main

type Apple struct {
  Mobile
}

func iPhone() Device {
  return &Apple{
    Mobile : Mobile{
      name : "iPhone 14 Pro Max",
      brand : "Apple",
      ram : "6 GB",
      price : "1,50,00 Rs",
    },
  }
}
