package reviews

import (
	"encoding/json"
	"net/http"

	"github.com/data_lodge/go_reviews/cassandra"
	"github.com/gorilla/mux"
)

// Get returns all the reviews for a given home
func Get(w http.ResponseWriter, r *http.Request) {
	var reviewList []Review
	vars := mux.Vars(r)
	homeID := vars["id"]
	m := map[string]interface{}{}
	iterable := cassandra.Session.Query("SELECT * FROM reviews WHERE home_id = ?", homeID).Iter()
	for iterable.MapScan(m) {
		reviewList = append(reviewList, Review{
			ID:            m["id"].(string),
			HomeID:        m["home_id"].(int),
			Accuracy:      m["accuracy"].(int),
			Communication: m["communication"].(int),
			Cleanliness:   m["cleanliness"].(int),
			CheckIn:       m["check_in"].(int),
			Value:         m["value"].(int),
			Location:      m["location"].(string),
			CreatedAt:     m["created_at"].(string),
			Name:          m["name"].(string),
			ImgURL:        m["img_url"].(string),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllReviews{Reviews: reviewList}.Reviews)
}
