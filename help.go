package help

func GetEmail(id string) string {
	return "lab@fake.com.br"
}

func Save(data interface{}) (bool, error) {
	return true, nil
}

func CheckIfExistSomeDebit(customerId string) int64 {
	return 10.00
}

func Update(data interface{}) (bool, error) {
	return true, nil
}

func Delete(id int) (bool, error) {
	return true, nil
}

func Get(id string) (bool, error) {
	return true, nil
}
