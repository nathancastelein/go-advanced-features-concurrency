package main

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

type Finder interface {
	Find(resourceName string) bool
	FindWithContext(ctx context.Context, resourceName string) (bool, error)
	FindWithError(ctx context.Context, resourceName string) (bool, error)
	Weight() int64
}

type datacenter struct {
	name          string
	responseDelay time.Duration
	err           error
	weight        int64
}

func (d *datacenter) LogValue() slog.Value {
	return slog.StringValue(d.name)
}

func (d *datacenter) Find(resourceName string) bool {
	time.Sleep(d.responseDelay)
	return true
}

func (d *datacenter) FindWithContext(ctx context.Context, resourceName string) (bool, error) {
	timer := time.NewTimer(d.responseDelay)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		slog.Info("deadline exceeded", slog.Any("finder", d))
		return false, context.DeadlineExceeded
	}
}

func (d *datacenter) FindWithError(ctx context.Context, resourceName string) (bool, error) {
	if d.err != nil {
		return false, d.err
	}
	timer := time.NewTimer(d.responseDelay)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		slog.Info("deadline exceeded", slog.Any("finder", d))
		return false, errors.New("deadline exceeded")
	}
}

func (d *datacenter) Weight() int64 {
	return d.weight
}

func SBG() Finder {
	return &datacenter{
		name:          "SBG",
		responseDelay: 100 * time.Millisecond,
		weight:        10,
	}
}

func GRA() Finder {
	return &datacenter{
		name:          "GRA",
		responseDelay: 200 * time.Millisecond,
		weight:        20,
	}
}

func BHS() Finder {
	return &datacenter{
		name:          "BHS",
		responseDelay: 300 * time.Millisecond,
		weight:        30,
	}
}

func RBX() Finder {
	return &datacenter{
		name:          "RBX",
		responseDelay: 100 * time.Millisecond,
		weight:        10,
	}
}

func WAW() Finder {
	return &datacenter{
		name:          "WAW",
		responseDelay: 100 * time.Millisecond,
		weight:        10,
	}
}

func DCError() Finder {
	return &datacenter{
		name:          "DCError",
		responseDelay: 100 * time.Millisecond,
		err:           errors.New("something went wrong"),
	}
}
