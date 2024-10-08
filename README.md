# Cachetest

Cachetest is my take to compare caching libraries for go4lage. [https://github.com/karl1b/go4lage/]

I don't know if they are so realistic, but better test than no test.


I just wanted to make some performance tests on what I should use for go4lage.
The custom_fast_cache is my take. Pretty impressive what Go can do so simple out of the box.


## Results

|Cache            |Number of Keys|Ops Per Worker|Cache Size (MB)|LoadDuration (s)|Duration (s)|Alloc (MB)|TotalAlloc (MB)|Sys (MB)|NumGC|NumCPU|NumGoroutine|GOMAXPROCS|
|-----------------|--------------|--------------|---------------|----------------|------------|----------|---------------|--------|-----|------|------------|----------|
|custom_cache     |100000        |100000        |100            |0.86            |0.74        |8         |48             |30      |8    |12    |1           |12        |
|custom_fast_cache|100000        |100000        |100            |0.87            |0.51        |8         |47             |26      |8    |12    |1           |12        |
|go_cache         |100000        |100000        |100            |0.92            |0.57        |15        |145            |46      |14   |12    |1           |12        |
|free_cache       |100000        |100000        |100            |0.4             |1.46        |9         |673            |35      |70   |12    |1           |12        |
|ristretto_cache  |100000        |100000        |100            |0.44            |0.87        |9         |674            |43      |61   |12    |3           |12        |
|custom_cache     |100000        |10000000      |100            |0.85            |78.45       |8         |440            |30      |58   |12    |1           |12        |
|custom_fast_cache|100000        |10000000      |100            |0.89            |49.98       |8         |432            |30      |57   |12    |1           |12        |
|go_cache         |100000        |10000000      |100            |0.96            |63.63       |15        |1234           |47      |89   |12    |1           |12        |
|free_cache       |100000        |10000000      |100            |0.4             |143.79      |9         |44183          |35      |4164 |12    |1           |12        |
|ristretto_cache  |100000        |10000000      |100            |0.44            |90.07       |9         |17016          |42      |1727 |12    |3           |12        |
|custom_cache     |10000         |10000000      |100            |0.08            |73.05       |2         |237            |18      |85   |12    |1           |12        |
|custom_fast_cache|10000         |10000000      |100            |0.08            |41.1        |2         |232            |18      |87   |12    |1           |12        |
|go_cache         |10000         |10000000      |100            |0.09            |37.22       |3         |1118           |18      |342  |12    |1           |12        |
|free_cache       |10000         |10000000      |100            |0.04            |138.32      |4         |43979          |26      |8729 |12    |1           |12        |
|ristretto_cache  |10000         |10000000      |100            |0.05            |84.19       |3         |15911          |30      |4274 |12    |3           |12        |
|custom_cache     |100000        |50000000      |100            |0.85            |387.89      |8         |2053           |30      |252  |12    |1           |12        |
|custom_fast_cache|100000        |50000000      |100            |0.87            |237.42      |8         |2033           |30      |257  |12    |1           |12        |
|go_cache         |100000        |50000000      |100            |0.93            |329.9       |15        |5632           |47      |388  |12    |1           |12        |
|free_cache       |100000        |50000000      |100            |0.4             |713.56      |9         |219970         |42      |19510|12    |1           |12        |
|ristretto_cache  |100000        |50000000      |100            |0.44            |444.75      |9         |82961          |47      |7894 |12    |3           |12        |
|custom_cache     |1000000       |50000000      |100            |8.77            |687.67      |63        |2416           |140     |43   |12    |1           |12        |
|custom_fast_cache|1000000       |50000000      |100            |8.93            |627.83      |63        |2411           |140     |44   |12    |1           |12        |
|go_cache         |1000000       |50000000      |100            |10.72           |751.26      |163       |6868           |466     |49   |12    |1           |12        |
|free_cache       |1000000       |50000000      |100            |4               |802.69      |64        |222014         |153     |3379 |12    |1           |12        |
|ristretto_cache  |1000000       |50000000      |100            |4.29            |548.13      |69        |87972          |178     |1293 |12    |3           |12        |
|custom_cache     |10000         |50000000      |250            |0.08            |403.91      |3         |15             |18      |8    |12    |1           |12        |
|custom_fast_cache|10000         |50000000      |250            |0.08            |211.82      |3         |10             |18      |5    |12    |1           |12        |
|go_cache         |10000         |50000000      |250            |0.09            |185.1       |3         |5514           |22      |1616 |12    |1           |12        |
|free_cache       |10000         |50000000      |250            |0.04            |697.19      |4         |219767         |38      |38087|12    |1           |12        |
|ristretto_cache  |10000         |50000000      |250            |0.05            |427.12      |4         |68744          |30      |13970|12    |3           |12        |


## Not shown

https://github.com/emar-kar/cache was tested and worked, but It was too slow by far. (Maybe I used it wrong)
https://github.com/orcaman/concurrent-map did crash without an error. Maybe I was also too stupid to use it correctly ¯\_(ツ)_/¯