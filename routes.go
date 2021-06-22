// www.restapiexample.com/golang-tutorial/how-to-set-and-get-cache-using-golang-rest-api/

// go get github.com/patrickmn/go-cache

apiV1.Handle("/get-cache", controller.GetCache).Methods("GET")
