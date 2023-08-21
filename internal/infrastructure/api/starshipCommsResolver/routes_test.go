package starshipcommsresolver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes2(t *testing.T) {
	// Crear un enrutador de prueba
	router := gin.Default()

	// Registrar las rutas utilizando la función RegisterRoutes2
	RegisterRoutes2(router)

	// Crear una solicitud de prueba para la ruta "/topsecret/"
	reqBody := []byte(`{"key": "value"}`) // Cambia esto al cuerpo que necesitas
	req, err := http.NewRequest("POST", "/topsecret/", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Error al crear la solicitud de prueba: %v", err)
	}

	// Crear un registrador de respuesta para capturar la salida
	respRecorder := httptest.NewRecorder()

	// Enviar la solicitud al enrutador
	router.ServeHTTP(respRecorder, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusNotFound, respRecorder.Code, "Código de estado esperado: 404 Not Found")

	// Verificar que la ruta "/topsecret/" esté registrada correctamente en el enrutador
	routes := router.Routes()
	assert.True(t, routeExists(routes, "POST", "/topsecret/"), "La ruta '/topsecret/' debería estar registrada en el enrutador")
}

func routeExists(routes gin.RoutesInfo, method, path string) bool {
	for _, route := range routes {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}
