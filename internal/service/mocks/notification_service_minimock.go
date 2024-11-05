// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/knoxie-s/notification-service/internal/service.NotificationService -o notification_service_minimock.go -n NotificationServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/knoxie-s/notification-service/internal/model"
)

// NotificationServiceMock implements mm_service.NotificationService
type NotificationServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, notificationInfo *model.NotificationInfo) (np1 *model.Notification, err error)
	funcCreateOrigin    string
	inspectFuncCreate   func(ctx context.Context, notificationInfo *model.NotificationInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mNotificationServiceMockCreate

	funcSendNotificationsToClients          func(ctx context.Context) (err error)
	funcSendNotificationsToClientsOrigin    string
	inspectFuncSendNotificationsToClients   func(ctx context.Context)
	afterSendNotificationsToClientsCounter  uint64
	beforeSendNotificationsToClientsCounter uint64
	SendNotificationsToClientsMock          mNotificationServiceMockSendNotificationsToClients
}

// NewNotificationServiceMock returns a mock for mm_service.NotificationService
func NewNotificationServiceMock(t minimock.Tester) *NotificationServiceMock {
	m := &NotificationServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mNotificationServiceMockCreate{mock: m}
	m.CreateMock.callArgs = []*NotificationServiceMockCreateParams{}

	m.SendNotificationsToClientsMock = mNotificationServiceMockSendNotificationsToClients{mock: m}
	m.SendNotificationsToClientsMock.callArgs = []*NotificationServiceMockSendNotificationsToClientsParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mNotificationServiceMockCreate struct {
	optional           bool
	mock               *NotificationServiceMock
	defaultExpectation *NotificationServiceMockCreateExpectation
	expectations       []*NotificationServiceMockCreateExpectation

	callArgs []*NotificationServiceMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// NotificationServiceMockCreateExpectation specifies expectation struct of the NotificationService.Create
type NotificationServiceMockCreateExpectation struct {
	mock               *NotificationServiceMock
	params             *NotificationServiceMockCreateParams
	paramPtrs          *NotificationServiceMockCreateParamPtrs
	expectationOrigins NotificationServiceMockCreateExpectationOrigins
	results            *NotificationServiceMockCreateResults
	returnOrigin       string
	Counter            uint64
}

// NotificationServiceMockCreateParams contains parameters of the NotificationService.Create
type NotificationServiceMockCreateParams struct {
	ctx              context.Context
	notificationInfo *model.NotificationInfo
}

// NotificationServiceMockCreateParamPtrs contains pointers to parameters of the NotificationService.Create
type NotificationServiceMockCreateParamPtrs struct {
	ctx              *context.Context
	notificationInfo **model.NotificationInfo
}

// NotificationServiceMockCreateResults contains results of the NotificationService.Create
type NotificationServiceMockCreateResults struct {
	np1 *model.Notification
	err error
}

// NotificationServiceMockCreateOrigins contains origins of expectations of the NotificationService.Create
type NotificationServiceMockCreateExpectationOrigins struct {
	origin                 string
	originCtx              string
	originNotificationInfo string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mNotificationServiceMockCreate) Optional() *mNotificationServiceMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for NotificationService.Create
func (mmCreate *mNotificationServiceMockCreate) Expect(ctx context.Context, notificationInfo *model.NotificationInfo) *mNotificationServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NotificationServiceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &NotificationServiceMockCreateParams{ctx, notificationInfo}
	mmCreate.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for NotificationService.Create
func (mmCreate *mNotificationServiceMockCreate) ExpectCtxParam1(ctx context.Context) *mNotificationServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NotificationServiceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &NotificationServiceMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreate.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreate
}

// ExpectNotificationInfoParam2 sets up expected param notificationInfo for NotificationService.Create
func (mmCreate *mNotificationServiceMockCreate) ExpectNotificationInfoParam2(notificationInfo *model.NotificationInfo) *mNotificationServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NotificationServiceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &NotificationServiceMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.notificationInfo = &notificationInfo
	mmCreate.defaultExpectation.expectationOrigins.originNotificationInfo = minimock.CallerInfo(1)

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the NotificationService.Create
func (mmCreate *mNotificationServiceMockCreate) Inspect(f func(ctx context.Context, notificationInfo *model.NotificationInfo)) *mNotificationServiceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for NotificationServiceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by NotificationService.Create
func (mmCreate *mNotificationServiceMockCreate) Return(np1 *model.Notification, err error) *NotificationServiceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NotificationServiceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &NotificationServiceMockCreateResults{np1, err}
	mmCreate.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// Set uses given function f to mock the NotificationService.Create method
func (mmCreate *mNotificationServiceMockCreate) Set(f func(ctx context.Context, notificationInfo *model.NotificationInfo) (np1 *model.Notification, err error)) *NotificationServiceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the NotificationService.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the NotificationService.Create method")
	}

	mmCreate.mock.funcCreate = f
	mmCreate.mock.funcCreateOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// When sets expectation for the NotificationService.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mNotificationServiceMockCreate) When(ctx context.Context, notificationInfo *model.NotificationInfo) *NotificationServiceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NotificationServiceMock.Create mock is already set by Set")
	}

	expectation := &NotificationServiceMockCreateExpectation{
		mock:               mmCreate.mock,
		params:             &NotificationServiceMockCreateParams{ctx, notificationInfo},
		expectationOrigins: NotificationServiceMockCreateExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up NotificationService.Create return parameters for the expectation previously defined by the When method
func (e *NotificationServiceMockCreateExpectation) Then(np1 *model.Notification, err error) *NotificationServiceMock {
	e.results = &NotificationServiceMockCreateResults{np1, err}
	return e.mock
}

// Times sets number of times NotificationService.Create should be invoked
func (mmCreate *mNotificationServiceMockCreate) Times(n uint64) *mNotificationServiceMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of NotificationServiceMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	mmCreate.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreate
}

func (mmCreate *mNotificationServiceMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements mm_service.NotificationService
func (mmCreate *NotificationServiceMock) Create(ctx context.Context, notificationInfo *model.NotificationInfo) (np1 *model.Notification, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	mmCreate.t.Helper()

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, notificationInfo)
	}

	mm_params := NotificationServiceMockCreateParams{ctx, notificationInfo}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.np1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := NotificationServiceMockCreateParams{ctx, notificationInfo}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("NotificationServiceMock.Create got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.notificationInfo != nil && !minimock.Equal(*mm_want_ptrs.notificationInfo, mm_got.notificationInfo) {
				mmCreate.t.Errorf("NotificationServiceMock.Create got unexpected parameter notificationInfo, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originNotificationInfo, *mm_want_ptrs.notificationInfo, mm_got.notificationInfo, minimock.Diff(*mm_want_ptrs.notificationInfo, mm_got.notificationInfo))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("NotificationServiceMock.Create got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreate.CreateMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the NotificationServiceMock.Create")
		}
		return (*mm_results).np1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, notificationInfo)
	}
	mmCreate.t.Fatalf("Unexpected call to NotificationServiceMock.Create. %v %v", ctx, notificationInfo)
	return
}

// CreateAfterCounter returns a count of finished NotificationServiceMock.Create invocations
func (mmCreate *NotificationServiceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of NotificationServiceMock.Create invocations
func (mmCreate *NotificationServiceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to NotificationServiceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mNotificationServiceMockCreate) Calls() []*NotificationServiceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*NotificationServiceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *NotificationServiceMock) MinimockCreateDone() bool {
	if m.CreateMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMock.invocationsDone()
}

// MinimockCreateInspect logs each unmet expectation
func (m *NotificationServiceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NotificationServiceMock.Create at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to NotificationServiceMock.Create at\n%s", m.CreateMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to NotificationServiceMock.Create at\n%s with params: %#v", m.CreateMock.defaultExpectation.expectationOrigins.origin, *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Errorf("Expected call to NotificationServiceMock.Create at\n%s", m.funcCreateOrigin)
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to NotificationServiceMock.Create at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), m.CreateMock.expectedInvocationsOrigin, afterCreateCounter)
	}
}

type mNotificationServiceMockSendNotificationsToClients struct {
	optional           bool
	mock               *NotificationServiceMock
	defaultExpectation *NotificationServiceMockSendNotificationsToClientsExpectation
	expectations       []*NotificationServiceMockSendNotificationsToClientsExpectation

	callArgs []*NotificationServiceMockSendNotificationsToClientsParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// NotificationServiceMockSendNotificationsToClientsExpectation specifies expectation struct of the NotificationService.SendNotificationsToClients
type NotificationServiceMockSendNotificationsToClientsExpectation struct {
	mock               *NotificationServiceMock
	params             *NotificationServiceMockSendNotificationsToClientsParams
	paramPtrs          *NotificationServiceMockSendNotificationsToClientsParamPtrs
	expectationOrigins NotificationServiceMockSendNotificationsToClientsExpectationOrigins
	results            *NotificationServiceMockSendNotificationsToClientsResults
	returnOrigin       string
	Counter            uint64
}

// NotificationServiceMockSendNotificationsToClientsParams contains parameters of the NotificationService.SendNotificationsToClients
type NotificationServiceMockSendNotificationsToClientsParams struct {
	ctx context.Context
}

// NotificationServiceMockSendNotificationsToClientsParamPtrs contains pointers to parameters of the NotificationService.SendNotificationsToClients
type NotificationServiceMockSendNotificationsToClientsParamPtrs struct {
	ctx *context.Context
}

// NotificationServiceMockSendNotificationsToClientsResults contains results of the NotificationService.SendNotificationsToClients
type NotificationServiceMockSendNotificationsToClientsResults struct {
	err error
}

// NotificationServiceMockSendNotificationsToClientsOrigins contains origins of expectations of the NotificationService.SendNotificationsToClients
type NotificationServiceMockSendNotificationsToClientsExpectationOrigins struct {
	origin    string
	originCtx string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Optional() *mNotificationServiceMockSendNotificationsToClients {
	mmSendNotificationsToClients.optional = true
	return mmSendNotificationsToClients
}

// Expect sets up expected params for NotificationService.SendNotificationsToClients
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Expect(ctx context.Context) *mNotificationServiceMockSendNotificationsToClients {
	if mmSendNotificationsToClients.mock.funcSendNotificationsToClients != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by Set")
	}

	if mmSendNotificationsToClients.defaultExpectation == nil {
		mmSendNotificationsToClients.defaultExpectation = &NotificationServiceMockSendNotificationsToClientsExpectation{}
	}

	if mmSendNotificationsToClients.defaultExpectation.paramPtrs != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by ExpectParams functions")
	}

	mmSendNotificationsToClients.defaultExpectation.params = &NotificationServiceMockSendNotificationsToClientsParams{ctx}
	mmSendNotificationsToClients.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSendNotificationsToClients.expectations {
		if minimock.Equal(e.params, mmSendNotificationsToClients.defaultExpectation.params) {
			mmSendNotificationsToClients.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendNotificationsToClients.defaultExpectation.params)
		}
	}

	return mmSendNotificationsToClients
}

// ExpectCtxParam1 sets up expected param ctx for NotificationService.SendNotificationsToClients
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) ExpectCtxParam1(ctx context.Context) *mNotificationServiceMockSendNotificationsToClients {
	if mmSendNotificationsToClients.mock.funcSendNotificationsToClients != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by Set")
	}

	if mmSendNotificationsToClients.defaultExpectation == nil {
		mmSendNotificationsToClients.defaultExpectation = &NotificationServiceMockSendNotificationsToClientsExpectation{}
	}

	if mmSendNotificationsToClients.defaultExpectation.params != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by Expect")
	}

	if mmSendNotificationsToClients.defaultExpectation.paramPtrs == nil {
		mmSendNotificationsToClients.defaultExpectation.paramPtrs = &NotificationServiceMockSendNotificationsToClientsParamPtrs{}
	}
	mmSendNotificationsToClients.defaultExpectation.paramPtrs.ctx = &ctx
	mmSendNotificationsToClients.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmSendNotificationsToClients
}

// Inspect accepts an inspector function that has same arguments as the NotificationService.SendNotificationsToClients
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Inspect(f func(ctx context.Context)) *mNotificationServiceMockSendNotificationsToClients {
	if mmSendNotificationsToClients.mock.inspectFuncSendNotificationsToClients != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("Inspect function is already set for NotificationServiceMock.SendNotificationsToClients")
	}

	mmSendNotificationsToClients.mock.inspectFuncSendNotificationsToClients = f

	return mmSendNotificationsToClients
}

// Return sets up results that will be returned by NotificationService.SendNotificationsToClients
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Return(err error) *NotificationServiceMock {
	if mmSendNotificationsToClients.mock.funcSendNotificationsToClients != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by Set")
	}

	if mmSendNotificationsToClients.defaultExpectation == nil {
		mmSendNotificationsToClients.defaultExpectation = &NotificationServiceMockSendNotificationsToClientsExpectation{mock: mmSendNotificationsToClients.mock}
	}
	mmSendNotificationsToClients.defaultExpectation.results = &NotificationServiceMockSendNotificationsToClientsResults{err}
	mmSendNotificationsToClients.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSendNotificationsToClients.mock
}

// Set uses given function f to mock the NotificationService.SendNotificationsToClients method
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Set(f func(ctx context.Context) (err error)) *NotificationServiceMock {
	if mmSendNotificationsToClients.defaultExpectation != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("Default expectation is already set for the NotificationService.SendNotificationsToClients method")
	}

	if len(mmSendNotificationsToClients.expectations) > 0 {
		mmSendNotificationsToClients.mock.t.Fatalf("Some expectations are already set for the NotificationService.SendNotificationsToClients method")
	}

	mmSendNotificationsToClients.mock.funcSendNotificationsToClients = f
	mmSendNotificationsToClients.mock.funcSendNotificationsToClientsOrigin = minimock.CallerInfo(1)
	return mmSendNotificationsToClients.mock
}

// When sets expectation for the NotificationService.SendNotificationsToClients which will trigger the result defined by the following
// Then helper
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) When(ctx context.Context) *NotificationServiceMockSendNotificationsToClientsExpectation {
	if mmSendNotificationsToClients.mock.funcSendNotificationsToClients != nil {
		mmSendNotificationsToClients.mock.t.Fatalf("NotificationServiceMock.SendNotificationsToClients mock is already set by Set")
	}

	expectation := &NotificationServiceMockSendNotificationsToClientsExpectation{
		mock:               mmSendNotificationsToClients.mock,
		params:             &NotificationServiceMockSendNotificationsToClientsParams{ctx},
		expectationOrigins: NotificationServiceMockSendNotificationsToClientsExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSendNotificationsToClients.expectations = append(mmSendNotificationsToClients.expectations, expectation)
	return expectation
}

// Then sets up NotificationService.SendNotificationsToClients return parameters for the expectation previously defined by the When method
func (e *NotificationServiceMockSendNotificationsToClientsExpectation) Then(err error) *NotificationServiceMock {
	e.results = &NotificationServiceMockSendNotificationsToClientsResults{err}
	return e.mock
}

// Times sets number of times NotificationService.SendNotificationsToClients should be invoked
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Times(n uint64) *mNotificationServiceMockSendNotificationsToClients {
	if n == 0 {
		mmSendNotificationsToClients.mock.t.Fatalf("Times of NotificationServiceMock.SendNotificationsToClients mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendNotificationsToClients.expectedInvocations, n)
	mmSendNotificationsToClients.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSendNotificationsToClients
}

func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) invocationsDone() bool {
	if len(mmSendNotificationsToClients.expectations) == 0 && mmSendNotificationsToClients.defaultExpectation == nil && mmSendNotificationsToClients.mock.funcSendNotificationsToClients == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendNotificationsToClients.mock.afterSendNotificationsToClientsCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendNotificationsToClients.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendNotificationsToClients implements mm_service.NotificationService
func (mmSendNotificationsToClients *NotificationServiceMock) SendNotificationsToClients(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmSendNotificationsToClients.beforeSendNotificationsToClientsCounter, 1)
	defer mm_atomic.AddUint64(&mmSendNotificationsToClients.afterSendNotificationsToClientsCounter, 1)

	mmSendNotificationsToClients.t.Helper()

	if mmSendNotificationsToClients.inspectFuncSendNotificationsToClients != nil {
		mmSendNotificationsToClients.inspectFuncSendNotificationsToClients(ctx)
	}

	mm_params := NotificationServiceMockSendNotificationsToClientsParams{ctx}

	// Record call args
	mmSendNotificationsToClients.SendNotificationsToClientsMock.mutex.Lock()
	mmSendNotificationsToClients.SendNotificationsToClientsMock.callArgs = append(mmSendNotificationsToClients.SendNotificationsToClientsMock.callArgs, &mm_params)
	mmSendNotificationsToClients.SendNotificationsToClientsMock.mutex.Unlock()

	for _, e := range mmSendNotificationsToClients.SendNotificationsToClientsMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.Counter, 1)
		mm_want := mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.params
		mm_want_ptrs := mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.paramPtrs

		mm_got := NotificationServiceMockSendNotificationsToClientsParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSendNotificationsToClients.t.Errorf("NotificationServiceMock.SendNotificationsToClients got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendNotificationsToClients.t.Errorf("NotificationServiceMock.SendNotificationsToClients got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendNotificationsToClients.SendNotificationsToClientsMock.defaultExpectation.results
		if mm_results == nil {
			mmSendNotificationsToClients.t.Fatal("No results are set for the NotificationServiceMock.SendNotificationsToClients")
		}
		return (*mm_results).err
	}
	if mmSendNotificationsToClients.funcSendNotificationsToClients != nil {
		return mmSendNotificationsToClients.funcSendNotificationsToClients(ctx)
	}
	mmSendNotificationsToClients.t.Fatalf("Unexpected call to NotificationServiceMock.SendNotificationsToClients. %v", ctx)
	return
}

// SendNotificationsToClientsAfterCounter returns a count of finished NotificationServiceMock.SendNotificationsToClients invocations
func (mmSendNotificationsToClients *NotificationServiceMock) SendNotificationsToClientsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendNotificationsToClients.afterSendNotificationsToClientsCounter)
}

// SendNotificationsToClientsBeforeCounter returns a count of NotificationServiceMock.SendNotificationsToClients invocations
func (mmSendNotificationsToClients *NotificationServiceMock) SendNotificationsToClientsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendNotificationsToClients.beforeSendNotificationsToClientsCounter)
}

// Calls returns a list of arguments used in each call to NotificationServiceMock.SendNotificationsToClients.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendNotificationsToClients *mNotificationServiceMockSendNotificationsToClients) Calls() []*NotificationServiceMockSendNotificationsToClientsParams {
	mmSendNotificationsToClients.mutex.RLock()

	argCopy := make([]*NotificationServiceMockSendNotificationsToClientsParams, len(mmSendNotificationsToClients.callArgs))
	copy(argCopy, mmSendNotificationsToClients.callArgs)

	mmSendNotificationsToClients.mutex.RUnlock()

	return argCopy
}

// MinimockSendNotificationsToClientsDone returns true if the count of the SendNotificationsToClients invocations corresponds
// the number of defined expectations
func (m *NotificationServiceMock) MinimockSendNotificationsToClientsDone() bool {
	if m.SendNotificationsToClientsMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendNotificationsToClientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendNotificationsToClientsMock.invocationsDone()
}

// MinimockSendNotificationsToClientsInspect logs each unmet expectation
func (m *NotificationServiceMock) MinimockSendNotificationsToClientsInspect() {
	for _, e := range m.SendNotificationsToClientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NotificationServiceMock.SendNotificationsToClients at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSendNotificationsToClientsCounter := mm_atomic.LoadUint64(&m.afterSendNotificationsToClientsCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendNotificationsToClientsMock.defaultExpectation != nil && afterSendNotificationsToClientsCounter < 1 {
		if m.SendNotificationsToClientsMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to NotificationServiceMock.SendNotificationsToClients at\n%s", m.SendNotificationsToClientsMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to NotificationServiceMock.SendNotificationsToClients at\n%s with params: %#v", m.SendNotificationsToClientsMock.defaultExpectation.expectationOrigins.origin, *m.SendNotificationsToClientsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendNotificationsToClients != nil && afterSendNotificationsToClientsCounter < 1 {
		m.t.Errorf("Expected call to NotificationServiceMock.SendNotificationsToClients at\n%s", m.funcSendNotificationsToClientsOrigin)
	}

	if !m.SendNotificationsToClientsMock.invocationsDone() && afterSendNotificationsToClientsCounter > 0 {
		m.t.Errorf("Expected %d calls to NotificationServiceMock.SendNotificationsToClients at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SendNotificationsToClientsMock.expectedInvocations), m.SendNotificationsToClientsMock.expectedInvocationsOrigin, afterSendNotificationsToClientsCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NotificationServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockSendNotificationsToClientsInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NotificationServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *NotificationServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockSendNotificationsToClientsDone()
}
