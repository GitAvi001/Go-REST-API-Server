package handler

import "net/http"

func postNews() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getAllNews() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetNewsByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotImplemented)
	}
}

func UpdateNewsByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotImplemented)
	}
}

func DeleteNewsByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotImplemented)
	}
}
