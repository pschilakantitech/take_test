package pg_persist

import (
	"database/sql"

	"github.com/pkg/errors"
)

type CoinResultsPGDB struct {
	CoinID    int
	CoinCount sql.NullFloat64
	ID        int
}

var (
	ErrNoRecords      = errors.New("no_records")
	ErrNoRowsAffected = errors.New("no_rows_affected")
)

const (
	selCoinsSQL           = "SELECT coin_id,coin_count FROM GetUserCoinAndCounts WHERE hash =$1"
	insUser               = `INSERT INTO "user"."user"(hash, created_on) VALUES ( $1, current_timestamp)`
	selUserIDByHash       = `SELECT user_id FROM "user"."user" WHERE hash=$1`
	selIDByUserAndCoinIDs = `SELECT id, coin_count FROM user_coins.user_coins WHERE user_id=$1  AND  coin_id=$2`
	insUserCoin           = `INSERT INTO user_coins.user_coins(user_id, coin_id, coin_count) VALUES ($1, $2, $3)`
	updUserCoin           = `UPDATE user_coins.user_coins SET coin_count=$1 WHERE id=$2`
)

//func GetUserCoinsAndCounts(hash string) (arr []*CoinResultsPGDB, err error) {
//	prepare, _ := Db.Prepare(selCoinsSQL)
//	rows, err := prepare.Query(hash)
//	defer rows.Close()
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to get coin details for hash"+hash)
//	}
//	for rows.Next() {
//		res := CoinResultsPGDB{}
//		if err := rows.Scan(&res.CoinID, &res.CoinCount); err != nil {
//			log.Error(err)
//		}
//		arr = append(arr, &res)
//	}
//	if len(arr) == 0 {
//		return nil, ErrNoRecords
//	}
//	return
//}

//func CreateUser(hash string) (err error) {
//	res, err := Db.Exec(insUser, hash)
//	if err != nil {
//		return errors.Wrap(err, "failed to create user for "+hash)
//	}
//	if v, _ := res.RowsAffected(); v <= 0 {
//		return ErrNoRowsAffected
//	}

//	return
//}

//func GetUserIDByHash(hash string) (id int, err error) {
//	prepare, _ := Db.Prepare(selUserIDByHash)
//	rows, err := prepare.Query(hash)
//	defer rows.Close()
//	if err != nil {
//		return 0, errors.Wrap(err, "failed to get coin details for hash"+hash)
//	}
//	if rows.Next() {
//		if err := rows.Scan(&id); err != nil {
//			return 0, errors.Wrap(err, "failed to scan user id for hash "+hash)
//		}
//	} else {
//		return 0, ErrNoRecords
//	}
//	return
//}

//func AddOrUpdateUserCoins(hash string, coinID int, coinCount float64) (err error) {
//	userID := 0
//	if userID, err = GetUserIDByHash(hash); err != nil {
//		return
//	}

//	prepare, _ := Db.Prepare(selIDByUserAndCoinIDs)
//	rows, err := prepare.Query(userID, coinID)
//	defer rows.Close()
//	if err != nil {
//		return errors.Wrap(err, fmt.Sprintf("failed to get ID details for user id : %d and coin id %d", userID, coinID))
//	}
//	if rows.Next() {
//		coinResp := CoinResultsPGDB{}
//		if err = rows.Scan(&coinResp.ID, &coinResp.CoinCount); err != nil {
//			return errors.Wrap(err, fmt.Sprintf("unable to get user_coin pk by for user id : %d, coin id : %d ", userID, coinID))
//		}
//		if coinCount != coinResp.CoinCount.Float64 {
//			// record found? yes!!!, update the record
//			return updateUserCoin(coinResp.ID, coinCount)
//		}
//	} else {
//		// No record found!!!, make an entry
//		return insertUserCoin(userID, coinID, coinCount)
//	}

//	return
//}

//func insertUserCoin(userID, coinID int, coinCount float64) (err error) {
//	res, err := Db.Exec(insUserCoin, userID, coinID, coinCount)
//	if err != nil {
//		return errors.Wrap(err, "add user coin failed")
//	}
//	if v, _ := res.RowsAffected(); v <= 0 {
//		return ErrNoRowsAffected
//	}

//	return
//}

//func updateUserCoin(id int, coinCount float64) (err error) {
//	res, err := Db.Exec(updUserCoin, coinCount, id)
//	if err != nil {
//		return errors.Wrap(err, "update user coin failed")
//	}
//	if v, _ := res.RowsAffected(); v <= 0 {
//		return ErrNoRowsAffected
//	}

//	return
//}

//func insertUserCoin(userID, coinID int, coinCount float64) (err error) {
//	res, err := Db.Exec(insUserCoin, userID, coinID, coinCount)
//	if err != nil {
//		return errors.Wrap(err, "add user coin failed")
//	}
//	if v, _ := res.RowsAffected(); v <= 0 {
//		return ErrNoRowsAffected
//	}

//	return
//}
