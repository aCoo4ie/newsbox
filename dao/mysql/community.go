package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"
)

func GetCommunityList() (cl []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&cl, sqlStr)
	fmt.Println(cl, err)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}
