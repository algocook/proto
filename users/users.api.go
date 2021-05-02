package users

import (
	context "context"

	grpc "google.golang.org/grpc"
)

// Client comment
type Client struct {
	conn *grpc.ClientConn
}

// NewClient init
func NewClient() (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	con, err := grpc.Dial("users:5300", opts...)
	return &Client{con}, err
}

// GetUser function
func (client *Client) GetUser(id int64) *GetUserResponse {
	cli := NewUsersClient(client.conn)
	request := GetUserRequest{
		Id: id,
	}
	response, err := cli.GetUser(context.Background(), &request)

	if err != nil {
		return &GetUserResponse{
			Error: err.Error(),
		}
	}

	return response
}

// PostUser function
func (client *Client) PostUser(username, title, description string) *PostUserResponse {
	cli := NewUsersClient(client.conn)
	request := PostUserRequest{
		Username:    username,
		Title:       title,
		Description: description,
	}

	response, err := cli.PostUser(context.Background(), &request)

	if err != nil {
		return &PostUserResponse{
			Error: err.Error(),
		}
	}

	return response
}

// CheckUsername method
func (client *Client) CheckUsername(username string) *CheckUsernameResponse {
	cli := NewUsersClient(client.conn)
	request := CheckUsernameRequest{
		Username: username,
	}

	response, err := cli.CheckUsername(context.Background(), &request)

	if err != nil {
		return &CheckUsernameResponse{
			Error: err.Error(),
		}
	}

	return response
}
