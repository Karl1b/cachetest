package cachetest

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	freecache "github.com/coocood/freecache"
	ristretto "github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	go_cache "github.com/patrickmn/go-cache"
	"golang.org/x/exp/rand"
)

const (
	smallHTTPSize  = 1024       // 1 KB
	largeHTTPSize  = 1024 * 3   // 3 KB
	smallImageSize = 250 * 1024 // 250 KB

)

var (
	// The different caches
	goCache           *go_cache.Cache  // go_cache "github.com/patrickmn/go-cache"
	customSimpleCache *CustomCache     // The custom cache I developed.
	customFastCache   *CustomFastCache // The custom cache I developed, but without so much info. and with Rlock instead.
	fcache            *freecache.Cache // "github.com/coocood/freecache"
	ristrettoCache    *ristretto.Cache[string, []byte]
	sizeSmall         []byte // Just random junk to mimic value size
	sizeMedium        []byte
	sizeLarge         []byte

	uuidSlice []string // Holds the keys I use. They are uuids

	startTime     time.Time
	startloadtime time.Time
	loadduration  time.Duration
	duration      time.Duration

	numWorkers         int // Number of test workers
	opsPerWorker       int // Operation of workers
	numberOfKeys       int // Should vary from 10k to 1M
	allowedCacheSizeMB int // Should vary. Maybe 10 mB to 200 mB?
)

func init() {
	sizeSmall = generateRandomContent(smallHTTPSize)
	sizeMedium = generateRandomContent(largeHTTPSize)
	sizeLarge = generateRandomContent(smallImageSize)
	rand.Seed(uint64(time.Now().UnixMilli()))
	runtime.GOMAXPROCS(runtime.NumCPU())
	numWorkers = runtime.NumCPU() * 4
}

func RunRistretto(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI

	uuidSlice = generateKeySlice(numberOfKeys)

	start("ristretto", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	var err error
	ristrettoCache, err = ristretto.NewCache[string, []byte](&ristretto.Config[string, []byte]{
		NumCounters: int64(numberOfKeys), // number of keys
		MaxCost:     int64(allowedCacheSizeMB * 1024 * 1024),
		BufferItems: 64,
	})
	if err != nil {
		log.Fatal(err)
	}
	startloadtime = time.Now()

	testCache(testRistrettoCache, 0)
	loadduration = time.Since(startloadtime)
	startTime = time.Now()
	testCache(testRistrettoCache, 1)
	duration = time.Since(startTime)

	fmt.Printf("Ristretto Cache test completed in %v\n", duration)
	end("ristretto_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)

}

func RunCustomFastCache(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI

	uuidSlice = generateKeySlice(numberOfKeys)
	start("custom_fast_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	fmt.Printf("customCache_nkeys:%v_opsPw:%v_cacheSize%v.prof", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	customFastCache = NewCustomFastCache(allowedCacheSizeMB)
	fmt.Println("custom cache created")

	startloadtime = time.Now()
	testCache(testcustomFastCache, 0)
	loadduration = time.Since(startloadtime)
	startTime = time.Now()
	testCache(testcustomFastCache, 1)
	duration = time.Since(startTime)
	fmt.Printf("CustomCache test completed in %v\n", duration)

	end("custom_fast_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)
}

func RunCustomCache(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI

	uuidSlice = generateKeySlice(numberOfKeys)
	start("custom_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB)

	fmt.Println("######################")
	fmt.Printf("customCache_nkeys:%v_opsPw:%v_cacheSize%v.prof", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	customSimpleCache = NewCustomSimpleCache(allowedCacheSizeMB)
	fmt.Println("custom cache created")
	startloadtime = time.Now()
	testCache(testcustomSimpleCache, 0)
	loadduration = time.Since(startloadtime)
	startTime = time.Now()
	testCache(testcustomSimpleCache, 1)
	duration = time.Since(startTime)
	fmt.Printf("CustomCache test completed in %v\n", duration)
	fmt.Printf("Total size: %d MB\n", customSimpleCache.totalsizeB/1024/1024)
	fmt.Printf("Hits: %d\n", customSimpleCache.hits)
	fmt.Printf("Misses: %d\n", customSimpleCache.misses)
	end("custom_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)
}

func RunGoCache(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI
	uuidSlice = generateKeySlice(numberOfKeys)
	fmt.Println("######################")
	fmt.Printf("gocache_nkeys:%v_opsPw:%v_cacheSize%v.prof", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	start("go_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	goCache = go_cache.New(go_cache.NoExpiration, 0)
	startloadtime = time.Now()
	testCache(testGoCache, 0)
	loadduration = time.Since(startloadtime)
	startTime = time.Now()
	testCache(testGoCache, 1)
	duration = time.Since(startTime)
	fmt.Printf("GoCache test completed in %v\n", duration)
	end("go_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)
}

func RunFreeCache(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI
	uuidSlice = generateKeySlice(numberOfKeys)
	start("free_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB)
	fcache = freecache.NewCache(allowedCacheSizeMB)
	startloadtime = time.Now()
	testCache(testCache_freeCache, 0)
	loadduration = time.Since(startloadtime)
	startTime = time.Now()
	testCache(testCache_freeCache, 1)
	duration = time.Since(startTime)
	fmt.Printf("FreeCache test completed in %v\n", duration)
	end("free_cache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)
}

func testCache_freeCache(uuidSlice *[]string, mode int) {
	if mode == 0 {

		for i := 0; i < numberOfKeys; i++ {
			key := (*uuidSlice)[i]

			content := getRandomContent()
			fcache.Set([]byte(key), content, 0)
		}

	} else {

		for i := 0; i < opsPerWorker; i++ {
			key := (*uuidSlice)[rand.Intn(numberOfKeys)]
			switch rand.Intn(10) {
			case 0:
				content := getRandomContent()
				fcache.Set([]byte(key), content, 0)
			case 1:
				fcache.Del([]byte(key))
			default:
				fcache.Get([]byte(key))
			}

			mode := rand.Intn(2)

			if mode == 0 {
				content := getRandomContent()
				fcache.Set([]byte(key), content, 0)
			} else {
				fcache.Get([]byte(key))
			}
		}
	}
}

func testGoCache(uuidSlice *[]string, mode int) {

	if mode == 0 {

		for i := 0; i < numberOfKeys; i++ {
			key := (*uuidSlice)[i]
			content := getRandomContent()
			goCache.Set(key, content, go_cache.NoExpiration)

		}

	} else {

		for i := 0; i < opsPerWorker; i++ {
			key := (*uuidSlice)[rand.Intn(numberOfKeys)]

			switch rand.Intn(10) {
			case 0:
				content := getRandomContent()
				goCache.Set(key, content, go_cache.NoExpiration)
			case 1:
				goCache.Delete(key)

			default:
				goCache.Get(key)
			}
		}
	}
}

func testcustomSimpleCache(uuidSlice *[]string, mode int) {

	if mode == 0 {
		for i := 0; i < numberOfKeys; i++ {
			key := (*uuidSlice)[i]
			content := getRandomContent()
			customSimpleCache.Set(key, content)

		}
	} else {

		for i := 0; i < opsPerWorker; i++ {
			key := (*uuidSlice)[rand.Intn(numberOfKeys)]

			switch rand.Intn(10) {
			case 0:
				content := getRandomContent()
				customSimpleCache.Set(key, content)
			case 1:
				customSimpleCache.Del(key)
			default:
				customSimpleCache.Get(key)
			}
		}
	}
}

func testcustomFastCache(uuidSlice *[]string, mode int) {

	if mode == 0 {

		for i := 0; i < numberOfKeys; i++ {
			key := (*uuidSlice)[i]

			content := getRandomContent()
			customFastCache.Set(key, content)

		}

	} else {

		for i := 0; i < opsPerWorker; i++ {
			key := (*uuidSlice)[rand.Intn(numberOfKeys)]

			switch rand.Intn(10) {
			case 0:
				content := getRandomContent()
				customFastCache.Set(key, content)
			case 1:
				customFastCache.Del(key)
			default:
				customFastCache.Get(key)
			}
		}
	}
}

func testRistrettoCache(uuidSlice *[]string, mode int) {

	if mode == 0 {
		// I know this happends concurrently, but lets test :)
		for i := 0; i < numberOfKeys; i++ {
			key := (*uuidSlice)[i]

			content := getRandomContent()
			ristrettoCache.Set(key, content, int64(len(content)))

		}
	} else {

		for i := 0; i < opsPerWorker; i++ {
			key := (*uuidSlice)[rand.Intn(numberOfKeys)]

			switch rand.Intn(10) {
			case 0:
				content := getRandomContent()
				ristrettoCache.Set(key, content, int64(len(content)))
			case 1:
				ristrettoCache.Del(key)
			default:
				_, _ = ristrettoCache.Get(key)
			}
		}

	}

}

func getRandomContent() []byte {
	switch rand.Intn(3) {
	case 0:
		return sizeSmall
	case 1:
		return sizeMedium
	default:
		return sizeLarge
	}
}

func start(name string, numberOfKeys int, opsPerWorker int, allowedCacheSizeB int) {
	cpufilename := fmt.Sprintf("cpu_%s_nkeys:%v_opsPw:%v_cacheSize%v.prof", name, numberOfKeys, opsPerWorker, allowedCacheSizeB)

	cpuFile, err := os.Create(cpufilename)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuFile)
}

func generateKeySlice(n int) []string {
	uuids := make([]string, n)
	for i := 0; i < n; i++ {
		uuids[i] = uuid.New().String()
	}
	return uuids
}

func generateRandomContent(size int) []byte {
	content := make([]byte, size)
	rand.Read(content)
	return content
}

// Utils
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Testwrapper.
func testCache(worker func(*[]string, int), mode int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker(&uuidSlice, mode)
		}()
	}

	wg.Wait()
}

func end(name string, numberOfKeys int, opsPerWorker int, allowedCacheSizeB int, duration time.Duration) {
	memFilename := fmt.Sprintf("mem_%s_nkeys:%v_opsPw:%v_cacheSize%v.prof", name, numberOfKeys, opsPerWorker, allowedCacheSizeB)
	pprof.StopCPUProfile()
	memFile, err := os.Create(memFilename)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer memFile.Close()
	runtime.GC() // Run a garbage collection
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Get CPU information
	numCPU := runtime.NumCPU()
	numGoroutine := runtime.NumGoroutine()

	// Append results to CSV file
	csvFile, err := os.OpenFile("benchmark_results2.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Check if the file is empty and write header if needed
	fileInfo, err := csvFile.Stat()
	if err != nil {
		log.Fatal("Failed to get file info:", err)
	}
	if fileInfo.Size() == 0 {
		header := []string{
			"Cache", "Number of Keys", "Ops Per Worker", "Cache Size (MB)", "LoadDuration (s)",
			"Duration (s)", "Alloc (MB)", "TotalAlloc (MB)", "Sys (MB)", "NumGC",
			"NumCPU", "NumGoroutine", "GOMAXPROCS",
		}
		if err := writer.Write(header); err != nil {
			log.Fatal("Failed to write CSV header:", err)
		}
	}

	// Write benchmark results
	record := []string{
		name,
		fmt.Sprintf("%d", numberOfKeys),
		fmt.Sprintf("%d", opsPerWorker),
		fmt.Sprintf("%d", allowedCacheSizeB),
		fmt.Sprintf("%.2f", loadduration.Seconds()),
		fmt.Sprintf("%.2f", duration.Seconds()),
		fmt.Sprintf("%d", bToMb(m.Alloc)),
		fmt.Sprintf("%d", bToMb(m.TotalAlloc)),
		fmt.Sprintf("%d", bToMb(m.Sys)),
		fmt.Sprintf("%d", m.NumGC),
		fmt.Sprintf("%d", numCPU),
		fmt.Sprintf("%d", numGoroutine),
		fmt.Sprintf("%d", runtime.GOMAXPROCS(0)),
	}

	if err := writer.Write(record); err != nil {
		log.Fatal("Failed to write CSV record:", err)
	}

	fmt.Printf("Benchmark results appended to benchmark_results.csv\n")
	printMemStats()
	printCPUStats(numCPU, numGoroutine)
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func printCPUStats(numCPU, numGoroutine int) {
	fmt.Printf("NumCPU = %d", numCPU)
	fmt.Printf("\tNumGoroutine = %d", numGoroutine)
	fmt.Printf("\tGOMAXPROCS = %d\n", runtime.GOMAXPROCS(0))
}
