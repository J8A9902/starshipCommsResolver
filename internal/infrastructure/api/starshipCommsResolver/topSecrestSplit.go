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

type topsecretSplitHandler struct {
	topsecretsplitService ports.CommunicationServices
	topsecretsplitUseCase *usecase.TopsecretSplitUseCase
}

func newSplitHandler(service ports.CommunicationServices) *topsecretSplitHandler {
	orderUseCase := usecase.NewTopsecretSplitUseCase(service)
	return &topsecretSplitHandler{
		topsecretsplitService: service,
		topsecretsplitUseCase: orderUseCase,
	}
}

func (o *topsecretSplitHandler) topsecretSplit(c *gin.Context) {

	var SatelliteData entity.SatelliteData
	satelliteName := c.Param("satellite_name")
	if err := c.BindJSON(&SatelliteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretsplitUseCase.CreateTopSecretSplit(&satelliteName, &SatelliteData)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}

func (o *topsecretSplitHandler) GetsecretSplit(c *gin.Context) {

	var Satellite entity.Satellitequery
	satelliteName := c.Param("satellite_name")
	if err := c.BindQuery(&Satellite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretsplitUseCase.GetAll(satelliteName, &Satellite)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}

func (o *topsecretSplitHandler) PatchtopsecretSplit(c *gin.Context) {

	var SatelliteData entity.SatelliteData
	satelliteName := c.Param("satellite_name")
	if err := c.BindJSON(&SatelliteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretsplitUseCase.PatchtopsecretSplit(satelliteName, &SatelliteData)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}
