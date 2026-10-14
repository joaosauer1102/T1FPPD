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

func filosofoSequencial(id int) int {
	ciclosConcluidos := 0

	for i := 0; i < numCiclos; i++ {
		fmt.Printf("Filósofo %d está pensando (Ciclo %d)\n", id, i+1)
		time.Sleep(tempoAcao)

		fmt.Printf("Filósofo %d adquiriu os garfos e está comendo (Ciclo %d)\n", id, i+1)
		time.Sleep(tempoAcao)

		fmt.Printf("Filósofo %d liberou os garfos\n", id)
		ciclosConcluidos++
	}

	return ciclosConcluidos
}

func main() {
	inicio := time.Now()
	ciclos := make([]int, numFilosofos)

	for i := 0; i < numFilosofos; i++ {
		ciclos[i] = filosofoSequencial(i)
	}

	tempoTotal := time.Since(inicio)

	fmt.Println("\n===== RESULTADOS V0 - SEQUENCIAL =====")
	fmt.Printf("Tempo total de execução: %v\n", tempoTotal)
	fmt.Println("Tempo médio de espera pelos garfos: 0 ms (sem disputa concorrente)")
	fmt.Println("Maior tempo de espera pelos garfos: 0 ms (sem disputa concorrente)")
	fmt.Println("Máximo de filósofos comendo simultaneamente: 1")

	for i, total := range ciclos {
		fmt.Printf("Filósofo %d: %d ciclos concluídos\n", i, total)
	}
}
