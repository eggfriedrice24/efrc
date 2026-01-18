package main

import (
	"fmt"
	"log"

	"efrc/internal/client"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current configuration and status",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := client.LoadConfig()
		if err != nil {
			log.Fatalf("No config found. Run 'efrc init' first: %v", err)
		}

		fmt.Println("=== EFRC Status ===")
		fmt.Printf("Device:     %s\n", config.DeviceName)
		fmt.Printf("Public Key: %s\n", config.PublicKey)
		fmt.Printf("Virtual IP: %s\n", valueOrNone(config.VirtualIP))
		fmt.Printf("Server:     %s\n", valueOrNone(config.ServerURL))
		fmt.Printf("Registered: %v\n", config.DeviceID != "")
	},
}

func valueOrNone(s string) string {
	if s == "" {
		return "(not set)"
	}
	return s
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
