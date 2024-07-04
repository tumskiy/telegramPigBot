package htopd

import (
	"database/sql"
	"fmt"
	"time"
)

func (pd *PD) CreateTodayPD() (bool, int64, error) {
	lastPDdate, _, err := pd.getLastPD()
	if err != nil {
		return false, 0, err
	}

	today := time.Now()
	if lastPDdate.Year() == today.Year() && lastPDdate.Month() == today.Month() && lastPDdate.Day() == today.Day() {
		_, pid, _ := pd.getLastPD()
		return false, pid, nil
	}

	user := &User{}
	findVictim := user.UserVictim()
	err = pd.createPD(findVictim)
	if err != nil {
		return false, 0, err
	}

	return true, findVictim, nil
}

func (pd *PD) createPD(userID int64) error {
	count, err := pd.countPD()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO pd VALUES (?, ?, ?, ?)", count+1, userID, pd.ChatID, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (pd *PD) countPD() (int64, error) {
	var result int64
	exec := db.QueryRow("SELECT COUNT(*) FROM pd")
	if exec == nil {
		result = 0
		return result, nil
	}
	err := exec.Scan(&result)
	if err != nil {
		fmt.Println("db troubles:", err)
		return 0, err // Возвращаем false и ошибку
	}

	return result, nil
}

func (pd *PD) getLastPD() (time.Time, int64, error) {
	exec := db.QueryRow(`SELECT id, user_id, chat_id, date
		FROM pd
		ORDER BY date DESC
		LIMIT 1`)
	if err := exec.Scan(&pd.ID, &pd.UserID, &pd.ChatID, &pd.Date); err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, 0, nil
		}
		return time.Time{}, 0, err
	}
	return pd.Date, pd.UserID, nil
}

func (pd *PD) GetCountPDByUserID(userID int64) (int, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) as count FROM PD WHERE user_id = ?`, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
