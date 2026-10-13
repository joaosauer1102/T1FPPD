package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	numExecucoes = 5
	tempoAcao    = 10 * time.Millisecond
)

type Estatisticas struct {
	mu                sync.Mutex
	esperaTotal       time.Duration
	maiorEspera       time.Duration
	aquisicoes        int
	comendoAgora      int
	maxComendo        int
	ciclosPorFilosofo []int
}

func (e *Estatisticas) registrarEspera(d time.Duration) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.esperaTotal += d
	e.aquisicoes++
	if d > e.maiorEspera {
		e.maiorEspera = d
	}
}

func (e *Estatisticas) comecouComer() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.comendoAgora++
	if e.comendoAgora > e.maxComendo {
		e.maxComendo = e.comendoAgora
	}
}

func (e *Estatisticas) terminouComer(id int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.comendoAgora--
	e.ciclosPorFilosofo[id]++
}

func executarVersao3() {
	var garfos [numFilosofos]sync.Mutex
	var wg sync.WaitGroup
	stats := Estatisticas{ciclosPorFilosofo: make([]int, numFilosofos)}

	inicio := time.Now()
	for id := 0; id < numFilosofos; id++ {
		goFilosofo := id
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			esq := id
			dir := (id + 1) % numFilosofos

			// Ordena os garfos pelo índice global para quebrar a espera circular.
			primeiro, segundo := esq, dir
			if primeiro > segundo {
				primeiro, segundo = segundo, primeiro
			}

			for ciclo := 0; ciclo < numCiclos; ciclo++ {
				time.Sleep(tempoAcao) // pensar

				inicioEspera := time.Now()
				garfos[primeiro].Lock()
				garfos[segundo].Lock()
				stats.registrarEspera(time.Since(inicioEspera))

				stats.comecouComer()
				time.Sleep(tempoAcao) // comer
				stats.terminouComer(id)

				garfos[segundo].Unlock()
				garfos[primeiro].Unlock()
			}
		}(goFilosofo)
	}
	wg.Wait()
	duracao := time.Since(inicio)

	fmt.Printf("V3 - ordenação na aquisição dos garfos | Execução concluída\n")
	fmt.Printf("Tempo total: %v\n", duracao)
	fmt.Printf("Espera média por aquisição dos dois garfos: %v\n", stats.esperaTotal/time.Duration(stats.aquisicoes))
	fmt.Printf("Maior espera por aquisição: %v\n", stats.maiorEspera)
	fmt.Printf("Máximo de filósofos comendo simultaneamente: %d\n", stats.maxComendo)
	fmt.Printf("Ciclos concluídos por filósofo: %v\n", stats.ciclosPorFilosofo)
}

func main() {
	for execucao := 1; execucao <= numExecucoes; execucao++ {
		fmt.Printf("\n========== Execução %d/%d ==========\n", execucao, numExecucoes)
		executarVersao3()
	}
}
