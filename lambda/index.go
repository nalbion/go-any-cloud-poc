package main

// /* Required, but no C code needed. */
import "C"

import (
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/nalbion/go-any-cloud-poc/lib/controllers"
	"github.com/nalbion/go-any-cloud-poc/lib/rest/eawsy"
)

func Handle(evt apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	return controllers.HandleIndex(eawsy.NewInboundRequest(&evt))
}
