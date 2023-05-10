package main

type Samsung struct {
  Mobile
}

func Fold() Device {
  return &Samsung {
    Mobile : Mobile{
      name : "Galaxy Fold 3",
      brand : "Samsung",
      ram : "12 GB",
      price : "1,50,00 Rs",
    },
  }
}
