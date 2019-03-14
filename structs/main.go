package reviews

// Review struct to hold review data on a home
type Review struct {
	ID            string `json:"id"`
	HomeID        int    `json:"home_id"`
	Accuracy      int    `json:"accuracy"`
	Communication int    `json:"communication"`
	Cleanliness   int    `json:"cleanliness"`
	CheckIn       int    `json:"check_in"`
	Value         int    `json:"value"`
	Location      int    `json:"location"`
	Complaints    bool   `json:"complaints"`
	CreatedAt     string `json:"created_at"`
	Name          string `json:"name"`
	ImgURL        string `json:"img_url"`
}

// getReview to form payload returning a Review struct
type getReview struct {
	Review Review `json:"review"`
}

// AllReviews to form payload of an array of Reviews structs
type AllReviews struct {
	Reviews []Review `json:"reviews"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
