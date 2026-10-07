package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	tempoAcao    = 10 * time.Millisecond
)

type Garfo struct {
	sync.Mutex
}

type Filosofo struct {
	id       int
	garfoEsq *Garfo
	garfoDir *Garfo
}

func (f Filosofo) jantar(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numCiclos; i++ {
		fmt.Printf("Filósofo %d pensando (Ciclo %d)\n", f.id, i+1)
		time.Sleep(tempoAcao)

		f.garfoEsq.Lock()
		fmt.Printf("Filósofo %d pegou o garfo esquerdo\n", f.id)

		time.Sleep(1 * time.Millisecond)

		f.garfoDir.Lock()
		fmt.Printf("Filósofo %d pegou o garfo direito e está comendo (Ciclo %d)\n", f.id, i+1)
		time.Sleep(tempoAcao)

		f.garfoDir.Unlock()
		f.garfoEsq.Unlock()
		fmt.Printf("Filósofo %d terminou de comer e liberou os garfos\n", f.id)
	}
}

func main() {
	garfos := make([]*Garfo, numFilosofos)
	for i := 0; i < numFilosofos; i++ {
		garfos[i] = &Garfo{}
	}

	filosofos := make([]*Filosofo, numFilosofos)
	for i := 0; i < numFilosofos; i++ {
		filosofos[i] = &Filosofo{
			id:       i,
			garfoEsq: garfos[i],
			garfoDir: garfos[(i+1)%numFilosofos],
		}
	}

	var wg sync.WaitGroup
	inicio := time.Now()

	for i := 0; i < numFilosofos; i++ {
		wg.Add(1)
		go filosofos[i].jantar(&wg)
	}

	wg.Wait()
	fmt.Printf("\nTempo total (V1): %v\n", time.Since(inicio))
}