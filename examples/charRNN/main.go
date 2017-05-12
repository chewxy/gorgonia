package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/chewxy/cu"
	T "github.com/chewxy/gorgonia"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var blockprofile = flag.String("blockprofile", "", "write blocking profile to this file")

var bpFile io.Writer

// prediction params
var softmaxTemperature = 1.0
var maxCharGen = 100

// various global variable inits
var epochSize = -1
var inputSize = -1
var outputSize = -1

// gradient update stuff
var l2reg = 0.000001
var learnrate = 0.01
var clipVal = 5.0

type contextualError interface {
	error
	Node() *T.Node
	Value() T.Value
	InstructionID() int
}

func cleanup(sigChan chan os.Signal, doneChan chan bool, profiling bool, bp *pprof.Profile) {
	select {
	case <-sigChan:
		log.Println("EMERGENCY EXIT!")
		if profiling {
			pprof.StopCPUProfile()
		}
		if bp != nil {
			bp.WriteTo(bpFile, 0)
		}
		os.Exit(1)

	case <-doneChan:
		return
	}
}

func main() {
	flag.Parse()
	rand.Seed(1337)
	// intercept Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// doneChan := make(chan bool, 1)
	// defer func() {
	// 	nn, cc, ec := T.GraphCollisionStats()
	// 	log.Printf("COLLISION COUNT: %d/%d. Expected : %d", cc, nn, ec)
	// }()

	// var profiling bool
	var blockprof *pprof.Profile
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// profiling = true
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *blockprofile != "" {
		f, err := os.Create(*blockprofile)
		if err != nil {
			log.Fatal(err)
		}
		bpFile = f

		blockprof = pprof.Lookup("block")
		runtime.SetBlockProfileRate(1)

		defer func() {
			blockprof.WriteTo(f, 0)
		}()
	}
	// go cleanup(sigChan, doneChan, profiling, blockprof)

	m := NewLSTMModel(inputSize, embeddingSize, outputSize, hiddenSizes)

	solver := T.NewRMSPropSolver(T.WithLearnRate(learnrate), T.WithL2Reg(l2reg), T.WithClip(clipVal))
	start := time.Now()
	eStart := start
	for i := 0; i <= 100; i++ {
		// log.Printf("Iter: %d", i)
		// _, _, err := m.run(i, solver)
		cost, perp, err := m.run(i, solver)
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}

		if i%100 == 0 {
			timetaken := time.Since(eStart)
			fmt.Printf("Time Taken: %v\tCost: %v\tPerplexity: %v\n", timetaken, cost, perp)
			eStart = time.Now()
		}

		if i%1000 == 0 {
			log.Printf("Going to predict now")
			m.predict()
			log.Printf("Done predicting")
		}

		if *memprofile != "" && i == 1000 {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
			return
		}

	}

	end := time.Now()
	fmt.Printf("%v", end.Sub(start))
	fmt.Printf("%+3.3s", m.embedding.Value())
	log.Printf("AverageQueueLength %v", cu.AverageQueueLength())
	log.Printf("Blocking Callers %v", cu.BlockingCallers())
	ql := cu.QueueLengths()
	var sum int
	for _, q := range ql {
		sum += q
	}
	log.Printf("Calls %d", sum)
	log.Printf("Queue Lengths %v", ql)
}
