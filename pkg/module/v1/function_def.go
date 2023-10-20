package module

type DiscoveryFunction func(conf *Config, spec []TypeSpec) (error, []File)
