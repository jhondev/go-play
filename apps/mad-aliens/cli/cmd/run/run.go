package run

import (
	"fmt"
	"mad-aliens/pkg/data/providers"
	"mad-aliens/pkg/data/records"
	"mad-aliens/pkg/simulation"
	"mad-aliens/pkg/world"
	"math/rand"

	"github.com/spf13/cobra"
)

const (
	flagWorldFilePath = "path"
	flagAliensNumber  = "naliens"
	flagMaxMoves      = "maxmoves"
	defaultAliens     = 5
	defaultMaxMovs    = 10000
)

func New() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the aliens invasion simulation",
		Long:  ``,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			path, _ := flags.GetString(flagWorldFilePath)
			naliens, _ := flags.GetInt(flagAliensNumber)
			maxmoves, _ := flags.GetInt(flagMaxMoves)
			err := runSimulation(path, naliens, maxmoves)
			if err != nil {
				fmt.Printf("error running the simulation: %v", err)
			}
		},
	}
	runCmd.Flags().StringP(flagWorldFilePath, "p", "", "world map file path")
	_ = runCmd.MarkFlagRequired(flagWorldFilePath)
	runCmd.Flags().IntP(flagAliensNumber, "n", defaultAliens, "number of aliens in the simulation")
	runCmd.Flags().IntP(flagMaxMoves, "m", defaultMaxMovs, "max number of moves")

	return runCmd
}

func runSimulation(path string, naliens int, maxmoves int) error {

	fmt.Printf("Starting alien invasion simulation with the following setup:\n\n"+
		"  number of aliens: %d\n"+
		"  max moves:        %d\n\n", naliens, maxmoves)

	fmt.Println("👽 Loading world.")

	// use the file provider.
	// we can create different provider types
	// to grap the info from other sources like apis or databases
	prov := providers.NewFile(path)
	// load the world with cities and random aliens positions
	wld, err := world.Load(prov, naliens, maxmoves, rand.Intn)
	if err != nil {
		return err
	}

	// use memory records store.
	// we can create different records store types
	// to manage the history information of a simulation
	// and recrete it based on the events and states
	rec, err := records.NewMem(wld)
	if err != nil {
		return err
	}

	// create a new simulation object injecting the loaded values
	sim := simulation.New(wld, rec, rand.Intn)

	fmt.Printf("👽 Mad aliens are invading:\n\n")

	err = sim.Run()
	if err != nil {
		return err
	}

	fmt.Printf("👽 The mad aliens invasion has finished with the following report:\n\n")
	sim.PrintReport()

	return nil
}
