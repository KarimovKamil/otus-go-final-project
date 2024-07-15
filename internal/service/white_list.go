package service //nolint:dupl

import (
	"strings"

	"github.com/KarimovKamil/otus-go-final-project/internal/repository"
)

type WhiteList struct {
	repo *repository.WhiteListRepo
}

func NewWhiteList(repo *repository.WhiteListRepo) *WhiteList {
	return &WhiteList{
		repo: repo,
	}
}

func (w *WhiteList) IsContains(ip string) (bool, error) {
	networks, err := w.repo.GetAll()
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

func (w *WhiteList) Add(network string) error {
	networkEntity := GetNetwork(network)

	if w.repo.IsExists(networkEntity.IP, networkEntity.Mask) {
		return ErrNetworkAlreadyExists
	}

	err := w.repo.Add(networkEntity)
	if err != nil {
		return err
	}
	return nil
}

func (w *WhiteList) Remove(network string) error {
	networkEntity := GetNetwork(network)

	err := w.repo.Remove(networkEntity.IP, networkEntity.Mask)
	if err != nil {
		return err
	}
	return nil
}
