package profile

import (
	"time"

	ruvixapi "github.com/cagodoy/ruvix-api"
)

// FactoryCreate ...
func FactoryCreate(addr string, options ruvixapi.ClientOptions) (*Profile, *Profile, error) {
	pc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	before := &Profile{
		Fingerprint:   "",
		Age:           18,
		Birth:         2000,
		MaritalStatus: "single",
		Childs:        0,
		Gender:        "male",
	}

	after, err := pc.Create(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryGet ...
func FactoryGet(addr string, options ruvixapi.ClientOptions) (*Profile, *Profile, error) {
	pc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	before := &Profile{
		Fingerprint:   "",
		Age:           18,
		Birth:         2000,
		MaritalStatus: "single",
		Childs:        0,
		Gender:        "male",
	}

	_, err = pc.Create(before)
	if err != nil {
		return nil, nil, err
	}

	after, err := pc.Get()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryUpdate ...
func FactoryUpdate(addr string, options ruvixapi.ClientOptions) (*Profile, *Profile, error) {
	pc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	_, before, err := FactoryCreate(addr, options)
	if err != nil {
		return nil, nil, err
	}

	time.Sleep(1 * time.Second)

	after, err := pc.Update(&Profile{
		Fingerprint:   "",
		Age:           31,
		Birth:         2005,
		MaritalStatus: "married",
		Childs:        2,
		Gender:        "female",
	})
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}
