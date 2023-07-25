package end

type CacheGet interface {
	Get() (error, interface{})
}

type CacheSet interface {
	Set(key string, value interface{}) (error, interface{})
}
