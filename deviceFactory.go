package main

import "fmt"

func getMobile(mobileType string) (Device, error)  {
  if mobileType=="samsung"{
    return Fold(), nil
  }
  if mobileType=="apple"{
    return iPhone(), nil
  }

  return nil, fmt.Errorf("Invalid device")
}
