Defines 3 structs - 
1. `Trade` struct is used to validate field names and data type from the binary. 
2. `Store` struct keeps track of the current data seen for each market
3. `Result` struct is used after all the trades are completed, it formats Store into the expected output

I also used a fixed-length (13000) array `record`, which has an index that corresponds to the `market` being traded.

At first I took an approach where I tried to minimize the amount of computation for each trade and generate expected output after the 10 million trades are completed. This made sense since there were much more trades than markets. To do this, I determined that `Store` only needs 4 fields - number of transactions, the total volume, the total price for that trade and number of buy orders. The expected output can be computed from these 4 sets of information

1. total_volume (self explanatory)
2. mean_price: total_price / num_of_transactions
3. mean_volume: total_volume / num_of_transactions
4. Volume_weighted_average_price: total_price / total_volume
5. Percentage_buy: num_buy / num_of_transactions

However, computing at the end meant that the stats would not be available until all the trades are completed. I then changed my approach to keep track of these in real-time. Hence, `Store` would need the total_price field and some of the above logic would need to be done in the `process` function, the `Result` struct would also become redundant.


Areas of improvement
1. The `record` array has a fixed size of 13000, the app could be improved to dynamically increase its size when new markets are added. Alternatively, a map could be used instead, but would run the risk of collisions when the number of markets grow.
2. Scaling the app - With more trades, a few approaches could be considered to scale with the volume.
   1. Batching computations - 