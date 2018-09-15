package usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"encoding/json"

	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
)

var (
	//HTTPServiceType service type of http
	HTTPServiceType ServiceType = "HTTPServiceType"
	//SubsriberEventType service type of subsrciber
	SubsriberEventType ServiceType = "SubscriberType"
)

//UseCase base contract
type UseCase interface {
	Name() string
	HandleUseCase(ctx context.Context, res *usecaseres.UseCaseResource, data *UseCaseData) (interface{}, error)
	Notify(ctx context.Context, res *usecaseres.UseCaseResource, data *UseCaseData) error
	NotifyResult() interface{}
}

//ServiceType as use case service type
type ServiceType string

//HandleUseCase abstraction for use case business operation
type HandleUseCase func(ctx context.Context, res *usecaseres.UseCaseResource, data *UseCaseData) (interface{}, error)

//SubscribeData subscriber data
type SubscribeData struct {
	Data []byte
}

//UseCaseData structure
type UseCaseData struct {
	ServiceType   *ServiceType `json:"data"`
	HTTPData      interface{}
	NSQBody       []byte
	Subscribers   []UseCase
	SubscribeData *SubscribeData
	Language      string
	Arguments     []string // for exec binary only
}

//Cast casts target data based on service
//args:
//  target: target of casted data
//error:
//  error response
func (u *UseCaseData) Cast(target interface{}) error {
	if u.ServiceType == nil {
		return errors.New("service type is not defined")
	}
	switch *u.ServiceType {
	case HTTPServiceType:
		return u.castHTTPRequest(target)
	case SubsriberEventType:
		return u.castSubscribeData(target)
	default:
		return errors.New("unimplemented service type for casting")
	}
}

//castHTTPRequest cast use case's request data to specific structure
//args:
//	target: cast target
//returns:
//	error: operation
func (u *UseCaseData) castHTTPRequest(target interface{}) error {
	if u.HTTPData == nil {
		return errors.New("empty http data")
	}
	byteData, err := json.Marshal(u.HTTPData)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &target)
}

//castExecRequest cast use case's bytes data to specific structure
//args:
//	target: cast target
//returns:
//	error: operation
func (u *UseCaseData) castExecRequest(target interface{}) error {
	if len(u.Arguments) == 0 {
		return errors.New("arguments data len is 0")
	}
	byteData, err := json.Marshal(u.Arguments)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &target)
}

//castNSQBody cast use case's bytes data to specific structure
//args:
//	target: cast target
//returns:
//	error: operation
func (u *UseCaseData) castNSQBody(target interface{}) error {
	if u.NSQBody == nil {
		return errors.New("undefined/ empty nsq body")
	}
	return json.Unmarshal(u.NSQBody, &target)
}

//castSubscribeData cast subscribe data to specific structure
//args:
//	target: cast target
//returns:
//	error: operation
func (u *UseCaseData) castSubscribeData(target interface{}) error {
	return json.Unmarshal(u.SubscribeData.Data, &target)
}

//NotifySubscribers will notify all subscriber with async operation
//args:
//	ctx: context,
//	res: usecase resource,
//	data: data which will be passed to subscriber
//returns:
//	error: operation
func (u *UseCaseData) NotifySubscribers(ctx context.Context, timeout time.Duration, res *usecaseres.UseCaseResource, data interface{}) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	err := u.setSubscribeData(data)
	if err != nil {
		return err
	}
	doneChan := make(chan bool)
	errChan := make(chan error, len(u.Subscribers))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(u.Subscribers))
		for _, s := range u.Subscribers {
			go func(s UseCase) {
				defer wg.Done()
				errChan <- s.Notify(ctx, res, u)
			}(s)

		}
		wg.Wait()
		close(errChan)
		doneChan <- true
	}()
	select {
	case <-timeoutCtx.Done():
		cancel()
		return errors.New("context deadline exceeded")
	case <-doneChan:
		var msgs []string
		for e := range errChan {
			if e != nil {
				msgs = append(msgs, e.Error())
			}
		}
		if len(msgs) > 0 {
			byteData, err := json.Marshal(msgs)
			if err != nil {
				return err
			}
			return fmt.Errorf("fail to notify one of several use case %s",
				string(byteData))
		}
	}
	return err
}

//GetNotifyResultByName will get notify result from specific subscriber
//args:
//	ctx: context,
//  name: name subsriber which will be notified,
//returns:
//	result: Result from notify operation
func (u *UseCaseData) GetNotifyResultByName(ctx context.Context, name string) (result interface{}) {
	for _, s := range u.Subscribers {
		if s.Name() == name {
			result = s.NotifyResult()
			return result
		}
	}
	return nil
}

//GetSubscriberByName gets subscriber by name
//args:
//	name: name of subscriber use case
//returns:
//	filtered use case
func (u *UseCaseData) GetSubscriberByName(name string) UseCase {
	for _, s := range u.Subscribers {
		if s.Name() == name {
			return s
		}
	}
	return nil
}

//setSubscribeData will notify specific subscriber
//args:
//	data: data which will be set on SubscribeData that will be passed when notify a subscriber
//returns:
//	error: operation
func (u *UseCaseData) setSubscribeData(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	u.SubscribeData = new(SubscribeData)
	u.SubscribeData.Data = byteData
	u.ServiceType = &SubsriberEventType
	return nil
}
