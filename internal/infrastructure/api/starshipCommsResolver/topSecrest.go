package starshipcommsresolver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"starshipCommsResolver/internal/pkg/entity"
	"starshipCommsResolver/internal/pkg/ports"
	entity2 "starshipCommsResolver/internal/pkg/response"
	"starshipCommsResolver/internal/pkg/usecase"
)

type topsecretHandler struct {
	topsecretService ports.CommunicationServices
	topsecretUseCase *usecase.TopsecrestUseCase
}

func newHandler(service ports.CommunicationServices) *topsecretHandler {
	orderUseCase := usecase.NewTopsecretUseCase(service)
	return &topsecretHandler{
		topsecretService: service,
		topsecretUseCase: orderUseCase,
	}
}

func (o *topsecretHandler) topsecret(c *gin.Context) {

	var satellite entity.SatelliteRequest
	if err := c.BindJSON(&satellite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretUseCase.CreateTopSecret(&satellite)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}
