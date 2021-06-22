package helper
 
import (
    "github.com/patrickmn/go-cache"
    "time"
    "fmt"
    "strings"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)
type Emp []struct {
    ID        int    `json:"id"`
    Name string `json:"name"`
}

func SetCache(key string, emp interface{}) bool {
    Cache.Set(key, emp, cache.NoExpiration)
    return true
}
