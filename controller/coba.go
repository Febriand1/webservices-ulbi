package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	inimodel "github.com/Febriand1/Nilai/Model"
	inimodul "github.com/Febriand1/Nilai/Module"
	"github.com/Febriand1/webservices-ulbi/config"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	inimodellatihan "github.com/indrariksa/be_presensi/model"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
)

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensinew [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn1, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensinew/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn1, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

//latihan kemaren

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/Febriand1/webservices-ulbi",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

//func Homepage(c *fiber.Ctx) error {
//	ipaddr := musik.GetIPaddress()
//	return c.JSON(ipaddr)
//}

func GetPresensiold(c *fiber.Ctx) error {
	nl := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(nl)
}

//nilai
func GetPresensi(c *fiber.Ctx) error {
	nl := inimodul.GetPresensiFromMahasiswa("budiman", config.Ulbimongoconn, "presensi")
	return c.JSON(nl)
}

func GetMatakuliah(c *fiber.Ctx) error {
	nl := inimodul.GetMatakuliahFromJadwal("07.00", config.Ulbimongoconn, "matakuliah")
	return c.JSON(nl)
}

func GetGrade(c *fiber.Ctx) error {
	nl := inimodul.GetGradeFromMahasiswa(121395, config.Ulbimongoconn, "nilai")
	return c.JSON(nl)
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllNilaiFromNamaMahasiswa("budiman", config.Ulbimongoconn, "nilai")
	return c.JSON(nl)
}

func InsertIdentitas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa inimodel.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := inimodul.InsertMahasiswa(db, "mahasiswa",
		mahasiswa.Nama,
		mahasiswa.NPM,
		mahasiswa.Phone_Number)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertMatkul(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var matakuliah inimodel.Matakuliah
	if err := c.BodyParser(&matakuliah); err != nil {
		return err
	}
	insertedID := inimodul.InsertMatakuliah(db, "matakuliah",
		matakuliah.Nama_MK,
		matakuliah.SKS,
		matakuliah.Jadwal,
		matakuliah.Pengampu)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Matkul berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var dosen inimodel.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return err
	}
	insertedID := inimodul.InsertDosen(db, "dosen",
		dosen.Nama_Dosen,
		dosen.NIK,
		dosen.Phone_NumberD)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Dosen berhasil disimpan.",
		"inserted_id": insertedID,
	})
}



func GetAllMahasiswa(c *fiber.Ctx) error {
	ps := inimodul.GetMahasiswa(config.Ulbimongoconn, "mahasiswa")
	return c.JSON(ps)
}

func GetAllMatakuliah(c *fiber.Ctx) error {
	ps := inimodul.GetMatakuliah(config.Ulbimongoconn, "matakuliah")
	return c.JSON(ps)
}

//TB
// GetAllNilai godoc
// @Summary Get All Data Nilai.
// @Description Mengambil semua data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Success 200 {object} Nilai
// @Router /nilai [get]
func GetAllNilai(c *fiber.Ctx) error {
	ps := inimodul.GetAllNilai(config.Ulbimongoconn, "nilai")
	return c.JSON(ps)
}

// GetNilaiID godoc
// @Summary Get By ID Data Nilai.
// @Description Ambil per ID data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /nilai/{id} [get]
func GetNilaiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	nl, err := inimodul.GetNilaiFromID(objID, config.Ulbimongoconn, "nilai")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(nl)
}

// func InsertDataNilai(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var nilai inimodel.Nilai
// 	if err := c.BodyParser(&nilai); err != nil {
// 		return err
// 	}
// 	insertedID := inimodul.InsertNilai(db, "nilai",
// 		nilai.All_Tugas,
// 		nilai.UTS,
// 		nilai.UAS,
// 		nilai.Grade,
// 		nilai.Kategori,
// 		nilai.Absensi)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }

// InsertData godoc
// @Summary Insert data nilai.
// @Description Input data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param request body Nilai true "Payload Body [RAW]"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 500
// @Router /insnilai [post]
func InsertDataNilai(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var nilai inimodel.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodul.InsertNilai(db, "nilai",
		nilai.All_Tugas,
		nilai.UTS,
		nilai.UAS,
		nilai.Grade,
		nilai.Kategori,
		nilai.Absensi)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data nilai.
// @Description Ubah data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Nilai true "Payload Body [RAW]"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 500
// @Router /updnilai/{id} [put]
func UpdateDataNilai(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a nilai object
	var nilai inimodel.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateNilai function with the parsed ID and the Nilai object
	err = inimodul.UpdateNilai(db, "nilai",
		objectID,
		nilai.All_Tugas,
		nilai.UTS,
		nilai.UAS,
		nilai.Grade,
		nilai.Kategori,
		nilai.Absensi)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteNilaiByID godoc
// @Summary Delete data nilai.
// @Description Hapus data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delnilai/{id} [delete]
func DeleteNilaiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodul.DeleteNilaiByID(objID, config.Ulbimongoconn, "nilai")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
//TB

var jwtSecretKey = []byte("secretKey")

func LoginAdmin(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var admin inimodel.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	authenticated, err := inimodul.LoginAdmin(db, "admin", admin.Username, admin.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if authenticated {
		// Membuat token JWT
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = admin.Username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Menandatangani token dengan secret key
		tokenString, err := token.SignedString(jwtSecretKey)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "Login successful",
			"token":   tokenString,
		})
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"status":  http.StatusUnauthorized,
		"message": "Invalid credentials",
	})
}

func Authenticated(c *fiber.Ctx) error {
	// tokenString := c.Get("Authorization")

	var token inimodel.Token
	if err := c.BodyParser(&token); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
    tkn := token.Token_String
	// Check if token exists
	if tkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Parse token
	initoken, err := jwt.Parse(tkn, func(initoken *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := initoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Validate token claims
	if claims, ok := initoken.Claims.(jwt.MapClaims); ok && initoken.Valid {
		// c.Locals("username", claims["username"])
		// return c.Next()
		return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"username":      claims["username"],
	})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Invalid token",
	})
}








// //login
// func InsertLoginAdmin(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var admin inimodel.Admin
// 	if err := c.BodyParser(&admin); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	insertedID, err := inimodul.InsertAdmin(db, "admin",
// 		admin.Username,
// 		admin.Password)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }

// func GetAdmin(c *fiber.Ctx) error {
// 	ps := inimodul.GetAdmin(config.Ulbimongoconn, "admin")
// 	return c.JSON(ps)
// }

// func GetAdminID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Wrong parameter",
// 		})
// 	}
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "Invalid id parameter",
// 		})
// 	}
// 	nl, err := inimodul.GetAdminFromID(objID, config.Ulbimongoconn, "admin")
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 				"status":  http.StatusNotFound,
// 				"message": fmt.Sprintf("No data found for id %s", id),
// 			})
// 		}
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": fmt.Sprintf("Error retrieving data for id %s", id),
// 		})
// 	}
// 	return c.JSON(nl)
// }

// func InsertLoginUser(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var user inimodel.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	insertedID, err := inimodul.InsertUser(db, "user",
// 		user.UsernameM,
// 		user.PasswordM)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }

// func GetUser(c *fiber.Ctx) error {
// 	ps := inimodul.GetUser(config.Ulbimongoconn, "user")
// 	return c.JSON(ps)
// }

// func GetUserID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Wrong parameter",
// 		})
// 	}
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "Invalid id parameter",
// 		})
// 	}
// 	nl, err := inimodul.GetUserFromID(objID, config.Ulbimongoconn, "user")
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 				"status":  http.StatusNotFound,
// 				"message": fmt.Sprintf("No data found for id %s", id),
// 			})
// 		}
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": fmt.Sprintf("Error retrieving data for id %s", id),
// 		})
// 	}
// 	return c.JSON(nl)
// }
// //login

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData1(c *fiber.Ctx) error {
	db := config.Ulbimongoconn1
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}


// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn1

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}


// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn1, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}