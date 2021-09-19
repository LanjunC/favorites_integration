package model


const (
	GetFavorsApi = "https://api.zhihu.com/people/creater-60/collections"	//limit&offset非必须
	GetFavorItemsByCidApi = "https://api.zhihu.com/collections/%v/items"
)

type ZhihuFavors struct {
	Paging struct {
		IsEnd    bool   `json:"is_end"`
		IsStart  bool   `json:"is_start"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Totals   int    `json:"totals"`
	} `json:"paging"`
	Data []struct {
		ID            int    `json:"id"`
		Title         string `json:"title"`
		Description   string `json:"description"`
		ItemCount     int    `json:"item_count"`
		CreatedTime   int    `json:"created_time"`
		UpdatedTime   int    `json:"updated_time"`
	} `json:"data"`
}

type ZhihuFavorItems struct {
	Paging struct {
		IsEnd    bool   `json:"is_end"`
		IsStart  bool   `json:"is_start"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Totals   int    `json:"totals"`
	} `json:"paging"`
	Data []struct {
		Content struct {
			ID          int    `json:"id"`
			Type        string `json:"type"`
			URL         string `json:"url"`
			CreatedTime int    `json:"created_time"`
			UpdatedTime int    `json:"updated_time"`
			Question    struct {
				ID           int    `json:"id"`
				Title        string `json:"title"`
				Created      int    `json:"created"`
				UpdatedTime  int    `json:"updated_time"`
				URL          string `json:"url"`
			} `json:"question,omitempty"`
			Author struct {
				ID                string `json:"id"`
				URLToken          string `json:"url_token"`
				Name              string `json:"name"`
				UseDefaultAvatar  bool   `json:"use_default_avatar"`
				AvatarURL         string `json:"avatar_url"`
				AvatarURLTemplate string `json:"avatar_url_template"`
			} `json:"author"`
			VoteupCount        int    `json:"voteup_count"`
			CommentCount       int    `json:"comment_count"`
			ThanksCount        int    `json:"thanks_count"`
			//Content            string `json:"content"`
			/*
			  answer普通回答title为""，简略回答在excerpt
			  answer视频回答同普通回答
			  article文章title不为空，简略回答在excerpt_title
			  zvideo视频title不为空
			*/
			Title              string `json:"title"`
			ExcerptTitle string `json:"excerpt_title"`	//当type为article时才会有
			Excerpt            string `json:"excerpt"`	//当type为answer时
			Attachment struct {
				Type         string `json:"type"`
				AttachmentID string `json:"attachment_id"`
			} `json:"attachment"`
		} `json:"content,omitempty"`
	} `json:"data"`
}

