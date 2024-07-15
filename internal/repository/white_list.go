package repository //nolint:dupl

import (
	"github.com/KarimovKamil/otus-go-final-project/internal/entity"
	"github.com/KarimovKamil/otus-go-final-project/internal/repository/client"
)

var (
	isExistsWL  = "SELECT EXISTS(SELECT 1 FROM white_list WHERE ip = $1 and mask = $2)"
	insertWL    = "INSERT INTO white_list (ip, mask, binary_prefix) VALUES ($1, $2, $3)"
	selectAllWL = "SELECT ip, mask, binary_prefix FROM white_list"
	deleteWL    = "DELETE FROM white_list WHERE ip = $1 and mask = $2"
)

type WhiteListRepo struct {
	client *client.PostgresSQL
}

func NewWhiteListRepo(client *client.PostgresSQL) *WhiteListRepo {
	return &WhiteListRepo{
		client: client,
	}
}

func (r *WhiteListRepo) IsExists(ip, mask string) bool {
	var exists bool
	err := r.client.DB.QueryRow(isExistsWL, ip, mask).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *WhiteListRepo) Add(network entity.Network) error {
	_, err := r.client.DB.Exec(insertWL, network.IP, network.Mask, network.BinaryPrefix)
	if err != nil {
		return err
	}
	return nil
}

func (r *WhiteListRepo) GetAll() ([]entity.Network, error) {
	networkList := make([]entity.Network, 10)
	err := r.client.DB.Select(&networkList, selectAllWL)
	if err != nil {
		return nil, err
	}
	return networkList, nil
}

func (r *WhiteListRepo) Remove(ip, mask string) error {
	_, err := r.client.DB.Exec(deleteWL, ip, mask)
	if err != nil {
		return err
	}
	return nil
}
