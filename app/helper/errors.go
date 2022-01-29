package helper

import "net/http"

const (
	InvalidJsonCode           = 1
	GetSpacecraftErrorCode    = 2
	CreateSpacecraftErrorCode = 3
	UpdateSpacecraftErrorCode = 4
	DeleteSpacecraftErrorCode = 5
	SomethingWentWrongCode    = 6
)

const (
	InvalidJson           = "Cannot parse request body."
	GetSpacecraftError    = "Cannot get spacecrafts."
	CreateSpacecraftError = "Unable to save spacecrafts."
	UpdateSpacecraftError = "Unable to update spacecrafts."
	DeleteSpacecraftError = "Unable to delete spacecrafts."
	SomethingWentWrong    = "Something went wrong."
)

var errorsMap map[int]string

func init() {
	errorsMap = make(map[int]string)
	errorsMap[InvalidJsonCode] = InvalidJson
	errorsMap[GetSpacecraftErrorCode] = GetSpacecraftError
	errorsMap[CreateSpacecraftErrorCode] = CreateSpacecraftError
	errorsMap[UpdateSpacecraftErrorCode] = UpdateSpacecraftError
	errorsMap[DeleteSpacecraftErrorCode] = DeleteSpacecraftError
	errorsMap[SomethingWentWrongCode] = SomethingWentWrong
}

func CheckErr(w http.ResponseWriter, r *http.Request, code, status int, e error) bool {
	if e != nil {
		SendResponse(w, r, code, status, e, "")
		return true
	}
	return false
}
