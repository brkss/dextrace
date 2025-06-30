package domain



type AuthRepository interface {
	Login(user User) (string, error)
}

type GlucoseRepository interface {
	GetData(token string, userID string) (*GlucoseDataResponse, error)
}

type NightscoutRepository interface {
	PushData(data []GetDataResponse) error;
	GetLastRecord() (*NightscoutEntry, error);
}