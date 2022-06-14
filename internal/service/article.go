package service

type CountAticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type AticleListRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTAticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateAticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteAticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
