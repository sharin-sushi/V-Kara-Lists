package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharin-sushi/0016go_next_relation/domain"
	"github.com/sharin-sushi/0016go_next_relation/interfaces/controllers/common"
)

func (controller *Controller) GetJoinVtubersMoviesKaraokes(c *gin.Context) {
	// transmitKaraoke, err := controller.VtuberContentInteractor.GetVtubersMoviesKaraokes()
	VtsMosKasWithFav, err := controller.FavoriteInteractor.GetVtubersMoviesKaraokesWithFavCnts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"resultStsのerror": err.Error()})
		return
	}
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	myFav, err := controller.FavoriteInteractor.FindFavoritesCreatedByListenerId(listenerId)
	fmt.Printf("myFav=%v\n", myFav)
	transmitKaraokes := common.AddIsFavToKaraokeWithFav(VtsMosKasWithFav, myFav)

	c.JSON(http.StatusOK, gin.H{
		"vtubers_movies_karaokes": transmitKaraokes,
	})
	return
}
func (controller *Controller) CreateVtuber(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var vtuber domain.Vtuber
	if err := c.ShouldBind(&vtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	vtuber.VtuberInputterId = listenerId

	if err := controller.VtuberContentInteractor.CreateVtuber(vtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invailed Registered the New Vtuber",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Registered the New Vtuber",
	})
	return
}

func (controller *Controller) CreateMovie(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var movie domain.Movie
	if err := c.ShouldBind(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	fmt.Printf("1-1:*値よな？ %v \n", movie.VtuberId) //4

	movie.MovieInputterId = listenerId
	if err := controller.VtuberContentInteractor.CreateMovie(movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invailed Registered the New Movie",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Registered the New Movie",
	})
	return
}
func (controller *Controller) CreateKaraoke(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var karaoke domain.Karaoke
	if err := c.ShouldBind(&karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	karaoke.KaraokeInputterId = listenerId
	if err := controller.VtuberContentInteractor.CreateKaraoke(karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invailed Registered the New Karaoke",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Registered the New Karaoke",
	})
	return
}
func (controller *Controller) EditVtuber(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var vtuber domain.Vtuber
	if err := c.ShouldBind(&vtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	vtuber.VtuberInputterId = listenerId
	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyVtuber(listenerId, vtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.UpdateVtuber(vtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Inputter can modify each data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Update",
	})
	return
}
func (controller *Controller) EditMovie(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var Movie domain.Movie
	if err := c.ShouldBind(&Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	Movie.MovieInputterId = listenerId
	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyMovie(listenerId, Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.UpdateMovie(Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Inputter can modify each data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Update",
	})
	return
}
func (controller *Controller) EditKaraoke(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var Karaoke domain.Karaoke
	if err := c.ShouldBind(&Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	Karaoke.KaraokeInputterId = listenerId
	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyKaraoke(listenerId, Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.UpdateKaraoke(Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Inputter can modify each data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Update",
	})
	return
}

func (controller *Controller) DeleteVtuber(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var selectedVtuber domain.Vtuber
	if err := c.ShouldBind(&selectedVtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	fmt.Print("Vtuber", selectedVtuber, "\n")
	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyVtuber(listenerId, selectedVtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.DeleteVtuber(selectedVtuber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only Inputter can modify each data",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Delete",
	})
	return
}

func (controller *Controller) DeleteMovie(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var Movie domain.Movie
	if err := c.ShouldBind(&Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyMovie(listenerId, Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.DeleteMovie(Movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only Inputter can modify each data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Delete",
	})
	return
}

func (controller *Controller) DeleteKaraoke(c *gin.Context) {
	listenerId, err := common.TakeListenerIdFromJWT(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching listener info",
		})
		return
	}
	var Karaoke domain.Karaoke
	if err := c.ShouldBind(&Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if isAuth, err := controller.VtuberContentInteractor.VerifyUserModifyKaraoke(listenerId, Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Auth Check is failed.(we could not Verify)",
		})
		return
	} else if isAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only The Inputter can modify each data",
		})
		return
	}
	if err := controller.VtuberContentInteractor.DeleteKaraoke(Karaoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only Inputter can modify each data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Delete",
	})
	return
}

func (controller *Controller) ReturnTopPageData(c *gin.Context) {
	var errs []error
	allVts, err := controller.VtuberContentInteractor.GetVtubers()
	if err != nil {
		errs = append(errs, err)
	}
	VtsMosWithFav, err := controller.FavoriteInteractor.GetVtubersMoviesWithFavCnts()
	if err != nil {
		fmt.Print("err:", err)
		errs = append(errs, err)
	}
	VtsMosKasWithFav, err := controller.FavoriteInteractor.GetVtubersMoviesKaraokesWithFavCnts()
	if err != nil {
		errs = append(errs, err)
	}
	fmt.Printf("VtsMosKasWithFav= \n%v\n", VtsMosKasWithFav)
	listenerId, err := common.TakeListenerIdFromJWT(c) //非ログイン時でも処理は続ける
	fmt.Printf("listenerId=%v\n", listenerId)
	if err != nil || listenerId == 0 {
		errs = append(errs, err)
		c.JSON(http.StatusOK, gin.H{
			"vtubers":                 allVts,
			"vtubers_movies":          VtsMosWithFav,
			"vtubers_movies_karaokes": VtsMosKasWithFav,
			"error":                   errs,
			"message":                 "dont you Loged in ?",
		})
		return
	}
	fmt.Printf("VtsMosWitFav= \n %+v\n", VtsMosWithFav)
	fmt.Printf("VtsMosKasWitFav= \n %+v\n", VtsMosKasWithFav)
	myFav, err := controller.FavoriteInteractor.FindFavoritesCreatedByListenerId(listenerId)
	fmt.Printf("myFav= \n %v\n", myFav)
	TransmitMovies := common.AddIsFavToMovieWithFav(VtsMosWithFav, myFav)
	TransmitKaraokes := common.AddIsFavToKaraokeWithFav(VtsMosKasWithFav, myFav)

	c.JSON(http.StatusOK, gin.H{
		"vtubers":                 allVts,
		"vtubers_movies":          TransmitMovies,
		"vtubers_movies_karaokes": TransmitKaraokes,
		"error":                   errs,
	})
	return
}

// dropdown用かな？
func (controller *Controller) GetVtuverMovieKaraoke(c *gin.Context) {
	var errs []error
	allVts, err := controller.VtuberContentInteractor.GetVtubers()
	if err != nil {
		errs = append(errs, err)
	}
	allMos, err := controller.VtuberContentInteractor.GetMovies()
	if err != nil {
		errs = append(errs, err)
	}
	allKas, err := controller.VtuberContentInteractor.GetKaraokes()
	if err != nil {
		errs = append(errs, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"vtubers":                 allVts,
		"vtubers_movies":          allMos,
		"vtubers_movies_karaokes": allKas,
		"error":                   errs,
	})
}

// 管理者用　全データ取得のイメージ、とにかくデータを網羅して確認したいとき
func (controller *Controller) Enigma(c *gin.Context) {
	var errs []error
	allVts, err := controller.VtuberContentInteractor.GetVtubers()
	if err != nil {
		errs = append(errs, err)
	}
	allMos, err := controller.VtuberContentInteractor.GetMovies()
	if err != nil {
		errs = append(errs, err)
	}
	allKas, err := controller.VtuberContentInteractor.GetKaraokes()
	if err != nil {
		errs = append(errs, err)
	}
	if err != nil {
		errs = append(errs, err)
	}
	allVtsMos, err := controller.VtuberContentInteractor.GetVtubersMovies()
	if err != nil {
		errs = append(errs, err)

	}

	allVtsMosKas, err := controller.VtuberContentInteractor.GetVtubersMoviesKaraokes()
	if err != nil {
		errs = append(errs, err)
	}
	all, err := controller.VtuberContentInteractor.GetVtubersMoviesKaraokes()
	if err != nil {
		errs = append(errs, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"vtubers":                 allVts,
		"movies":                  allMos,
		"karaokes":                allKas,
		"vtubers_movies":          allVtsMos,
		"vtubers_movies_karaokes": allVtsMosKas,
		"all":                     all,
		"error":                   errs,
	})
}
