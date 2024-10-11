// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/nidoqueen1/article-api/entity"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: ctx, article
func (_m *IService) CreateArticle(ctx context.Context, article *entity.Article) error {
	ret := _m.Called(ctx, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArticle provides a mock function with given fields: ctx, articleID
func (_m *IService) GetArticle(ctx context.Context, articleID uint) (*entity.Article, error) {
	ret := _m.Called(ctx, articleID)

	var r0 *entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*entity.Article, error)); ok {
		return rf(ctx, articleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Article); ok {
		r0 = rf(ctx, articleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticlesByTagAndDate provides a mock function with given fields: ctx, tagName, date
func (_m *IService) GetArticlesByTagAndDate(ctx context.Context, tagName string, date time.Time) ([]*entity.Article, int64, error) {
	ret := _m.Called(ctx, tagName, date)

	var r0 []*entity.Article
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) ([]*entity.Article, int64, error)); ok {
		return rf(ctx, tagName, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) []*entity.Article); ok {
		r0 = rf(ctx, tagName, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) int64); ok {
		r1 = rf(ctx, tagName, date)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, time.Time) error); ok {
		r2 = rf(ctx, tagName, date)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewIService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIService(t mockConstructorTestingTNewIService) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
