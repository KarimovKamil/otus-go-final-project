package repository //nolint:dupl

import (
	"github.com/KarimovKamil/otus-go-final-project/internal/entity"
	"github.com/KarimovKamil/otus-go-final-project/internal/repository/client"
)

var (
	isExistsBL  = "SELECT EXISTS(SELECT 1 FROM black_list WHERE ip = $1 and mask = $2)"
	insertBL    = "INSERT INTO black_list (ip, mask, binary_prefix) VALUES ($1, $2, $3)"
	selectAllBL = "SELECT ip, mask, binary_prefix FROM black_list"
	deleteBL    = "DELETE FROM black_list WHERE ip = $1 and mask = $2"
)

type BlackListRepo struct {
	client *client.PostgresSQL
}

func NewBlackListRepo(client *client.PostgresSQL) *BlackListRepo {
	return &BlackListRepo{
		client: client,
	}
}

func (r *BlackListRepo) IsExists(ip, mask string) bool {
	var exists bool
	err := r.client.DB.QueryRow(isExistsBL, ip, mask).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *BlackListRepo) Add(network entity.Network) error {
	_, err := r.client.DB.Exec(insertBL, network.IP, network.Mask, network.BinaryPrefix)
	if err != nil {
		return err
	}
	return nil
}

func (r *BlackListRepo) GetAll() ([]entity.Network, error) {
	networkList := make([]entity.Network, 10)
	err := r.client.DB.Select(&networkList, selectAllBL)
	if err != nil {
		return nil, err
	}
	return networkList, nil
}

func (r *BlackListRepo) Remove(ip, mask string) error {
	_, err := r.client.DB.Exec(deleteBL, ip, mask)
	if err != nil {
		return err
	}
	return nil
}
