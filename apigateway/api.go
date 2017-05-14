package apigateway

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

type APIGateway struct {
	sess *session.Session
	svc  *apigateway.APIGateway
}

func NewAPIGateway() *APIGateway {
	apiGw := &APIGateway{}
	apiGw.sess = session.Must(session.NewSession())
	apiGw.svc = apigateway.New(apiGw.sess)
	return apiGw
}

func emptyIfNil(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (apiGw *APIGateway) ListAPIs() ([]API, error) {
	params := &apigateway.GetRestApisInput{
		Limit: aws.Int64(500),
	}
	//FIXME(samalba) should support pagination for a list of APIs >500
	resp, err := apiGw.svc.GetRestApis(params)
	if err != nil {
		return nil, err
	}
	apis := make([]API, len(resp.Items))
	for i, item := range resp.Items {
		api := API{
			Id:          emptyIfNil(item.Id),
			Name:        emptyIfNil(item.Name),
			Version:     emptyIfNil(item.Version),
			CreatedDate: *item.CreatedDate,
			Description: emptyIfNil(item.Description),
			//BinaryMediaTypes: *item.BinaryMediaTypes,
		}
		apis[i] = api
	}
	return apis, nil
}

func (apiGw *APIGateway) GetResources(apiId string) ([]Resource, error) {
	//FIXME(samalba) should support pagination for a list of resources >500
	params := &apigateway.GetResourcesInput{
		RestApiId: aws.String(apiId),
		Limit:     aws.Int64(500),
	}
	resp, err := apiGw.svc.GetResources(params)
	if err != nil {
		return nil, err
	}
	ress := make([]Resource, len(resp.Items))
	for i, item := range resp.Items {
		res := Resource{
			Id:       emptyIfNil(item.Id),
			ParentId: emptyIfNil(item.ParentId),
			Path:     emptyIfNil(item.Path),
			PathPart: emptyIfNil(item.PathPart),
		}
		ress[i] = res
	}
	return ress, nil
}
