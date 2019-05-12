package payments

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/gorilla/mux"
)

var dao = DAO{}
var config = Config{}

func init() {
	config.LoadConfig()
	dao.Host = config.Host
	dao.Database = config.Database
	dao.Collection = config.Collection
	dao.Connect()
	log.Println("Successfully initialized MongoDB connectivity")
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/payments", ListPayments).Methods("GET")
	router.HandleFunc("/payments/{id}", GetPayment).Methods("GET")
	router.HandleFunc("/payments", AddPayment).Methods("POST")
	router.HandleFunc("/payments/{id}", UpdatePayment).Methods("PUT")
	router.HandleFunc("/payments/{id}", DeletePayment).Methods("DELETE")
	address := fmt.Sprintf("localhost:%d", config.APIPort)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Error initializing API: %v", err)
	}
}

func ListPayments(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside ListPayments")
	payments, err := dao.FindAll()
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	JSONResponse(w, http.StatusOK, payments)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payment, err := dao.FindById(params["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Payment not found",
		})
		return
	}
	JSONResponse(w, http.StatusOK, payment)
}

func AddPayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		JSONResponse(w, http.StatusBadRequest, "Unable to parse request")
		return
	}
	payment.ID = bson.NewObjectId()
	if err := dao.Insert(payment); err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	defer r.Body.Close()
	payment, err := dao.FindById(params["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Payment not found",
		})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}
	if err := dao.Update(payment); err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payment, err := dao.FindById(params["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Payment not found",
		})
		return
	}
	if err := dao.Delete(payment); err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}

func JSONResponse(w http.ResponseWriter, returnCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(returnCode)
	w.Write(response)
}
