package handlers

import (
	"mirae-va/app/helpers"
	"mirae-va/app/resources"
	"mirae-va/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta MetaData               `json:"meta"`
	Data map[string]interface{} `json:"data"`
}

type MetaData struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type VirtualAccountTransactionHandler struct {
	vaService *services.VirtualAccountTransactionService
}

func NewVirtualAccountTransactionHandler() *VirtualAccountTransactionHandler {
	return &VirtualAccountTransactionHandler{
		vaService: services.NewVirtualAccountTransactionService(),
	}
}

func (h *VirtualAccountTransactionHandler) BillGenerate(c *gin.Context) {
	var dataResource resources.BillGenerate
	var errMsg string
	var responseData Response

	if err := c.ShouldBindJSON(&dataResource); err != nil {

		switch err.(type) {

		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "CustomerID" && fieldError.Tag() == "required" {
					errMsg = "Customer ID is required"
				} else if fieldError.Field() == "BillAmount" && fieldError.Tag() == "required" {
					errMsg = "Bill amount is required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = err.Error()

		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)
		return

	}

	result, err := h.vaService.BillGenerate(dataResource)

	if err != nil {
		responseData = Response{
			Meta: MetaData{
				Status:  false,
				Code:    422,
				Message: err.Error(),
			},
			Data: nil,
		}
		c.JSON(http.StatusOK, responseData)
		return

	}

	responseData = Response{
		Meta: MetaData{
			Status:  true,
			Code:    200,
			Message: "Bill successfully created",
		},
		Data: result,
	}
	c.JSON(http.StatusOK, responseData)
	return
}

func (h *VirtualAccountTransactionHandler) Inquiry(c *gin.Context) {
	var dataResource resources.Inquiry
	var errMsg string
	var responseData Response

	if err := c.ShouldBindJSON(&dataResource); err != nil {

		switch err.(type) {

		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "VirtualAccountNumber" && fieldError.Tag() == "required" {
					errMsg = "Virtual Account Number ID is required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = err.Error()

		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)
		return

	}

	result, err := h.vaService.Inquiry(dataResource)

	if err != nil {
		responseData = Response{
			Meta: MetaData{
				Status:  false,
				Code:    422,
				Message: err.Error(),
			},
			Data: nil,
		}
		c.JSON(http.StatusOK, responseData)
		return

	}

	responseData = Response{
		Meta: MetaData{
			Status:  true,
			Code:    200,
			Message: "Inquiry success",
		},
		Data: result,
	}
	c.JSON(http.StatusOK, responseData)
	return
}

func (h *VirtualAccountTransactionHandler) Payment(c *gin.Context) {
	var dataResource resources.Payment
	var errMsg string
	var responseData Response

	if err := c.ShouldBindJSON(&dataResource); err != nil {

		switch err.(type) {

		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "VirtualAccountNumber" && fieldError.Tag() == "required" {
					errMsg = "Virtual Account Number ID is required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = err.Error()

		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)
		return

	}

	result, err := h.vaService.Payment(dataResource)

	if err != nil {
		responseData = Response{
			Meta: MetaData{
				Status:  false,
				Code:    422,
				Message: err.Error(),
			},
			Data: nil,
		}
		c.JSON(http.StatusOK, responseData)
		return

	}

	responseData = Response{
		Meta: MetaData{
			Status:  true,
			Code:    200,
			Message: "Payment success",
		},
		Data: result,
	}
	c.JSON(http.StatusOK, responseData)
	return
}

func (h *VirtualAccountTransactionHandler) Report(c *gin.Context) {
	var dataResource resources.Report
	var errMsg string
	var responseData Response

	if err := c.ShouldBindJSON(&dataResource); err != nil {

		switch err.(type) {

		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "VirtualAccountNumber" && fieldError.Tag() == "required" {
					errMsg = "Virtual Account Number ID is required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = err.Error()

		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)
		return

	}

	result, err := h.vaService.Report(dataResource)

	if err != nil {
		responseData = Response{
			Meta: MetaData{
				Status:  false,
				Code:    422,
				Message: err.Error(),
			},
			Data: nil,
		}
		c.JSON(http.StatusOK, responseData)
		return

	}

	responseData = Response{
		Meta: MetaData{
			Status:  true,
			Code:    200,
			Message: "Get report",
		},
		Data: result,
	}
	c.JSON(http.StatusOK, responseData)
	return
}
