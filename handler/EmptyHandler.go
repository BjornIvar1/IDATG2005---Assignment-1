package handler

import (
	"IDATG2005---Assignment-1/utils"
	"fmt"
	"net/http"
)

// EmptyHandler
// handles requests to the root path and provides information about available endpoints.
func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. Please use the paths <a href=\"" +
		utils.StatusPath + "\">" + "Status" + "</a>, " +
		"<a href=\"" + utils.ExchangePath + "\">" + "Exchange" + "</a>, " +
		"or <a href=\"" + utils.InfoPath + "\">" + "Info" + "</a>."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}
