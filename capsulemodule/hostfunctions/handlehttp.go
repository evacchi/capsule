package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
    "github.com/bots-garden/capsule/commons"
)

/* previous version
var handleHttpFunction func(bodyReq string, headersReq map[string]string) (
    bodyResp string, headersResp map[string]string, errResp error)
*/

var handleHttpFunction func(req Request) (resp Response, errResp error)

func SetHandleHttp(function func(request Request) (Response, error)) {
    handleHttpFunction = function
}

// The name "callHandleHttp" of the exported function is defined/declared
// in `wasmrunner.go`, function: GetNewWasmRuntimeForHttp

//export callHandleHttp
//go:linkname callHandleHttp
func callHandleHttp(reqId uint32) (strPtrPosSize uint64) {

    reqParams, errReqParams := RequestParamsGet(reqId)
    if errReqParams != nil {
        // TODO
    }
    bodyParameter := reqParams[0]
    headersParameter := reqParams[1]
    uriParameter := reqParams[2]
    methodParameter := reqParams[3]

    headersSlice := commons.CreateSliceFromString(headersParameter, commons.StrHeadersSeparator)
    headers := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

    var result string
    responseReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(Request{bodyParameter, headers, uriParameter, methodParameter})

    returnHeaderString := commons.CreateStringFromSlice(commons.CreateSliceFromMap(responseReturnByHandleFunction.Headers), commons.StrSeparator)

    if errorReturnByHandleFunction != nil {
        result = commons.CreateStringError(errorReturnByHandleFunction.Error(), 0)
    } else {
        result = CreateBodyString(responseReturnByHandleFunction.Body)
    }

    // merge body and headers response
    pos, length := getStringPtrPositionAndSize(CreateResponseString(result, returnHeaderString))

    return packPtrPositionAndSize(pos, length)
}

func CreateBodyString(message string) string {
    return "[BODY]" + message
}

func CreateResponseString(result, headers string) string {
    return result + "[HEADERS]" + headers
}
