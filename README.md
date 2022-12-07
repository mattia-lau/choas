# CheckList
[x] fetch a BTCUSD pair from any publicly available API, once per minute  
[x] periodically store the exchange rate in a sensible database of your choice  
[x] implement API handlers against the database  
[x] get the last price  
[x] get the price at a given timestamp, come up with a way to serve a price if you don't have price at the requested second  
[x] compute the average price in a time range  

# Build
```
docker build -t choas-backend .
docker run -p 8080:8080 choas-backend
```