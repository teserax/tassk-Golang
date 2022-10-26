//заполнить матрицу построчно от начала и до конца зигзагом

package main

import "fmt"

func main() {
	const N = 9
	const R = 9
	m := [N][R]int{}
	var vertical, horizon, count int
	count = 1
	for count < N*R {

		m[vertical][horizon] = count
		//позиция 1 движение в лево и вниз
		if vertical == 0 && horizon%2 == 0 && horizon != len(m)-1 {
			if horizon < len(m)-1 {
				horizon++
			}
			count++
			m[vertical][horizon] = count
			fmt.Println(" ty1")
			for horizon != 0 {
				horizon--
				vertical++
				count++
				m[vertical][horizon] = count

			}
			//позиция 2  стремимся  вправо и верх
		} else if horizon == 0 && vertical != 0 {

			fmt.Println("ty2", vertical, horizon)
			vertical++
			count++
			m[vertical][horizon] = count
			for vertical != 0 {
				vertical--
				horizon++
				count++
				m[vertical][horizon] = count

			}
			//позиция 3 влево и вниз
		} else if horizon == len(m)-1 && horizon%2 == 0 {
			fmt.Println("ty3", vertical, horizon)
			vertical++
			count++
			m[vertical][horizon] = count
			fmt.Println(m[vertical][horizon])
			for vertical != len(m)-1 {

				vertical++
				horizon--
				count++
				m[vertical][horizon] = count

			}
			//позиция 4
			// в право и верх
		} else if vertical == len(m)-1 && horizon != 0 {
			fmt.Println("ty4")
			horizon++
			count++
			m[vertical][horizon] = count
			for horizon != len(m)-1 {
				vertical--
				horizon++
				count++
				m[vertical][horizon] = count
			}

		}

	}

	for _, val := range m {
		fmt.Printf("%2d\n", val)
	}

}
