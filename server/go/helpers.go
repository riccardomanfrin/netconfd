/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.3.0
 * Contact: support@athonet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/riccardomanfrin/netconfd/logger"
	"github.com/riccardomanfrin/netconfd/nc"
)

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}

//PostErrorResponse maps Post requests errors into HTTP status codes
func PostErrorResponse(err error, body interface{}) (ImplResponse, error) {
	if err != nil {
		logger.Log.Warning(err)
		switch err.(type) {
		case *nc.SemanticError:
			{
				return Response(http.StatusBadRequest, err), err
			}
		case *nc.ConflictError:
			{
				return Response(http.StatusConflict, err), err
			}
		default:
			{
				return Response(http.StatusInternalServerError, err), err
			}
		}
	}
	return Response(http.StatusCreated, body), nil
}

//GetErrorResponse maps Get requests errors into HTTP status codes
func GetErrorResponse(err error, body interface{}) (ImplResponse, error) {
	if err != nil {
		switch err.(type) {
		case *nc.NotFoundError:
			{
				return Response(http.StatusNotFound, err), err
			}
		case *nc.SemanticError:
			{
				return Response(http.StatusBadRequest, err), err
			}
		default:
			{
				return Response(http.StatusInternalServerError, err), err
			}
		}
	}
	return Response(http.StatusOK, body), nil
}

//PutErrorResponse maps Put requests errors into HTTP status codes
func PutErrorResponse(err error, body interface{}) (ImplResponse, error) {
	if err != nil {
		switch err.(type) {
		case *nc.SemanticError:
			{
				return Response(http.StatusBadRequest, err), err
			}
		default:
			{
				return Response(http.StatusInternalServerError, err), err
			}
		}
	}
	return Response(http.StatusOK, body), nil
}

//PatchErrorResponse maps Put requests errors into HTTP status codes
func PatchErrorResponse(err error, body interface{}) (ImplResponse, error) {
	if err != nil {
		switch err.(type) {
		case *nc.SemanticError:
			{
				return Response(http.StatusBadRequest, err), err
			}
		default:
			{
				return Response(http.StatusInternalServerError, err), err
			}
		}
	}
	return Response(http.StatusOK, body), nil
}

//DeleteErrorResponse maps Delete requests errors into HTTP status codes
func DeleteErrorResponse(err error, body interface{}) (ImplResponse, error) {
	if err != nil {
		switch err.(type) {
		case *nc.NotFoundError:
			{
				return Response(http.StatusNotFound, err), err
			}
		default:
			{
				return Response(http.StatusInternalServerError, err), err
			}
		}
	}
	return Response(http.StatusOK, body), nil
}
