package go_design_patterns

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
