package mysql

import (
	"bluebell/models"
	"database/sql"
)

func GetCommunityList() (cl []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&cl, sqlStr)
	// fmt.Println(cl, err)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (cd *models.CommunityDetail, err error) {
	sqlStr := `select community_name, introduction from community where community_id = ?`
	var detail models.CommunityDetail
	err = db.Get(&detail, sqlStr, id)
	// fmt.Println(detail, err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidId
		}
		return nil, err
	}
	return &detail, nil
}
