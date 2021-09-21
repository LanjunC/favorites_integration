package model

const (
	GetPinboxFavorsApi          = "https://withpinbox.com/api/user/695055d019/collection"                  //?count=30&category=all&order=create&sort=desc
	GetPinboxFavorItemsByCidApi = "https://withpinbox.com/api/user/695055d019/collection/%v/item?count=%v" //category=all&order=create&sort=desc
)

type PinboxFavors []struct {
	ID         int    `json:"id"` //0是默认收藏夹，弃用
	Name       string `json:"name"`
	Color      string `json:"color"`
	EditedAt   string `json:"edited_at"`
	CreatedAt  string `json:"created_at"`
	ItemsCount int    `json:"items_count"`
}

type PinboxFavorItems struct {
	Items []struct {
		ID           int      `json:"id"`
		CollectionID int      `json:"collection_id"`
		ItemType     string   `json:"item_type"`
		Note         string   `json:"note"`
		Tags         []string `json:"tags"`
		URL          string   `json:"url"`
		Title        string   `json:"title"`
		Thumbnail    string   `json:"thumbnail"`
		CreatedAt    string   `json:"created_at"`
	} `json:"items"`
	ItemsCount int `json:"items_count"`
}
