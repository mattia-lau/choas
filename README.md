# CheckList
[x] fetch a BTCUSD pair from any publicly available API, once per minute  
[x] periodically store the exchange rate in a sensible database of your choice  
[x] implement API handlers against the database  
[x] get the last price  
[x] get the price at a given timestamp, come up with a way to serve a price if you don't have price at the requested second  
[x] compute the average price in a time range  

# Secrets
1. cp .env.example .env

# Polygon
1. Create the account in https://polygon.io/
2. Copy API Key from https://polygon.io/dashboard/api-keys
3. replace `POLYGON_API_KEY=<Your api key>` in .env
