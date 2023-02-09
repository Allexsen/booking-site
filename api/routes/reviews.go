package routes

import (
	"github.com/Allexsen/booking-site/pkg/controllers"
)

func initReviews() {
	r.GET("/reviews", controllers.LoadReviews)

	r.GET("/reviews/add/:token", controllers.IsRewieavableToken, controllers.LoadNewReviewPage)

	r.POST("/reviews/add/:token", controllers.DisableToken, controllers.AddReview)
}
