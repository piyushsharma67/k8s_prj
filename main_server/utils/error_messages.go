package utils

import "errors"

var USER_ALREADY_EXISTS = errors.New("user already exists in DB")
var INTERNAL_SERVER_ERROR = errors.New("Internal server error")
var ERROR_FETCHING_RECORDS = errors.New("Error Fetching records!!")
var DB_INSTANCE_REQUIRED = errors.New("DB instance required!!")