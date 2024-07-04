package htopd

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func existUser(userID int64) (bool, error) {
	var result int
	exec := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userID)
	err := exec.Scan(&result)
	if err != nil {
		fmt.Println("db troubles:", err)
		return false, err // Возвращаем false и ошибку
	}
	if result != 0 {
		return true, nil // Возвращаем true и nil (без ошибки)
	}

	return false, nil // Возвращаем false и nil (без ошибки)
}

func (u *User) UserVictim() int64 {
	var results []User
	rows, err := db.Query("SELECT * FROM users") //получаем массив юзеров
	if err != nil {
		fmt.Println("Проблемы с базой данных:", err)
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			fmt.Println("Проблемы с базой данных:", err)
			continue
		}
		results = append(results, user) //прибавляем к результату
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Проблемы с базой данных:", err)
		return 0
	}

	if len(results) == 0 {
		return 0
	}

	randomIndex := rand.Intn(len(results)) // рандомим
	randomUser := results[randomIndex]     // назначаем индекс

	return randomUser.ID

}

func (u *User) CreateUser() error {
	exist, err := existUser(u.ID)
	if err != nil {
		return err
	}

	if !exist {
		_, err := db.Exec("INSERT INTO users VALUES (?, ?)", u.ID, u.Name)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (u *User) GetUserByID(userID int64) (*User, error) {
	exec := db.QueryRow("SELECT id, name FROM users WHERE id = ?", userID) // Указываем два столбца для сканирования
	err := exec.Scan(&u.ID, &u.Name)                                       // Сканируем результат запроса в поля структуры User
	if err != nil {
		fmt.Println("Проблемы с базой данных:", err)
		return nil, err
	}
	return u, nil
}

func (u *User) GetAllUsers() ([]*User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		fmt.Println("Проблемы с базой данных:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			fmt.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Ошибка при обходе строк:", err)
		return nil, err
	}

	return users, nil
}
