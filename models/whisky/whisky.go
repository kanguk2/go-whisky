package whisky

type Whisky struct {
	ID           int64  `db:"id, primarykey, autoincrement" json:"id"`
	CategoryID   int64  `db:"category_id" json:"category_id"`
	CountryID    string `db:"country_id" json:"country_id"`
	DistilleryID string `db:"distillery_id" json:"distillery_id"`
	Strength     int64  `db:"strength" json:"strength"`
	Size         int64  `db:"size" json:"size"`
	CreatorID    int64  `db:"creator_id" json:"creator_id"`
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`
	CreatedAt    int64  `db:"created_at" json:"created_at"`
}
