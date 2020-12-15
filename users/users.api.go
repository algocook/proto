package users

import (
	context "context"

	pb "github.com/algocook/proto/users"
	grpc "google.golang.org/grpc"
)

// Client comment
type Client struct {
	conn *grpc.ClientConn
}

// NewClient init
func NewClient() (Client, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	con, err := grpc.Dial("users:5300", opts...)
	return Client{con}, err
}

// User struct
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

// GetUser function
func (client *Client) GetUser(id int64) User {
	cli := pb.NewUsersClient(client.conn)
	request := pb.GetUserRequest{
		Id: id,
	}
	response, err := cli.GetUser(context.Background(), &request)

	if err != nil {
		return User{
			Error: err.Error(),
		}
	}

	return User{
		ID:          response.Id,
		Username:    response.Username,
		Title:       response.Title,
		Description: response.Description,
	}
}

// PostUser function
func (client *Client) PostUser(username string, title string, description string) User {
	cli := pb.NewUsersClient(client.conn)
	request := pb.PostUserRequest{
		Username:    username,
		Title:       title,
		Description: description,
	}

	response, err := cli.PostUser(context.Background(), &request)

	if err != nil {
		return User{
			Error: err.Error(),
		}
	}

	return User{
		ID:          response.Id,
		Username:    username,
		Title:       title,
		Description: description,
	}
}

// IsAvailable struct
type IsAvailable struct {
	Username    string `json:"username"`
	IsAvailable bool   `json:"is_available"`
	Error       string `json:"error"`
}

// CheckUsername method
func (client *Client) CheckUsername(username string) IsAvailable {
	cli := pb.NewUsersClient(client.conn)
	request := pb.CheckUsernameRequest{
		Username: username,
	}

	response, err := cli.CheckUsername(context.Background(), &request)

	if err != nil {
		return IsAvailable{
			Error: err.Error(),
		}
	}

	return IsAvailable{
		Username:    username,
		IsAvailable: response.IsAvailable,
	}
}
