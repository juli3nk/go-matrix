package matrix

import (
	"github.com/matrix-org/gomatrix"
)

type MatrixConfig struct {
	client *gomatrix.Client
}

func New(url, userID, token string) (*MatrixConfig, error) {
	cli, err := gomatrix.NewClient(url, userID, token)
	if err != nil {
		return nil, err
	}

	config := MatrixConfig{
		client: cli,
	}

	return &config, nil
}

func (m *MatrixConfig) SendText(room, text string) error {
	resp, err := m.client.JoinRoom(room, "", nil)
	if err != nil {
		return err
	}

	_, err = m.client.SendText(resp.RoomID, text)
	if err != nil {
		return err
	}

	return nil
}
