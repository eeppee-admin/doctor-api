package modelv2

// 文件库表

import (
	"time"

	"gorm.io/gorm"
)

// Patient 存储患者的基本信息
type Patient struct {
	gorm.Model
	PatientID        uint      `gorm:"primary_key;auto_increment;not null" json:"patient_id"`
	PatientName      string    `gorm:"type:varchar(50);not null" json:"patient_name"`
	PatientSex       string    `gorm:"type:varchar(10);not null" json:"patient_sex"`
	PatientBirthDate time.Time `gorm:"type:date;not null" json:"patient_birth_date"`
	PatientPhone     string    `gorm:"type:varchar(15);not null" json:"patient_phone"`
	PatientAddress   string    `gorm:"type:varchar(100)" json:"patient_address"`
	PatientIDCard    string    `gorm:"type:varchar(20);not null;unique" json:"patient_id_card"`
	// CreateTime       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	IsDeleted int `gorm:"not null" json:"is_deleted"`
}

// Doctor 存储医生的基本信息
type Doctor struct {
	gorm.Model
	DoctorID    uint   `gorm:"primary_key;auto_increment;not null" json:"doctor_id"`
	DoctorName  string `gorm:"type:varchar(50);not null" json:"doctor_name"`
	DoctorSex   string `gorm:"type:varchar(10);not null" json:"doctor_sex"`
	DoctorPhone string `gorm:"type:varchar(15);not null" json:"doctor_phone"`
	DoctorTitle string `gorm:"type:varchar(15)" json:"doctor_title"`
	// CreateTime       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	IsDeleted        int    `gorm:"not null" json:"is_deleted"`
	DoctorDepartment string `gorm:"type:varchar(50);not null" json:"doctor_department"`
}

// Visit 存储患者的就诊信息
type Visit struct {
	gorm.Model
	JZID          uint      `gorm:"primary_key;auto_increment;not null" json:"jz_id"`
	PatientID     uint      `gorm:"index;not null;foreignKey:PatientID" json:"patient_id"`
	DoctorID      uint      `gorm:"index;not null;foreignKey:DoctorID" json:"doctor_id"`
	VisitTime     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"visit_time"`
	Diagnosis     string    `gorm:"type:text" json:"diagnosis"`
	DiagnosisName string    `gorm:"type:varchar(100)" json:"diagnosis_name"`
	Status        string    `gorm:"type:varchar(20);not null" json:"status"`
	IsDeleted     int       `gorm:"not null" json:"is_deleted"`
}

// Order 存储医生为患者开具的医嘱
type Order struct {
	gorm.Model
	YZID         uint   `gorm:"primary_key;auto_increment;not null" json:"yz_id"`
	VisitID      uint   `gorm:"index;not null;foreignKey:VisitID" json:"visit_id"`
	DoctorID     uint   `gorm:"index;not null;foreignKey:DoctorID" json:"doctor_id"`
	OrderContent string `gorm:"type:varchar(100);not null" json:"order_content"`
	// CreateTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	IsDeleted int  `gorm:"not null" json:"is_deleted"`
	YZSign    int  `gorm:"not null" json:"yz_sign"`
	PatientID uint `gorm:"index;not null;foreignKey:PatientID" json:"patient_id"`
}

// LabTest 存储检验检查单项目表
type LabTest struct {
	gorm.Model
	LabTestID   string `gorm:"primary_key;type:varchar(20);not null" json:"labtest_id"`
	TestType    string `gorm:"type:varchar(10);not null" json:"test_type"`
	TestName    string `gorm:"type:varchar(50);not null" json:"test_name"`
	TestContent string `gorm:"type:text" json:"test_content"`
	TestUnit    string `gorm:"type:varchar(10)" json:"test_unit"`
}

// MedicalRecord 存储患者的病历信息
// type MedicalRecord struct {
// 	gorm.Model
// 	BLID          uint   `gorm:"primary_key;auto_increment;not null" json:"bl_id"`
// 	VisitID       uint   `gorm:"index;not null;foreignKey:VisitID" json:"visit_id"`
// 	DoctorID      uint   `gorm:"index;not null;foreignKey:DoctorID" json:"doctor_id"`
// 	RecordContent string `gorm:"type:text;not null" json:"record_content"`
// 	// CreateTime    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
// 	IsDeleted int  `gorm:"not null" json:"is_deleted"`
// 	PatientID uint `gorm:"index;not null;foreignKey:PatientID" json:"patient_id"`
// }

// User 存储系统用户（医生、管理员等）的账号信息
type User struct {
	gorm.Model
	UserID   uint   `gorm:"primary_key;auto_increment;not null" json:"user_id"`
	Username string `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role     string `gorm:"type:varchar(20);not null" json:"role"`
	// CreateTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	IsDeleted int `gorm:"not null" json:"is_deleted"`
}

// AdmissionNotice 存储患者入院通知单信息
type AdmissionNotice struct {
	gorm.Model
	RYTZDID       uint   `gorm:"primary_key;auto_increment;not null" json:"rytzd_id"`
	VisitID       uint   `gorm:"index;not null;foreignKey:VisitID" json:"visit_id"`
	DoctorID      uint   `gorm:"index;not null;foreignKey:DoctorID" json:"doctor_id"`
	NoticeContent string `gorm:"type:text;not null" json:"notice_content"`
	// CreateTime    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	IsDeleted int `gorm:"not null" json:"is_deleted"`
}

// WesternMedicline 存储西药信息表
type WesternMedicline struct {
	gorm.Model
	WesternMediclineID   string `gorm:"primary_key;type:varchar(20);not null" json:"western_medicline_id"`
	WesternMediclineName string `gorm:"type:varchar(50);not null" json:"western_medicline_name"`
	Ypjx                 string `gorm:"type:varchar(20)" json:"ypjx"`
	Ypgg                 string `gorm:"type:varchar(100)" json:"ypgg"`
}

// ChineseMedicline 存储中草药信息表
type ChineseMedicline struct {
	gorm.Model
	ChineseMediclineID   string `gorm:"primary_key;type:varchar(20);not null" json:"chinese_medicline_id"`
	ChineseMediclineName string `gorm:"type:varchar(50);not null" json:"chinese_medicline_name"`
	Gxfl                 string `gorm:"type:varchar(20)" json:"gxfl"`
	Gxyl                 string `gorm:"type:varchar(100)" json:"ggyl"`
}

// ICD_10 存储诊断编码表
type ICD_10 struct {
	gorm.Model
	DiagnosisID   string `gorm:"primary_key;type:varchar(20);not null" json:"diagnosis_id"`
	DiagnosisName string `gorm:"type:varchar(100);not null" json:"diagnosis_name"`
}
