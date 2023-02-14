package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID       string
	ClientID string
	Position []Position
}

type Position struct {
	Lat  float64
	Long float64
}

// LoadPositions carrega de um arquivo .txt todas as posições (lat e long) para o atributo Position da struct

func (r *Route) LoadPosition() error {
	if r.ID == "" {
		return errors.New("Informe a Rota do ID")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[8], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil
}
