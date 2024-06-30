package tutorialgolangtesting

import "testing"

func TestUserService_GetUserName(t *testing.T) {
	mockRepo := &MockUserRepository{
		users: map[int]*User {
			1: {ID: 1, Name: "John Doe"},
		},
	}

	service := &UserService{repo: mockRepo}

	name, err := service.GetUserName(1)

	if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

	if name != "John Doe" {
        t.Errorf("expected name to be 'John Doe', got %v", name)
    }

	_, err = service.GetUserName(2)
    if err == nil {
        t.Fatalf("expected an error, got nil")
    }
}
