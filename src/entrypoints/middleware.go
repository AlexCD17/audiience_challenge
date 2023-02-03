package entrypoints

import (
	"audiience_challenge/resources"
	"net"

	"encoding/json"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// verifyMiddleware verifies url params from request
func verifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		stateRegex := "^[A-Z]{2}$"
		validTypes := []string{"normal", "premium"}
		validStates := []string{"NY", "TX", "OH", "AZ", "CA"}

		log.Println("Verifying params...")

		state := request.URL.Query().Get("state")

		if state == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.MissingState)
			return

		}

		match, _ := regexp.MatchString(stateRegex, state)
		if !match {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.WrongState)
			return
		}

		if !slices.Contains(validStates, state) {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.UnsupportedState)
			return
		}

		estimationType := request.URL.Query().Get("type")

		if estimationType == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.MissingType)
			return
		}

		if !slices.Contains(validTypes, strings.ToLower(estimationType)) {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.WrongType)
			return
		}

		distance := request.URL.Query().Get("distance")

		if distance == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.MissingDistance)
			return
		}

		if _, err := strconv.ParseFloat(distance, 32); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.WrongDistance)
			return
		}

		amount := request.URL.Query().Get("base_amount")

		if amount == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.MissingBaseAmount)
			return
		}

		if _, err := strconv.ParseFloat(distance, 32); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(resources.WrongBaseAmount)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

// ipValidatorMiddleware checks for the ip-client header and validates value is either ipv4 or ipv6
func ipValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Looking for valid IP...")

		//Get IP from the ip-client header
		ip := request.Header.Get("ip-client")
		netIP := net.ParseIP(ip)
		if netIP == nil {
			writer.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(writer).Encode(resources.InvalidIP)
			return
		}

		next.ServeHTTP(writer, request)

	})

}
