package db

import (
	"database/sql"
	"fmt"
	"sz_resume_202005/model"
	"sz_resume_202005/utils"
	"sz_resume_202005/utils/zlog"
	"time"

	"github.com/jmoiron/sqlx"
)

//TotalWork 计算作品总数
func TotalWork(args ...interface{}) (num int, err error) {
	var sqlStr string

	if len(args) == 0 {
		sqlStr = "SELECT COUNT(work_id) from work where status=1"
	} else {
		if args[0] == 0 {
			sqlStr = "SELECT COUNT(work_id) from work where status=1"
		} else {
			sqlStr = "SELECT COUNT(work_id) from work where status!=2"
		}

	}
	err = db.Get(&num, sqlStr)
	return
}

//TotalWorkByAuthor 计算作品总数
func TotalWorkByAuthor(userID int, args ...interface{}) (num int, err error) {
	var sqlStr string

	if len(args) == 0 {
		sqlStr = "SELECT COUNT(work_id) from work where user_id=?"
		err = db.Get(&num, sqlStr, userID)
		return
	}
	if n, ok := args[0].(int); ok && n >= 0 && n <= 2 {
		sqlStr = "SELECT COUNT(work_id) from work where status=? and user_id=?"
		err = db.Get(&num, sqlStr, n, userID)
		return
	}
	return 0, fmt.Errorf("参数不合法")
}

//TotalWorkByStatus 计算文章总数
func TotalWorkByStatus(args ...interface{}) (num int, err error) {
	var sqlStr string

	if len(args) == 0 {
		sqlStr = "SELECT COUNT(work_id) from work "
		err = db.Get(&num, sqlStr)
		return
	}
	if n, ok := args[0].(int); ok && n >= 0 && n <= 2 {
		sqlStr = "SELECT COUNT(work_id) from work where status=? "
		err = db.Get(&num, sqlStr, n)
		return
	}
	return 0, fmt.Errorf("参数不合法")
}

//GetWork 读取审核过的作品详情
func GetWork(workID int) (w *model.WorkDetail, err error) {
	w = &model.WorkDetail{}
	sqlStr := `
	SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
	       ,w.content,w.img,w.if_horizons
	FROM work w
	LEFT JOIN user u ON w.user_id=u.user_id
	WHERE w.status=1 AND work_id=?
  `

	err = db.Get(w, sqlStr, workID)
	if err != nil {
		return nil, err
	}

	sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
	err = db.Select(&(w.Tags), sqlStr1, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	sqlStr2 := `
			SELECT album_id,work_id,img_path,img_no,create_time,update_time
			FROM work_album
			WHERE work_id=?
			`
	err = db.Select(&(w.WorkImgs), sqlStr2, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	return
}

//GetPWork 读取所有状态的作品详情
func GetPWork(userID, workID int) (w *model.WorkDetail, err error) {
	w = &model.WorkDetail{}
	sqlStr := `
	SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
	       ,w.content,w.img,w.if_horizons
	FROM work w
	LEFT JOIN user u ON w.user_id=u.user_id
	WHERE w.user_id=? AND work_id=?
  `

	err = db.Get(w, sqlStr, userID, workID)
	if err != nil {
		return nil, err
	}

	sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
	err = db.Select(&(w.Tags), sqlStr1, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	sqlStr2 := `
			SELECT album_id,work_id,img_path,img_no,create_time,update_time
			FROM work_album
			WHERE work_id=?
			`
	err = db.Select(&(w.WorkImgs), sqlStr2, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	return
}

//GetRPWork 读取所有状态的作品详情
func GetRPWork(workID int) (w *model.WorkDetail, err error) {
	w = &model.WorkDetail{}
	sqlStr := `
	SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
	       ,w.content,w.img,w.if_horizons
	FROM work w
	LEFT JOIN user u ON w.user_id=u.user_id
	WHERE work_id=?
  `

	err = db.Get(w, sqlStr, workID)
	if err != nil {
		return nil, err
	}

	sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
	err = db.Select(&(w.Tags), sqlStr1, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	sqlStr2 := `
			SELECT album_id,work_id,img_path,img_no,create_time,update_time
			FROM work_album
			WHERE work_id=?
			`
	err = db.Select(&(w.WorkImgs), sqlStr2, workID)
	if err != nil {
		zlog.Error("err:", err)
		return nil, err
	}

	return
}

//GetWorksByAuthor 根据作者返回作品列表,用于后端管理
func GetWorksByAuthor(userID int, offset, limit int) (works []*model.Work, err error) {

	sqlStr := `
		SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
		FROM work w
		LEFT JOIN user u ON w.user_id=u.user_id
		WHERE w.user_id=?
		ORDER BY work_no desc
		LIMIT ? OFFSET ?
		`

	r, err := db.Queryx(sqlStr, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	for r.Next() {
		w := &model.Work{}
		err = r.StructScan(w)
		if err != nil {
			zlog.Error("err:", err)
			return nil, err
		}
		sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
		err = db.Select(&(w.Tags), sqlStr1, w.WorkID)
		if err != nil {
			zlog.Error("err:", err)
			return nil, err
		}
		works = append(works, w)
	}

	return works, nil
}

//GetWorksByTag 根据分页和分类返回作品列表，用于前端显示
func GetWorksByTag(tagID int, offset, limit int) (works []*model.Work, err error) {

	if tagID == 0 {
		fmt.Print("begin getworksbytag")
		sqlStr := `
		SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
		FROM work w
		LEFT JOIN user u ON w.user_id=u.user_id
		WHERE w.status=1 
		ORDER BY work_no desc
		LIMIT ? OFFSET ?
		`

		r, err := db.Queryx(sqlStr, limit, offset)
		if err != nil {
			return nil, err
		}
		for r.Next() {
			w := &model.Work{}
			err = r.StructScan(w)
			if err != nil {
				zlog.Error("err:", err)
				return nil, err
			}
			sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
			err = db.Select(&(w.Tags), sqlStr1, w.WorkID)
			if err != nil {
				zlog.Error("err:", err)
				return nil, err
			}
			works = append(works, w)
		}

	} else {
		sqlStr := `
		SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
		FROM work w
		LEFT JOIN user u ON w.user_id=u.user_id
		WHERE w.status=1  AND w.work_id in (SELECT work_id from work_tag_rl where tag_id=?)
		ORDER BY work_no desc
		LIMIT ? OFFSET ?
		`
		r, err := db.Queryx(sqlStr, tagID, limit, offset)
		if err != nil {
			return nil, err
		}
		for r.Next() {
			w := &model.Work{}
			err = r.StructScan(w)
			if err != nil {
				return nil, err
			}
			sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
			err := db.Select(&(w.Tags), sqlStr1, w.WorkID)
			if err != nil {
				return nil, err
			}
			works = append(works, w)
		}
	}
	return works, nil
}

//GetAllWorksByStatus 根据状态显示所有作品
func GetAllWorksByStatus(offset, limit int, args ...interface{}) (works []*model.Work, err error) {

	if len(args) == 0 {
		fmt.Print("begin getworksbytag")
		sqlStr := `
		SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
		FROM work w
		LEFT JOIN user u ON w.user_id=u.user_id
		ORDER BY work_no desc
		LIMIT ? OFFSET ?
		`

		r, err := db.Queryx(sqlStr, limit, offset)
		if err != nil {
			return nil, err
		}
		for r.Next() {
			w := &model.Work{}
			err = r.StructScan(w)
			if err != nil {
				zlog.Error("err:", err)
				return nil, err
			}
			sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
			err = db.Select(&(w.Tags), sqlStr1, w.WorkID)
			if err != nil {
				zlog.Error("err:", err)
				return nil, err
			}
			works = append(works, w)
		}

	}
	if n, ok := args[0].(int); ok && n >= 0 && n <= 2 {
		sqlStr := `
		SELECT w.work_id,w.title,w.cover_img,w.start_time,w.end_time,w.user_id,w.work_no,w.status,w.create_time,w.update_time,u.user_name,u.avatar,w.view_count
		FROM work w
		LEFT JOIN user u ON w.user_id=u.user_id
		WHERE w.status=?
		ORDER BY work_no desc
		LIMIT ? OFFSET ?
		`
		r, err := db.Queryx(sqlStr, n, limit, offset)
		if err != nil {
			return nil, err
		}
		for r.Next() {
			w := &model.Work{}
			err = r.StructScan(w)
			if err != nil {
				return nil, err
			}
			sqlStr1 := `
			SELECT r.rl_id,r.tag_id,t.tag_name,r.work_id,r.create_time,r.update_time
			FROM work_tag_rl r
			LEFT JOIN work_tag t ON r.tag_id=t.tag_id
			WHERE work_id=?
			`
			err := db.Select(&(w.Tags), sqlStr1, w.WorkID)
			if err != nil {
				return nil, err
			}
			works = append(works, w)
		}
	}
	return works, nil
}

//AddWork 新增作品--事务，多表
func AddWork(userID int, work *model.WorkDetail) (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			zlog.Errorf("trans failed, err:%v\n", err)
			e := tx.Rollback()
			if e != nil {
				zlog.Panicf("transaction err:%v", err)
				panic(e)
			}
		}
	}()

	//work表新增
	sqlStr := `
	INSERT INTO work (title,work_no,img,cover_img,content,start_time,end_time,if_horizons,user_id) 
	VALUES(?,?,?,?,?,?,?,?,?)
	`
	result, err := tx.Exec(sqlStr, work.Title, work.WorkNO, work.Img, work.CoverImg, work.Content, work.StartTime, work.EndTime, work.IfHorizons, userID)
	if err != nil {
		zlog.Error(err)
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		zlog.Error(err)
		return
	}

	workID := utils.Int64to32(id64)
	workImgs := work.WorkImgs
	//work_album表新增
	if l := len(workImgs); l > 0 {

		sqlStri := "INSERT INTO work_album (work_id,img_path,img_no)  VALUES (?,?,?)"
		for _, v := range workImgs {
			v.WorkID = workID

			_, err = tx.Exec(sqlStri, v.WorkID, v.ImgPath, v.ImgNo)
			if err != nil {
				zlog.Error(err)
				return err
			}
		}

	}

	//work_tag_rl表新增
	workTags := work.Tags
	if l := len(workTags); l > 0 {
		sqlStrt := "INSERT INTO work_tag_rl (tag_id,work_id) VALUES (?,?)"
		for _, v := range workTags {
			v.WorkID = workID
			_, err = tx.Exec(sqlStrt, v.TagID, v.WorkID)
			if err != nil {
				zlog.Error(err)
				return err
			}
		}
	}

	err = tx.Commit()
	return
}

//EditWork 修改作品--事务，多表
func EditWork(userID int, work *model.WorkDetail) (err error) {

	tx, err := db.Beginx() // 开启事务
	if err != nil {
		return
	}
	if work.WorkID == 0 {
		return fmt.Errorf("workID not exsit")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			zlog.Errorf("trans failed, err:%v\n", err)
			e := tx.Rollback()
			if e != nil {
				zlog.Panicf("transaction err:%v", err)
				panic(e)
			}
		}
	}()

	//work表新增
	sqlStr := `
	UPDATE work SET title=?,work_no=?,img=?,cover_img=?,content=?,start_time=?,end_time=?,if_horizons=?
	WHERE work_id=? And user_id=?
	`
	_, err = tx.Exec(sqlStr, work.Title, work.WorkNO, work.Img, work.CoverImg, work.Content, work.StartTime, work.EndTime, work.IfHorizons, work.WorkID, userID)
	if err != nil {
		zlog.Error(err)
		return
	}

	workImgs := work.WorkImgs
	//work_album表新增
	zlog.Debug("len(workImgs)", len(workImgs))
	if l := len(workImgs); l > 0 {
		ids := []int{}

		for _, v := range workImgs {
			v.WorkID = work.WorkID
			if v.AlbumID != 0 {
				ids = append(ids, v.AlbumID)
			}
		}
		zlog.Debug("图库ids:", ids)
		if len(ids) < 1 {
			sqlStr1 := `
			DELETE FROM work_album WHERE work_id=?
			`
			_, err = tx.Exec(sqlStr1, work.WorkID)
			if err != nil {
				zlog.Error(err)
				return err
			}
		} else {
			sqlStr1 := `
			DELETE FROM work_album WHERE album_id NOT IN(?) AND work_id=?
			`
			query, args, err := sqlx.In(sqlStr1, ids, work.WorkID)
			if err != nil {
				zlog.Error(err)
				return err
			}
			_, err = tx.Exec(query, args...)
			if err != nil {
				zlog.Error(err)
				return err
			}
		}

		sqlStr2 := fmt.Sprintf(`
		INSERT INTO work_album (album_id,work_id,img_path,img_no) VALUES %s
		ON DUPLICATE KEY 
		UPDATE img_path=VALUES(img_path),img_no=VALUES(img_no)
		`, batchStringParam(l))
		query, args, err := sqlx.In(sqlStr2, workImgs.ConvInterfaceArray()...)
		query += ",update_time=?"
		fmt.Printf("query:%v\n", query)
		args = append(args, time.Now())
		if err != nil {
			zlog.Error(err)
			return err
		}
		_, err = tx.Exec(query, args...)
		if err != nil {
			zlog.Error(err)
			return err
		}

	} else {
		sqlStr3 := `
		DELETE FROM work_album where work_id=?
		`
		_, err = tx.Exec(sqlStr3, work.WorkID)
		if err != nil {
			zlog.Error(err)
			return err
		}
	}

	//work_tag_rl表新增
	workTags := work.Tags
	if l := len(workTags); l > 0 {

		ids := []int{}

		for _, v := range workTags {
			v.WorkID = work.WorkID
			if v.RlID != 0 {
				ids = append(ids, v.RlID)
			}
		}
		zlog.Debug("标签tagids:", ids)
		if len(ids) < 1 {
			sqlStr1 := `
			DELETE FROM work_tag_rl WHERE work_id=?
			`
			_, err = tx.Exec(sqlStr1, work.WorkID)
			if err != nil {
				return err
			}
		} else {
			sqlStrt1 := `
			DELETE FROM work_tag_rl WHERE rl_id NOT IN(?) AND work_id=?
			`
			query, args, err := sqlx.In(sqlStrt1, ids, work.WorkID)
			if err != nil {
				return err
			}
			_, err = tx.Exec(query, args...)
			if err != nil {
				return err
			}
		}

		for _, v := range workTags {
			v.WorkID = work.WorkID
		}
		sqlStrt2 := fmt.Sprintf(`
		INSERT INTO work_tag_rl (rl_id,tag_id,work_id) VALUES %s
		ON DUPLICATE KEY 
		UPDATE`, batchStringParam(l))
		s, args, err := sqlx.In(sqlStrt2, workTags.ConvInterfaceArray()...)
		s += " update_time=?"
		fmt.Print("query:", s)
		args = append(args, time.Now())
		if err != nil {
			return err
		}
		_, err = tx.Exec(s, args...)
		if err != nil {
			return err
		}
	} else {
		sqlStrt3 := `
		DELETE FROM work_tag_rl where work_id=?
		`
		_, err = tx.Exec(sqlStrt3, work.WorkID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return
}

//DelWorksFE 永久删除作品--事务，多表
func DelWorksFE(ids []int) (err error) {

	sqlStr := "DELETE FROM work WHERE work_id IN (?)"
	query, args, err := sqlx.In(sqlStr, ids)
	_, err = db.Exec(query, args...)
	if err != nil {
		return
	}

	return
}

//DelWorks 删除作品--软删除 --事务，多表
func DelWorks(userID int, ids []int) (err error) {
	sqlStr := `
	  UPDATE work SET status=? WHERE user_id=? AND work_id IN (?) 
	`
	query, args, err := sqlx.In(sqlStr, 2, userID, ids)
	_, err = db.Exec(query, args...)
	return
}

//ExistWorkByAuth 检查是否有操作作品的权限
func ExistWorkByAuth(userID, workID int) (bool, error) {
	sqlStr := "select work_id from work where work_id=? and user_id=? "
	var rwork model.Work
	err := db.Get(&rwork, sqlStr, workID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//ExistWorkByID 前端判断作品是否存在
func ExistWorkByID(workID int) (bool, error) {
	sqlStr := "SELECT work_id from work where work_id=? and status=1 "
	var rwork model.Work
	err := db.Get(&rwork, sqlStr, workID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err

}

//ExistWork 管理员查看是否存在
func ExistWork(workID int) (bool, error) {
	sqlStr := "select w.work_id from work w where w.work_id=?"
	var rwork model.Work
	err := db.Get(&rwork, sqlStr, workID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//CheckWorks 作品审核
func CheckWorks(ids []int, status int) (err error) {
	sqlStr := "UPDATE work w SET w.status=? where w.work_id in (?) "
	query, args, err := sqlx.In(sqlStr, status, ids)
	if err != nil {
		return
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return
	}

	return

}

//CountWork 访问数量
func CountWork(id int) (err error) {
	sqlStr := "update work set view_count=view_count+1 where work_id=?"
	_, err = db.Exec(sqlStr, id)
	return
}
