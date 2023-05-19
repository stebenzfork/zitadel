package metrics

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
)

//const (
//	ActiveSessionCounter            = "zitadel.active_session_counter"
//	ActiveSessionCounterDescription = "Active session counter"
//	SpoolerDivCounter               = "zitadel.spooler_div_milliseconds"
//	SpoolerDivCounterDescription    = "Spooler div from last successful run to now in milliseconds"
//	Database                        = "database"
//	ViewName                        = "view_name"
//)

type Usage interface {
	GetUsageExporter() http.Handler
	GetUsageProvider() metric.MeterProvider
	RegisterUsageCounter(name, description string) error
	AddUsageCount(ctx context.Context, name string, value int64, labels map[string]attribute.Value) error
	RegisterUsageUpDownSumObserver(name, description string, callbackFunc instrument.Int64Callback) error
	RegisterUsageValueObserver(name, description string, callbackFunc instrument.Int64Callback) error
}

var U Usage

func GetUsageExporter() http.Handler {
	if U == nil {
		return nil
	}
	return U.GetUsageExporter()
}

func GetUsageProvider() metric.MeterProvider {
	if U == nil {
		return nil
	}
	return U.GetUsageProvider()
}

func RegisterUsageCounter(name, description string) error {
	if U == nil {
		return nil
	}
	return U.RegisterUsageCounter(name, description)
}

func AddUsageCount(ctx context.Context, name string, value int64, labels map[string]attribute.Value) error {
	if U == nil {
		return nil
	}
	return U.AddUsageCount(ctx, name, value, labels)
}

func RegisterUsageUpDownSumObserver(name, description string, callbackFunc instrument.Int64Callback) error {
	if U == nil {
		return nil
	}
	return U.RegisterUsageUpDownSumObserver(name, description, callbackFunc)
}

func RegisterUsageValueObserver(name, description string, callbackFunc instrument.Int64Callback) error {
	if U == nil {
		return nil
	}
	return U.RegisterUsageValueObserver(name, description, callbackFunc)
}
