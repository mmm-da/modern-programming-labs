package parking

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
	expRand "golang.org/x/exp/rand"
)

type Car struct {
	id       string
	patience time.Duration
	reserve  time.Duration
}

func (c Car) Process(parking chan int, stop chan struct{}) {

	endOfpatience := make(chan struct{})
	go func() {
		defer close(endOfpatience)
		time.Sleep(time.Millisecond * c.patience)
	}()

	select {
	case place := <-parking:
		fmt.Printf("%s take place %d\n", c.id, place)
		time.Sleep(time.Millisecond * c.reserve)
		parking <- place
		fmt.Printf("%s leave place %d\n", c.id, place)
	case <-endOfpatience:
		fmt.Printf("%s leave, end of patience\n", c.id)
	case <-stop:
		fmt.Printf("parking is closed, %s leaves\n", c.id)
	}
}

func CreateTraffic(lambda float64, stop chan struct{}) <-chan Car {
	traffic := make(chan Car)
	go func() {
		defer close(traffic)
		for {
			select {
			case <-stop:
				break
			default:
				traffic <- CreateMockCar()
				newDuration := int64(math.RoundToEven(expRand.ExpFloat64() / lambda))
				time.Sleep(time.Duration(newDuration) * time.Millisecond)
			}
		}
	}()
	return traffic
}

func InitPlaces(parking chan int, stop chan struct{}, count int) {
	for i := 1; i <= count; i++ {
		select {
		case <-stop:
			break
		default:
			parking <- i
			fmt.Printf("create parking place %d", i)
		}
	}
}

func CreateMockCar() Car {
	rand.Seed(time.Now().UTC().UnixNano())
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	patiente := time.Duration(rand.Int63n(300) + 1)
	reserve := time.Duration(rand.Int63n(300) + 1)
	result := Car{id: id.String(), patience: patiente, reserve: reserve}
	fmt.Printf("car %s, patiente %d, reserve %d\n", result.id, result.patience, result.reserve)
	return result
}
