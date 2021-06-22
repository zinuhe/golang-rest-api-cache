// www.restapiexample.com/golang-tutorial/how-to-set-and-get-cache-using-golang-rest-api/
// https://adityarama1210.medium.com/fast-golang-api-performance-with-in-memory-key-value-storing-cache-1b248c182bdb

// go get github.com/patrickmn/go-cache

apiV1.Handle("/get-cache", controller.GetCache).Methods("GET")
