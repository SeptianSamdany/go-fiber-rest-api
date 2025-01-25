package entities

type Job struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string `json:"title"`       // Nama pekerjaan
    Description string `json:"description"` // Deskripsi pekerjaan
    Users       []User `json:"users"`       // Relasi dengan User
}
