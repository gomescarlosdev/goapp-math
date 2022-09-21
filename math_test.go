package main

import "testing"

func TestSum(t *testing.T){
    total := Sum(15, 15)

    if total != 30 {
	    t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 30)
    }
}

func TestMulti(t *testing.T){
    total := Sum(10, 15)

    if total != 150 {
	    t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 150)
    }
}