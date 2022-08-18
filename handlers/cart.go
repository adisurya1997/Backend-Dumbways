package handlers

import (
	cartdto "waysbuck/dto/cart"
	dto "waysbuck/dto/result"
	"waysbuck/models"
	"waysbuck/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var toppingsId []int
	for _, r := range r.FormValue("toppingId") {
		if int(r-'0') >= 0 {
			toppingsId = append(toppingsId, int(r-'0'))
		}
    }


	subamount, _ := strconv.Atoi(r.FormValue("subamount"))
	productId, _ := strconv.Atoi(r.FormValue("productId"))
	transactionId, _ := strconv.Atoi(r.FormValue("transactionId"))

	request := cartdto.CartRequest{
		SubAmount:     subamount,    
		ProductID:     productId,
		ToppingID:	   toppingsId,
		TransactionID:     transactionId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get all category data by id [] 
	topping, _ := h.CartRepository.FindToppingsById(toppingsId)

	cart := models.Cart{
		SubAmount:  request.SubAmount,
		ProductID:  productId,
		Topping:	topping,
	}

	cart, err = h.CartRepository.CreateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	cart, _ = h.CartRepository.GetCart(cart.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}