# TimeoutClient

A go helper for timeouting http.Client

License [Beerware](http://en.wikipedia.org/wiki/Beerware)

# Usage
```
// Create a new client
client := timeoutclient.NewTimeoutClient(3000*time.Millisecond, 3000*time.Millisecond)
// Use as normal http.client
resp, err := client.Get("http://example.com")
```
Original idea && credits go to : [Tanel Lebedev](https://github.com/tanel)
