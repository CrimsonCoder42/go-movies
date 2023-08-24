package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JSONResponse defines the structure for the JSON responses this application will use.

type JSONResponse struct {
	Error   bool        `json:"error"`          // Indicator of whether there was an error or not.
	Message string      `json:"message"`        // A message, usually describing the error or success.
	Data    interface{} `json:"data,omitempty"` // Actual data being returned. "omitempty" means the key is omitted from the object if its value is empty.
}

// writeJSON is a method for the application struct to write JSON data to an HTTP response.

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	// Convert the provided data into JSON format.

	out, err := json.Marshal(data)
	if err != nil {
		return err // Return any encountered error.
	}

	// If there are additional headers provided, set them for the response.
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// Set the "Content-Type" header to "application/json".
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status for the response.
	w.WriteHeader(status)

	// Write the JSON data to the response.
	_, err = w.Write(out)
	if err != nil {
		return err // Return any encountered error.
	}

	// Return nil if everything went smoothly.
	return nil
}


// Define a method for the application struct to read JSON data from an HTTP request.
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	
	// Define the maximum bytes to be read. This is set to one megabyte.
	maxBytes := 1024 * 1024

	// Wrap the request body reader with a reader that errors if the size exceeds maxBytes.
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Create a new JSON decoder for the request body.
	dec := json.NewDecoder(r.Body)

	// Disallow unknown fields in the JSON input to ensure input is as expected.
	dec.DisallowUnknownFields()

	// Decode the request body JSON into the provided data variable.
	err := dec.Decode(data)
	if err != nil {
		return err // Return any encountered error.
	}

	// Check for any additional JSON values after the first one and error if found.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		// Return an error if there's more than one JSON value.
		return errors.New("body must only have a single JSON value")
	}

	// Return nil if everything went smoothly.
	return nil
}


func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
	
}