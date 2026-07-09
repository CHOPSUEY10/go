package intersections

import (
	"crossroads/trafficlight"
	"fmt"
	"sync"
	"time"
)

type Intersections struct {
	Intersect map[string]*trafficlight.TrafficLight
}

func (i *Intersections) AddTrafficLight(t *trafficlight.TrafficLight) {

	if i.Intersect == nil {
		i.Intersect = make(map[string]*trafficlight.TrafficLight)
	}
	i.Intersect[t.RoadName] = t

}

func (i *Intersections) RemoveTrafficLight(t *trafficlight.TrafficLight) {

	if i.Intersect == nil {
		i.Intersect = make(map[string]*trafficlight.TrafficLight)
	}

	if t.RoadName == "" {

		fmt.Printf("Nama jalan tidak ditemukan")

	}
	delete(i.Intersect, t.RoadName)

}
func (i *Intersections) RunTrafficLight(t *trafficlight.TrafficLight) {

	fmt.Printf("\nLampu lalu lintas %s\n", t.RoadName)
	for i := 0; i < t.Time; i++ {
		fmt.Printf("\tWaktu tunggu : %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Selesai\n\n")

}

func (i *Intersections) SwitchTrafficLight(q *Queue) {

	// Sistem kontrol konkurensi (WaitGroup)
	var wg sync.WaitGroup
	for _, v := range i.Intersect {
		// Menambahkan goroutine ke workgroup setiap satu kali iterasi
		wg.Add(1)
		// goroutine yang dikerjakan dalam bentuk anon func
		go func(light *trafficlight.TrafficLight) {
			// Memberitahu Workgroup jika satu goroutine selesai
			defer wg.Done()
			// Pekerjaan yang dilakukan goroutine
			i.RunTrafficLight(light)
		}(v)

		wg.Wait()
	}

}
