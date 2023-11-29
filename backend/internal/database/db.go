package database

type Database interface {
	Database() Database
	Connect(string, string) error
	Close() error
	InsertData(data string) error
	GetData() ([]string, error)
}
