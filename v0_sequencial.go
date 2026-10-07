package main

import (
	"fmt"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	tempoAcao    = 10 * time.Millisecond
)

func filosofoSequencial(id int) {
	for i := 0; i < numCiclos; i++ {
		fmt.Printf("Filósofo %d está pensando (Ciclo %d)\n", id, i+1)
		time.Sleep(tempoAcao) 

		fmt.Printf("Filósofo %d adquiriu os garfos e está comendo (Ciclo %d)\n", id, i+1)
		time.Sleep(tempoAcao) 

		fmt.Printf("Filósofo %d liberou os garfos\n", id)
	}
}

func main() {
	inicio := time.Now()
	
	for i := 0; i < numFilosofos; i++ {
		filosofoSequencial(i)
	}
	
	fmt.Printf("\nTempo total (V0 - Sequencial): %v\n", time.Since(inicio))
}