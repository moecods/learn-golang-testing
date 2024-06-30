package tutorialgolangtesting

import "github.com/stretchr/testify/mock"


type Post struct {
	ID int
	Title string
}

type PostRepository interface {
	GetPost(id int) (*Post, error)
}

// ----------- Mock Post Repository -------------
type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) GetPost(id int) (*Post, error) {
    args := m.Called(id)
    if post := args.Get(0); post != nil {
        return post.(*Post), args.Error(1)
    }
    return nil, args.Error(1)
}

// ---------- Mongo Post Repository -----------
type MongoPostRepository struct {}

// func (m *MongoPostRepository) GetPost(id int) (*Post, error) {
// 	// ...
// }

// ------------- Post Service -----------------
type PostService struct {
	repo PostRepository
}

func (s *PostService) GetPostTitle(id int) (string, error) {
	post, err := s.repo.GetPost(id)

	if err != nil {
		return "", err
	}

	return post.Title, nil
}
