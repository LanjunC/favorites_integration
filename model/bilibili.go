package model

const (
	GetBiliFavors             = "https://api.bilibili.com/x/v3/fav/folder/created/list-all?up_mid=371742091"
	GetBiliFavorItemsByMidApi = "https://api.bilibili.com/x/v3/fav/resource/list?media_id=%v&pn=%v&ps=%v"
)

type BiliFavors struct {
	Code    int    `json:"code,omitempty"` //这里允许为空可以在收到bili api的code为0时，向前端传时忽略该字段
	Message string `json:"message,omitempty"`
	Data    struct {
		Count int `json:"count"`
		List  []struct {
			ID         int    `json:"id"`
			Fid        int    `json:"fid"`
			Mid        int    `json:"mid"`
			Title      string `json:"title"`
			MediaCount int    `json:"media_count"`
		} `json:"list"`
	} `json:"data"`
}

type BiliFavorItems struct {
	Code    int       `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
	Data    *struct { //这里用指针，若bili api传过来是null，说明对应mediaId不存在
		Info struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Cover string `json:"cover"`
			Upper struct {
				Mid  int    `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"upper"`
			CntInfo struct {
				Collect int `json:"collect"`
				Play    int `json:"play"`
				ThumbUp int `json:"thumb_up"`
				Share   int `json:"share"`
			} `json:"cnt_info"`
			Intro      string `json:"intro"`
			Ctime      int    `json:"ctime"`
			Mtime      int    `json:"mtime"`
			MediaCount int    `json:"media_count"`
		} `json:"info"`
		Medias []struct {
			ID       int    `json:"id"`
			Type     int    `json:"type"`
			Title    string `json:"title"`
			Cover    string `json:"cover"`
			Intro    string `json:"intro"`
			Page     int    `json:"page"`
			Duration int    `json:"duration"`
			Upper    struct {
				Mid  int    `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"upper"`
			CntInfo struct {
				Collect int `json:"collect"`
				Play    int `json:"play"`
				Danmaku int `json:"danmaku"`
			} `json:"cnt_info"`
			Link     string `json:"link"`
			Ctime    int    `json:"ctime"`
			Pubtime  int    `json:"pubtime"`
			FavTime  int    `json:"fav_time"`
			BvID     string `json:"bv_id"`
			VisitUrl string `json:"visit_url"` //访问该视频的url
		} `json:"medias"`
		HasMore bool `json:"has_more"`
	} `json:"data"`
}
