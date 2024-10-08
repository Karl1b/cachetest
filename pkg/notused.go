package cachetest

// Stuff here is not used, because it was too slow

/* func RunEmarCache(numberOfKeysI int, opsPerWorkerI int, allowedCacheSizeMBI int) {
	numberOfKeys = numberOfKeysI
	opsPerWorker = opsPerWorkerI
	allowedCacheSizeMB = allowedCacheSizeMBI

	uuidSlice = generateKeySlice(numberOfKeys)

	start("emarcache", numberOfKeys, opsPerWorker, allowedCacheSizeMB)

	emCache = emarcache.New(emarcache.WithMaxSize(uint64(1024 * 1024 * allowedCacheSizeMB)))

	startTime = time.Now()
	testCache(testEmarCache)
	duration = time.Since(startTime)
	fmt.Printf("Emar Cache test completed in %v\n", duration)
	end("emarcache", numberOfKeys, opsPerWorker, allowedCacheSizeMB, duration)

} */

/* func testEmarCache(uuidSlice *[]string) {
	for i := 0; i < opsPerWorker; i++ {
		key := (*uuidSlice)[rand.Intn(numberOfKeys)]

		switch rand.Intn(10) {
		case 0:
			content := getRandomContent()
			emCache.Set(key, content)

		case 1:
			emCache.Delete(key)

		default:
			emCache.Get(key)

		}
	}
} */
