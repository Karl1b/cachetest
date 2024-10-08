
#!/bin/bash
go build
# aprox 1 min per run
#numberOfKeys=100000
#opsPerWorker=10000000
#allowedCacheSizeMB=100



numberOfKeys=100000
opsPerWorker=100000
allowedCacheSizeMB=100


#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB


numberOfKeys=100000
opsPerWorker=10000000
allowedCacheSizeMB=100


#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB


numberOfKeys=10000
opsPerWorker=10000000
allowedCacheSizeMB=100


#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB


#numberOfKeys=100000
#opsPerWorker=50000000
#allowedCacheSizeMB=100

numberOfKeys=100000
opsPerWorker=50000000
allowedCacheSizeMB=100

#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB

numberOfKeys=1000000
opsPerWorker=50000000
allowedCacheSizeMB=100

#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB

numberOfKeys=10000
opsPerWorker=50000000
allowedCacheSizeMB=250

#numberOfKeys, opsPerWorker , allowedCacheSizeMB 
./cachetest custom $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest customfast $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest gocache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest freecache $numberOfKeys $opsPerWorker $allowedCacheSizeMB
./cachetest ristretto $numberOfKeys $opsPerWorker $allowedCacheSizeMB