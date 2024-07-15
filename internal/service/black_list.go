package service //nolint:dupl

import (
	"strings"

	"github.com/KarimovKamil/otus-go-final-project/internal/repository"
)

type BlackList struct {
	repo *repository.BlackListRepo
}

func NewBlackList(repo *repository.BlackListRepo) *BlackList {
	return &BlackList{
		repo: repo,
	}
}

func (b *BlackList) IsContains(ip string) (bool, error) {
	networks, err := b.repo.GetAll()
	if err != nil {
		return false, err
	}

	binaryIP := IPAddressToBinary(ip)
	for _, network := range networks {
		if strings.HasPrefix(binaryIP, network.BinaryPrefix) {
			return true, nil
		}
	}
	return false, err
}

func (b *BlackList) Add(network string) error {
	networkEntity := GetNetwork(network)

	if b.repo.IsExists(networkEntity.IP, networkEntity.Mask) {
		return ErrNetworkAlreadyExists
	}

	err := b.repo.Add(networkEntity)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlackList) Remove(network string) error {
	networkEntity := GetNetwork(network)

	err := b.repo.Remove(networkEntity.IP, networkEntity.Mask)
	if err != nil {
		return err
	}
	return nil
}
