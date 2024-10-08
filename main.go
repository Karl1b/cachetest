package main

import (
	"log"
	"strconv"

	cachetest "github.com/karl1b/cachetest/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cachetest",
	Short: "cachetest",
	Long:  `cachetest`,
}

var customCacheCmd = &cobra.Command{
	Use:   "custom",
	Short: "Run CustomCache test",
	Run: func(cmd *cobra.Command, args []string) {

		cachetest.RunCustomCache(convertInput(args))
	},
}

var customFastCacheCmd = &cobra.Command{
	Use:   "customfast",
	Short: "Run CustomCache test",
	Run: func(cmd *cobra.Command, args []string) {

		cachetest.RunCustomFastCache(convertInput(args))
	},
}

var goCacheCmd = &cobra.Command{
	Use:   "gocache",
	Short: "Run Go-Cache test",
	Run: func(cmd *cobra.Command, args []string) {

		cachetest.RunGoCache(convertInput(args))
	},
}

var freeCacheCmd = &cobra.Command{
	Use:   "freecache",
	Short: "Run freecache test",
	Run: func(cmd *cobra.Command, args []string) {

		cachetest.RunFreeCache(convertInput(args))
	},
}

var ristrettoCacheCmd = &cobra.Command{
	Use:   "ristretto",
	Short: "Run ristretto test",
	Run: func(cmd *cobra.Command, args []string) {

		cachetest.RunRistretto(convertInput(args))
	},
}

func init() {
	rootCmd.AddCommand(customCacheCmd)
	rootCmd.AddCommand(customFastCacheCmd)
	rootCmd.AddCommand(goCacheCmd)
	rootCmd.AddCommand(freeCacheCmd)
	rootCmd.AddCommand(ristrettoCacheCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

func convertInput(args []string) (int, int, int) {

	numberOfKeys, err := strconv.Atoi(args[0])
	if err != nil {

		log.Fatalln(err)
	}

	opsPerWorker, err := strconv.Atoi(args[1])
	if err != nil {

		log.Fatalln(err)
	}

	allowedCacheSizeMB, err := strconv.Atoi(args[2])
	if err != nil {

		log.Fatalln(err)
	}

	return numberOfKeys, opsPerWorker, allowedCacheSizeMB

}
