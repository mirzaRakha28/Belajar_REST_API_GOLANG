package models

import (
	"ECHO-REST/db"
	"net/http"
)

type Pegawai struct {
	Id        int    "json:'id'"
	Nama      string "json: 'nama'"
	Alamat    string "json: 'alamat'"
	Telephone string "json: 'telephone'"
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telephone)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama string, alamat string, telephone string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT pegawai (nama, alamat, telephone) VALUES (?,?,?)"
	statement, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := statement.Exec(nama, alamat, telephone)
	if err != nil {
		return res, nil
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]int64{
		"insert_id": lastInsertId,
	}
	return res, nil

}

func UpdatePegawai(nama string, alamat string, telephone string, id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE pegawai SET nama = ?, alamat = ?, telephone = ? WHERE id = ?"

	statement, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	alamat = alamat + " 1 "
	result, err := statement.Exec(nama, alamat, telephone, id)
	if err != nil {
		return res, nil
	}
	rowAffacted, err := result.RowsAffected()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]int64{

		"row_affected": rowAffacted,
	}
	return res, nil

}
func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
