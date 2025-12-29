package media

type MediaService struct {
	repo *Repository
}

func NewService(repo *Repository) *MediaService {
	return &MediaService{
		repo: repo,
	}
}
//file size for the images
const(
	MaxImageSize=8*1024*1024
)
//file types for the images 

func (s *MediaService) CreateMedia(media *Media) error {

	return s.repo.CreateMedia(media)
}

func (s *MediaService) GetMediaByWrap(wrapUUID string) ([]*Media, error) {
	return s.repo.GetMediaByWrap(wrapUUID)
}

func (s *MediaService) DeleteMedia(mediaUUID string) error {
	return s.repo.DeleteMedia(mediaUUID)
}
