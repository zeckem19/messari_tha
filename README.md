Defines 3 structs - 
1. `Trade` struct is used to validate field names and data type from the binary. 
2. `Store` struct keeps track of the current data seen for each market
3. `PrettyData` struct is used to format the market data after all trades are completed

A fixed-length (13000) array `record` is used which has an index that corresponds to the `market` being traded.

At first I took an approach where I tried to minimize the amount of computation for each trade and generate expected output (at the end) after the 10 million trades are completed. This made sense since there were more trades than markets. To do this, I determined that `Store` only needs 4 fields 

1. number of transactions
2. the total volume
3. the total price for that trade
4. number of buy orders. 
   
The expected output can be computed from these 4 sets of information

1. total_volume (self explanatory)
2. mean_price: total_price / num_of_transactions
3. mean_volume: total_volume / num_of_transactions
4. Volume_weighted_average_price: total_price / total_volume
5. Percentage_buy: num_buy / num_of_transactions

However, computing at the end meant that the stats would not be available until all the trades are completed. I then changed my approach to keep track of these in real-time. Hence, `Store` would need the total_price field and some of the above logic would need to be done in the `process` function. The `process` function thus updates all the above fields, on top of the number of transactions and buy orders.

Results 

Running just the binary
```bash
# time  go run main.go > /dev/null 
go run main.go > /dev/null  7.42s user 1.52s system 99% cpu 8.997 total
```

Running with my code
```bash
# time go run main.go | go run ../main.go > /dev/null
go run main.go > /dev/null  15.01s user 0.59s system 103% cpu 15.049 total

```

Areas of improvement
1. The `record` array has a fixed size of 13000, the app could be improved to dynamically increase its size when new markets are added. Alternatively, a map could be used instead, but would run the risk of collisions when the number of markets grow.
2. Scaling the app - With more trades, a few approaches could be considered to scale with the volume.
   1. Batching computations - if new stats are added that might require more complex calculations, they could be computed on demand
   2. Parallelizing operations - as volume of trades increase, we could parallelize across multiple processes via distributing the computation around `market` or round robin with Id. However that would require a thread safe data structure

