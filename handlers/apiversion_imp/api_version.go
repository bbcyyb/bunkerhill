package apiversion_imp

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/restapi/operations/apiversion"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func GetAPIVersion(params apiversion.GetAPIVersionParams) middleware.Responder {
	err_payload := &models.GenericError{
		Message: "api version file not found or could not be read",
	}

	log.Println(params)

	if file, err := os.Open("version/version"); err == nil {
		if data, err := ioutil.ReadAll(file); err == nil {
			stdout := strings.Fields(string(data))
			payload := &models.GetAPIVersionOKBody{
				Data: &models.Apiversion{
					APIVersion: swag.String(stdout[0]),
				},
			}
			defer file.Close()

			return apiversion.NewGetAPIVersionOK().WithPayload(payload)
		}
	}

	return apiversion.NewGetAPIVersionInternalServerError().WithPayload(err_payload)
}
