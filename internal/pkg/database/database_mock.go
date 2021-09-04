// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package database

import (
	"context"
	"github.com/diwise/api-notify/pkg/models"
	"sync"
)

// Ensure, that DbMock does implement Db.
// If this is not the case, regenerate this file with moq.
var _ Db = &DbMock{}

// DbMock is a mock implementation of Db.
//
// 	func TestSomethingThatUsesDb(t *testing.T) {
//
// 		// make and configure a mocked Db
// 		mockedDb := &DbMock{
// 			CreateSubscriptionFunc: func(ctx context.Context, subscription *models.Subscription) error {
// 				panic("mock out the CreateSubscription method")
// 			},
// 			DeleteSubscriptionFunc: func(ctx context.Context, subscriptionId string) error {
// 				panic("mock out the DeleteSubscription method")
// 			},
// 			GetSubscriptionByIdFunc: func(subscriptionId string) (*models.Subscription, error) {
// 				panic("mock out the GetSubscriptionById method")
// 			},
// 			GetSubscriptionsByIdOrTypeFunc: func(ctx context.Context, id string, entityType string) ([]models.Subscription, error) {
// 				panic("mock out the GetSubscriptionsByIdOrType method")
// 			},
// 			ListSubscriptionsFunc: func(ctx context.Context, limit int) ([]models.Subscription, error) {
// 				panic("mock out the ListSubscriptions method")
// 			},
// 			UpdateSubscriptionFunc: func(ctx context.Context, subscription *models.Subscription) error {
// 				panic("mock out the UpdateSubscription method")
// 			},
// 		}
//
// 		// use mockedDb in code that requires Db
// 		// and then make assertions.
//
// 	}
type DbMock struct {
	// CreateSubscriptionFunc mocks the CreateSubscription method.
	CreateSubscriptionFunc func(ctx context.Context, subscription *models.Subscription) error

	// DeleteSubscriptionFunc mocks the DeleteSubscription method.
	DeleteSubscriptionFunc func(ctx context.Context, subscriptionId string) error

	// GetSubscriptionByIdFunc mocks the GetSubscriptionById method.
	GetSubscriptionByIdFunc func(subscriptionId string) (*models.Subscription, error)

	// GetSubscriptionsByIdOrTypeFunc mocks the GetSubscriptionsByIdOrType method.
	GetSubscriptionsByIdOrTypeFunc func(ctx context.Context, id string, entityType string) ([]models.Subscription, error)

	// ListSubscriptionsFunc mocks the ListSubscriptions method.
	ListSubscriptionsFunc func(ctx context.Context, limit int) ([]models.Subscription, error)

	// UpdateSubscriptionFunc mocks the UpdateSubscription method.
	UpdateSubscriptionFunc func(ctx context.Context, subscription *models.Subscription) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateSubscription holds details about calls to the CreateSubscription method.
		CreateSubscription []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Subscription is the subscription argument value.
			Subscription *models.Subscription
		}
		// DeleteSubscription holds details about calls to the DeleteSubscription method.
		DeleteSubscription []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SubscriptionId is the subscriptionId argument value.
			SubscriptionId string
		}
		// GetSubscriptionById holds details about calls to the GetSubscriptionById method.
		GetSubscriptionById []struct {
			// SubscriptionId is the subscriptionId argument value.
			SubscriptionId string
		}
		// GetSubscriptionsByIdOrType holds details about calls to the GetSubscriptionsByIdOrType method.
		GetSubscriptionsByIdOrType []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// EntityType is the entityType argument value.
			EntityType string
		}
		// ListSubscriptions holds details about calls to the ListSubscriptions method.
		ListSubscriptions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Limit is the limit argument value.
			Limit int
		}
		// UpdateSubscription holds details about calls to the UpdateSubscription method.
		UpdateSubscription []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Subscription is the subscription argument value.
			Subscription *models.Subscription
		}
	}
	lockCreateSubscription         sync.RWMutex
	lockDeleteSubscription         sync.RWMutex
	lockGetSubscriptionById        sync.RWMutex
	lockGetSubscriptionsByIdOrType sync.RWMutex
	lockListSubscriptions          sync.RWMutex
	lockUpdateSubscription         sync.RWMutex
}

// CreateSubscription calls CreateSubscriptionFunc.
func (mock *DbMock) CreateSubscription(ctx context.Context, subscription *models.Subscription) error {
	callInfo := struct {
		Ctx          context.Context
		Subscription *models.Subscription
	}{
		Ctx:          ctx,
		Subscription: subscription,
	}
	mock.lockCreateSubscription.Lock()
	mock.calls.CreateSubscription = append(mock.calls.CreateSubscription, callInfo)
	mock.lockCreateSubscription.Unlock()
	if mock.CreateSubscriptionFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CreateSubscriptionFunc(ctx, subscription)
}

// CreateSubscriptionCalls gets all the calls that were made to CreateSubscription.
// Check the length with:
//     len(mockedDb.CreateSubscriptionCalls())
func (mock *DbMock) CreateSubscriptionCalls() []struct {
	Ctx          context.Context
	Subscription *models.Subscription
} {
	var calls []struct {
		Ctx          context.Context
		Subscription *models.Subscription
	}
	mock.lockCreateSubscription.RLock()
	calls = mock.calls.CreateSubscription
	mock.lockCreateSubscription.RUnlock()
	return calls
}

// DeleteSubscription calls DeleteSubscriptionFunc.
func (mock *DbMock) DeleteSubscription(ctx context.Context, subscriptionId string) error {
	callInfo := struct {
		Ctx            context.Context
		SubscriptionId string
	}{
		Ctx:            ctx,
		SubscriptionId: subscriptionId,
	}
	mock.lockDeleteSubscription.Lock()
	mock.calls.DeleteSubscription = append(mock.calls.DeleteSubscription, callInfo)
	mock.lockDeleteSubscription.Unlock()
	if mock.DeleteSubscriptionFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.DeleteSubscriptionFunc(ctx, subscriptionId)
}

// DeleteSubscriptionCalls gets all the calls that were made to DeleteSubscription.
// Check the length with:
//     len(mockedDb.DeleteSubscriptionCalls())
func (mock *DbMock) DeleteSubscriptionCalls() []struct {
	Ctx            context.Context
	SubscriptionId string
} {
	var calls []struct {
		Ctx            context.Context
		SubscriptionId string
	}
	mock.lockDeleteSubscription.RLock()
	calls = mock.calls.DeleteSubscription
	mock.lockDeleteSubscription.RUnlock()
	return calls
}

// GetSubscriptionById calls GetSubscriptionByIdFunc.
func (mock *DbMock) GetSubscriptionById(subscriptionId string) (*models.Subscription, error) {
	callInfo := struct {
		SubscriptionId string
	}{
		SubscriptionId: subscriptionId,
	}
	mock.lockGetSubscriptionById.Lock()
	mock.calls.GetSubscriptionById = append(mock.calls.GetSubscriptionById, callInfo)
	mock.lockGetSubscriptionById.Unlock()
	if mock.GetSubscriptionByIdFunc == nil {
		var (
			subscriptionOut *models.Subscription
			errOut          error
		)
		return subscriptionOut, errOut
	}
	return mock.GetSubscriptionByIdFunc(subscriptionId)
}

// GetSubscriptionByIdCalls gets all the calls that were made to GetSubscriptionById.
// Check the length with:
//     len(mockedDb.GetSubscriptionByIdCalls())
func (mock *DbMock) GetSubscriptionByIdCalls() []struct {
	SubscriptionId string
} {
	var calls []struct {
		SubscriptionId string
	}
	mock.lockGetSubscriptionById.RLock()
	calls = mock.calls.GetSubscriptionById
	mock.lockGetSubscriptionById.RUnlock()
	return calls
}

// GetSubscriptionsByIdOrType calls GetSubscriptionsByIdOrTypeFunc.
func (mock *DbMock) GetSubscriptionsByIdOrType(ctx context.Context, id string, entityType string) ([]models.Subscription, error) {
	callInfo := struct {
		Ctx        context.Context
		ID         string
		EntityType string
	}{
		Ctx:        ctx,
		ID:         id,
		EntityType: entityType,
	}
	mock.lockGetSubscriptionsByIdOrType.Lock()
	mock.calls.GetSubscriptionsByIdOrType = append(mock.calls.GetSubscriptionsByIdOrType, callInfo)
	mock.lockGetSubscriptionsByIdOrType.Unlock()
	if mock.GetSubscriptionsByIdOrTypeFunc == nil {
		var (
			subscriptionsOut []models.Subscription
			errOut           error
		)
		return subscriptionsOut, errOut
	}
	return mock.GetSubscriptionsByIdOrTypeFunc(ctx, id, entityType)
}

// GetSubscriptionsByIdOrTypeCalls gets all the calls that were made to GetSubscriptionsByIdOrType.
// Check the length with:
//     len(mockedDb.GetSubscriptionsByIdOrTypeCalls())
func (mock *DbMock) GetSubscriptionsByIdOrTypeCalls() []struct {
	Ctx        context.Context
	ID         string
	EntityType string
} {
	var calls []struct {
		Ctx        context.Context
		ID         string
		EntityType string
	}
	mock.lockGetSubscriptionsByIdOrType.RLock()
	calls = mock.calls.GetSubscriptionsByIdOrType
	mock.lockGetSubscriptionsByIdOrType.RUnlock()
	return calls
}

// ListSubscriptions calls ListSubscriptionsFunc.
func (mock *DbMock) ListSubscriptions(ctx context.Context, limit int) ([]models.Subscription, error) {
	callInfo := struct {
		Ctx   context.Context
		Limit int
	}{
		Ctx:   ctx,
		Limit: limit,
	}
	mock.lockListSubscriptions.Lock()
	mock.calls.ListSubscriptions = append(mock.calls.ListSubscriptions, callInfo)
	mock.lockListSubscriptions.Unlock()
	if mock.ListSubscriptionsFunc == nil {
		var (
			subscriptionsOut []models.Subscription
			errOut           error
		)
		return subscriptionsOut, errOut
	}
	return mock.ListSubscriptionsFunc(ctx, limit)
}

// ListSubscriptionsCalls gets all the calls that were made to ListSubscriptions.
// Check the length with:
//     len(mockedDb.ListSubscriptionsCalls())
func (mock *DbMock) ListSubscriptionsCalls() []struct {
	Ctx   context.Context
	Limit int
} {
	var calls []struct {
		Ctx   context.Context
		Limit int
	}
	mock.lockListSubscriptions.RLock()
	calls = mock.calls.ListSubscriptions
	mock.lockListSubscriptions.RUnlock()
	return calls
}

// UpdateSubscription calls UpdateSubscriptionFunc.
func (mock *DbMock) UpdateSubscription(ctx context.Context, subscription *models.Subscription) error {
	callInfo := struct {
		Ctx          context.Context
		Subscription *models.Subscription
	}{
		Ctx:          ctx,
		Subscription: subscription,
	}
	mock.lockUpdateSubscription.Lock()
	mock.calls.UpdateSubscription = append(mock.calls.UpdateSubscription, callInfo)
	mock.lockUpdateSubscription.Unlock()
	if mock.UpdateSubscriptionFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.UpdateSubscriptionFunc(ctx, subscription)
}

// UpdateSubscriptionCalls gets all the calls that were made to UpdateSubscription.
// Check the length with:
//     len(mockedDb.UpdateSubscriptionCalls())
func (mock *DbMock) UpdateSubscriptionCalls() []struct {
	Ctx          context.Context
	Subscription *models.Subscription
} {
	var calls []struct {
		Ctx          context.Context
		Subscription *models.Subscription
	}
	mock.lockUpdateSubscription.RLock()
	calls = mock.calls.UpdateSubscription
	mock.lockUpdateSubscription.RUnlock()
	return calls
}