package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)  // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
	s.Stop()

	s.UpdateCharSet(spinner.CharSets[7])  // Update spinner to use a different character set
	s.Restart()                           // Restart the spinner
	time.Sleep(4 * time.Second)
	s.Stop()

	s.UpdateSpeed(200 * time.Millisecond) // Update the speed the spinner spins at
	s.Restart()
	time.Sleep(4 * time.Second)
	s.Stop()

	/*
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
    */
}
