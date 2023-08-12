package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

// TestSingleton is a function to test singleton
func TestSingleton(t *testing.T) {
	ins1 := GetSingletonInstance()
	ins2 := GetSingletonInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

// TestParallelSingleton is a function to test singleton in parallel
func TestParallelSingleton(t *testing.T) {
	// Create a channel to synchronize the start of goroutines.
	start := make(chan struct{})

	// Create a WaitGroup to wait for all goroutines to finish.
	wg := sync.WaitGroup{}
	wg.Add(parCount)

	// Create an array to store Singleton instances.
	instances := [parCount]Singleton{}

	// Launch multiple goroutines in a loop.
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// Wait for the start signal from the channel.
			<-start

			// Fetch a Singleton instance and store it in the array.
			instances[index] = GetSingletonInstance()

			// Notify the WaitGroup that the goroutine is done.
			wg.Done()
		}(i)
	}

	// Close the 'start' channel to signal all goroutines to start.
	close(start)

	// Wait for all goroutines to finish.
	wg.Wait()

	// Compare Singleton instances to check if they are equal.
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
