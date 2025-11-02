package types

// App represents an application
type App struct {
	Name        string  `json:"name"`
	Version     string  `json:"version"`
	Author      string  `json:"author"`
	Source      string  `json:"source"`
	Description string  `json:"description"`
	License     License `json:"license"`
}

// AppUser represents an application user
type AppUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
}

// AppUserActionLog represents an application user action log
type AppUserActionLog struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
}