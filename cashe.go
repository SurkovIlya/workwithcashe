package cash

import (
	"fmt"
	"time"
)

const hotelsCount = 1000

type Hotel struct {
	ID            uint32
	Name          string
	PriceByNight  uint16
	Location      string
	LastUsageTime time.Time
}

type Cash struct {
	Hotel map[uint32]Hotel
	TTL   uint32
}

func NewCash(ttlMs uint32) *Cash {
	hotels := make(map[uint32]Hotel, hotelsCount)

	fmt.Println(hotels)
	return &Cash{
		Hotel: hotels,
		TTL:   ttlMs,
	}

}

func (c *Cash) GetHotelByID(id uint32) (Hotel, error) {
	if hotelValue, ok := c.Hotel[id]; ok {
		fmt.Println(hotelValue)
		return hotelValue, nil
	} else {
		return Hotel{}, fmt.Errorf("Asd")
	}
}
func (c *Cash) AddHotel(hotel Hotel) {
	if _, ok := c.Hotel[hotel.ID]; ok {
		c.Hotel[hotel.ID] = hotel
	} else {
		c.Hotel[hotel.ID] = hotel
	}
}
func (c *Cash) Clean() {

}
