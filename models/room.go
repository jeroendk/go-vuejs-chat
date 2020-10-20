package models

type Room interface {
	GetId() string
	GetName() string
	GetPrivate() bool
}

type RoomRepository interface {
	AddRoom(room Room)
	FindRoomByName(name string) Room
}
