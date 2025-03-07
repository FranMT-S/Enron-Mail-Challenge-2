package profiling

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

type Profiler struct {
	CpuProf    *os.File
	MemProf    *os.File
	tracerProf *os.File
}

func NewProfiler() *Profiler {
	return &Profiler{
		CpuProf:    nil,
		MemProf:    nil,
		tracerProf: nil,
	}
}

func (prof *Profiler) StartMemAndCPUProfiler() {
	if prof.CpuProf != nil {
		prof.CpuProf.Close()
	}

	if prof.MemProf != nil {
		prof.MemProf.Close()
	}

	cpuProf, err := os.OpenFile("profiling/cpu.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("cannot start  CPU profile")
		log.Println(err)
	}

	prof.CpuProf = cpuProf

	if err := pprof.StartCPUProfile(cpuProf); err != nil {
		log.Println("cannot start  CPU profile")
		log.Println(err)
	}

	memProf, err := os.OpenFile("profiling/mem.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("cannot create memory profile")
		log.Println(err)
	}

	prof.MemProf = memProf

}

func (prof *Profiler) StopMemAndCPUProfiler() {

	pprof.StopCPUProfile()
	if prof.CpuProf != nil {
		prof.CpuProf.Close()
	}

	runtime.GC()
	if prof.MemProf != nil {
		if err := pprof.WriteHeapProfile(prof.MemProf); err != nil {
			log.Println("could not write memory profile")
			log.Println(err)
		}
		prof.MemProf.Close()
	}
}

func (prof *Profiler) StartTracerProfiler() {
	if prof.tracerProf != nil {
		prof.tracerProf.Close()
	}

	traceFile, err := os.OpenFile("profiling/trace.out", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error creating trace profile:", err)
	}

	prof.tracerProf = traceFile

	if err := trace.Start(prof.tracerProf); err != nil {
		fmt.Println("Error starting trace:", err)
	}
}

func (prof *Profiler) StopTracerProfiler() {
	trace.Stop()

	if prof.tracerProf != nil {
		prof.tracerProf.Close()
	}
}
