package custom_errs
import "github.com/pkg/errors"
var DBErrCreateWithID = errors.New("create with object id not allowed")
var DBErrGetWithID = errors.New("get with object id not allowed")
var DBErrUpdateWithID = errors.New("update with object id not allowed")
var DBErrDeleteWithID = errors.New("delete with object id not allowed")
var DBErrIDConversion = errors.New("id conversion error")
var ServerError = errors.New("server error")
var ParamError = errors.New("param error")
var InvalidInput = errors.New("invalid input")
var InvalidTag = errors.New("invalid tag")
var DbErrors = errors.New("database errors")
var DecodeErr = errors.New("can not decode")