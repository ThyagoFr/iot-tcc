package device

type Device struct {
	Tank     string             `bson:"tank" json:"tank"`
	DeviceID int                `bson:"deviceID" json:"deviceID"`
}
