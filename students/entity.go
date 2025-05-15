package students

type Students struct {
	ID						int 	 `json:"id"`
	Name					string `json:"name" binding:"required"`
	NIM						string `json:"nim" binding:"required"`
	Major					string `json:"major" binding:"required"`
	GPA      			float32 `json:"gpa" binding:"required"`
	Graduated			bool	 `json:"graduated"`
}
