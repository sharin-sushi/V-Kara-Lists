package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/sharin-sushi/0016go_next_relation/interfaces/controllers"
)

// 命名規則
// https://github.com/sharin-sushi/0016go_next_relation/issues/71#issuecomment-1843543763

func Routing(r *gin.Engine) {
	Controller := controllers.NewController(dbInit())

	ver := r.Group("/v1")
	{
		users := ver.Group("/users")
		{
			users.POST("/signup", Controller.CreateUser)
			users.PUT("/login", Controller.LogIn)
			users.PUT("/logout", controllers.Logout)
			users.DELETE("/withdraw", Controller.LogicalDeleteUser)
			users.GET("/gestlogin", controllers.GuestLogIn)
			users.GET("/profile", Controller.GetListenerProfile)
			users.GET("/mypage", Controller.ListenerPage)
		}
		vcontents := ver.Group("/vcontents")
		{
			vcontents.GET("/", Controller.ReturnTopPageData)
			vcontents.GET("/vtuber/:kana", Controller.ReturnVtuberPageData)
			vcontents.GET("/sings", Controller.GetJoinVtubersMoviesKaraokes)
			vcontents.GET("/original-song", Controller.ReturnOriginalSongPage)
			// /vtuber, /movie, /karaokeの文字列はフロント側で比較演算に使われてる
			// データ新規登録
			vcontents.POST("/create/vtuber", Controller.CreateVtuber)
			vcontents.POST("/create/movie", Controller.CreateMovie)
			vcontents.POST("/create/karaoke", Controller.CreateKaraoke)

			//データ編集
			vcontents.POST("/edit/vtuber", Controller.EditVtuber)
			vcontents.POST("/edit/movie", Controller.EditMovie)
			vcontents.POST("/edit/karaoke", Controller.EditKaraoke)

			// // データ削除(物理)
			vcontents.GET("/delete/deletePage", Controller.DeleteOfPage)
			vcontents.DELETE("/delete/vtuber", Controller.DeleteVtuber)
			vcontents.DELETE("/delete/movie", Controller.DeleteMovie)
			vcontents.DELETE("/delete/karaoke", Controller.DeleteKaraoke)

			//ドロップダウン用
			vcontents.GET("/getalldata", Controller.GetVtuverMovieKaraoke)

			// テスト用
			vcontents.GET("/dummy-top-page", Controller.ReturnDummyTopPage)
		}
		fav := ver.Group("/fav")
		{
			fav.POST("/favorite/movie", Controller.SaveMovieFavorite)
			fav.DELETE("/unfavorite/movie", Controller.DeleteMovieFavorite)
			fav.POST("/favorite/karaoke", Controller.SaveKaraokeFavorite)
			fav.DELETE("/unfavorite/karaoke", Controller.DeleteKaraokeFavorite)
		}
	}
}

// //開発者用　パスワード照会（ リポジトリ0019で作り直した）
// r.GET("/envpass", postrequest.EnvPass)
